/*
Copyright 2025 the Agent Control Plane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package framework

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	acp "github.com/humanlayer/agentcontrolplane/acp/api/v1alpha1"
	. "github.com/humanlayer/agentcontrolplane/acp/test/utils"
)

var _ = Describe("Framework Tests", func() {
	It("should have working Kubernetes client", func() {
		ctx := testFramework.GetContext()
		client := testFramework.GetClient()

		By("creating a simple secret to test connectivity")
		secret := &corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "connectivity-test",
				Namespace: "default",
			},
			Data: map[string][]byte{
				"test": []byte("data"),
			},
		}

		err := client.Create(ctx, secret)
		Expect(err).NotTo(HaveOccurred())

		By("verifying the secret was created")
		createdSecret := &corev1.Secret{}
		err = client.Get(ctx, types.NamespacedName{
			Name:      "connectivity-test",
			Namespace: "default",
		}, createdSecret)
		Expect(err).NotTo(HaveOccurred())
		Expect(createdSecret.Data["test"]).To(Equal([]byte("data")))

		By("cleaning up the secret")
		err = client.Delete(ctx, secret)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should create LLM resources", func() {
		ctx := testFramework.GetContext()
		client := testFramework.GetClient()

		By("creating an LLM resource")
		llm := &acp.LLM{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-llm-simple",
				Namespace: "default",
			},
			Spec: acp.LLMSpec{
				Provider: "openai",
			},
		}

		err := client.Create(ctx, llm)
		Expect(err).NotTo(HaveOccurred())

		By("verifying the LLM was created")
		createdLLM := &acp.LLM{}
		err = client.Get(ctx, types.NamespacedName{
			Name:      "test-llm-simple",
			Namespace: "default",
		}, createdLLM)
		Expect(err).NotTo(HaveOccurred())
		Expect(createdLLM.Spec.Provider).To(Equal("openai"))

		By("cleaning up the LLM")
		err = client.Delete(ctx, llm)
		Expect(err).NotTo(HaveOccurred())
	})
})

var _ = Describe("Basic Integration Test", func() {
	var (
		ctx        context.Context
		client     client.Client
		namespace  = "default"
		uniqueID   string
		testSecret *TestSecret
		testLLM    *TestLLM
		testAgent  *TestAgent
		testTask   *TestTask
		mockServer *httptest.Server
	)

	BeforeEach(func() {
		// Initialize context and client from framework
		ctx = testFramework.GetContext()
		client = testFramework.GetClient()

		// Setup mock server for LLM API calls
		mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Always return success for our tests
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			// Return appropriate OpenAI-compatible response
			_, err := w.Write([]byte(`{"id":"test-id","choices":[{"message":{"content":"test"}}]}`))
			if err != nil {
				http.Error(w, "Error writing response", http.StatusInternalServerError)
				return
			}
		}))

		// Generate unique test ID for resource names
		uniqueID = fmt.Sprintf("test-%d", time.Now().UnixNano())

		// Setup test resources with unique names
		testSecret = &TestSecret{
			Name: fmt.Sprintf("test-secret-%s", uniqueID),
		}

		testLLM = &TestLLM{
			Name:       fmt.Sprintf("test-llm-%s", uniqueID),
			SecretName: testSecret.Name,
		}

		testAgent = &TestAgent{
			Name:         fmt.Sprintf("test-agent-%s", uniqueID),
			SystemPrompt: "You are a helpful test assistant.",
			LLM:          testLLM.Name,
		}

		testTask = &TestTask{
			Name:        fmt.Sprintf("test-task-%s", uniqueID),
			AgentName:   testAgent.Name,
			UserMessage: "What is the capital of France?",
		}
	})

	AfterEach(func() {
		// Clean up mock server
		if mockServer != nil {
			mockServer.Close()
		}

		// Clean up test resources
		if testTask != nil {
			testTask.Teardown(ctx)
		}
		if testAgent != nil {
			testAgent.Teardown(ctx)
		}
		if testLLM != nil {
			// Manual cleanup for LLM since we created it manually
			llmResource := &acp.LLM{
				ObjectMeta: metav1.ObjectMeta{
					Name:      testLLM.Name,
					Namespace: namespace,
				},
			}
			_ = client.Delete(ctx, llmResource)
		}
		if testSecret != nil {
			testSecret.Teardown(ctx)
		}
	})

	It("should successfully create and process LLM → Agent → Task flow", func() {
		By("creating test secret for LLM API key")
		secret := testSecret.Setup(ctx, client)
		Expect(secret).NotTo(BeNil())

		By("creating LLM resource with mock server URL")
		llm := &acp.LLM{
			ObjectMeta: metav1.ObjectMeta{
				Name:      testLLM.Name,
				Namespace: namespace,
			},
			Spec: acp.LLMSpec{
				Provider: "openai",
				APIKeyFrom: &acp.APIKeySource{
					SecretKeyRef: acp.SecretKeyRef{
						Name: testSecret.Name,
						Key:  "api-key",
					},
				},
				Parameters: acp.BaseConfig{
					BaseURL: mockServer.URL, // Use mock server URL
					Model:   "test-model",
				},
			},
		}
		err := client.Create(ctx, llm)
		Expect(err).NotTo(HaveOccurred())
		Expect(llm.Spec.Provider).To(Equal("openai"))

		By("waiting for LLM to be ready")
		Eventually(func(g Gomega) {
			updatedLLM := &acp.LLM{}
			err := client.Get(ctx, types.NamespacedName{
				Name:      llm.Name,
				Namespace: namespace,
			}, updatedLLM)
			g.Expect(err).NotTo(HaveOccurred())

			// Debug output
			fmt.Printf("LLM Status: Ready=%v, Status=%s, StatusDetail=%s\n",
				updatedLLM.Status.Ready, updatedLLM.Status.Status, updatedLLM.Status.StatusDetail)

			g.Expect(updatedLLM.Status.Ready).To(BeTrue())
		}).Should(Succeed())

		By("creating Agent resource")
		agent := testAgent.Setup(ctx, client)
		Expect(agent).NotTo(BeNil())
		Expect(agent.Spec.LLMRef.Name).To(Equal(testLLM.Name))
		Expect(agent.Spec.System).To(Equal("You are a helpful test assistant."))

		By("waiting for Agent to be ready")
		Eventually(func(g Gomega) {
			updatedAgent := &acp.Agent{}
			err := client.Get(ctx, types.NamespacedName{
				Name:      agent.Name,
				Namespace: namespace,
			}, updatedAgent)
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(updatedAgent.Status.Ready).To(BeTrue())
		}).Should(Succeed())

		By("creating Task resource")
		task := testTask.Setup(ctx, client)
		Expect(task).NotTo(BeNil())
		Expect(task.Spec.AgentRef.Name).To(Equal(agent.Name))
		Expect(task.Spec.UserMessage).To(Equal("What is the capital of France?"))

		By("waiting for Task to initialize")
		Eventually(func(g Gomega) {
			updatedTask := &acp.Task{}
			err := client.Get(ctx, types.NamespacedName{
				Name:      task.Name,
				Namespace: namespace,
			}, updatedTask)
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(updatedTask.Status.Phase).To(Equal(acp.TaskPhaseInitializing))
			g.Expect(updatedTask.Status.SpanContext).NotTo(BeNil())
		}).Should(Succeed())

		By("waiting for Task to be ready for LLM")
		Eventually(func(g Gomega) {
			updatedTask := &acp.Task{}
			err := client.Get(ctx, types.NamespacedName{
				Name:      task.Name,
				Namespace: namespace,
			}, updatedTask)
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(updatedTask.Status.Phase).To(Equal(acp.TaskPhaseReadyForLLM))
			g.Expect(updatedTask.Status.ContextWindow).To(HaveLen(2))
			g.Expect(updatedTask.Status.ContextWindow[0].Role).To(Equal("system"))
			g.Expect(updatedTask.Status.ContextWindow[1].Role).To(Equal("user"))
		}).Should(Succeed())

		By("verifying the complete flow worked end-to-end")
		finalTask := &acp.Task{}
		err = client.Get(ctx, types.NamespacedName{
			Name:      task.Name,
			Namespace: namespace,
		}, finalTask)
		Expect(err).NotTo(HaveOccurred())

		// Verify task progression
		Expect(finalTask.Status.Phase).To(Equal(acp.TaskPhaseReadyForLLM))
		Expect(finalTask.Status.ContextWindow).To(HaveLen(2))
		Expect(finalTask.Status.ContextWindow[0].Content).To(ContainSubstring("helpful test assistant"))
		Expect(finalTask.Status.ContextWindow[1].Content).To(ContainSubstring("capital of France"))
	})

	It("should handle task with missing agent gracefully", func() {
		By("creating a task with non-existent agent")
		task := &acp.Task{
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("orphan-task-%s", uniqueID),
				Namespace: namespace,
			},
			Spec: acp.TaskSpec{
				AgentRef: acp.LocalObjectReference{
					Name: "non-existent-agent",
				},
				UserMessage: "This should fail",
			},
		}
		err := client.Create(ctx, task)
		Expect(err).NotTo(HaveOccurred())

		By("waiting for task to be in pending state")
		Eventually(func(g Gomega) {
			updatedTask := &acp.Task{}
			err := client.Get(ctx, types.NamespacedName{
				Name:      task.Name,
				Namespace: namespace,
			}, updatedTask)
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(updatedTask.Status.Phase).To(Equal(acp.TaskPhasePending))
			g.Expect(updatedTask.Status.StatusDetail).To(ContainSubstring("Waiting for Agent to exist"))
		}).Should(Succeed())

		By("cleaning up orphan task")
		err = client.Delete(ctx, task)
		Expect(err).NotTo(HaveOccurred())
	})
})
