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

package getting_started

import (
	"context"
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	acp "github.com/humanlayer/agentcontrolplane/acp/api/v1alpha1"
)

var _ = Describe("MCP Transport Tests", func() {
	var (
		ctx       context.Context
		client    client.Client
		namespace = "default"
		uniqueID  string
	)

	BeforeEach(func() {
		ctx = testFramework.GetContext()
		client = testFramework.GetClient()
		uniqueID = fmt.Sprintf("mcp-test-%d", time.Now().UnixNano())
	})

	It("should create MCPServer with stdio transport", func() {
		By("creating MCPServer with stdio transport")
		mcpServer := &acp.MCPServer{
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("%s-stdio", uniqueID),
				Namespace: namespace,
			},
			Spec: acp.MCPServerSpec{
				Transport: "stdio",
				Command:   "python",
				Args:      []string{"-m", "mcp_server", "--stdio"},
			},
		}

		err := client.Create(ctx, mcpServer)
		Expect(err).NotTo(HaveOccurred())

		By("verifying the MCPServer was created")
		createdMCP := &acp.MCPServer{}
		err = client.Get(ctx, types.NamespacedName{
			Name:      mcpServer.Name,
			Namespace: namespace,
		}, createdMCP)
		Expect(err).NotTo(HaveOccurred())
		Expect(createdMCP.Spec.Transport).To(Equal("stdio"))
		Expect(createdMCP.Spec.Command).To(Equal("python"))

		By("cleaning up the MCPServer")
		err = client.Delete(ctx, mcpServer)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should create MCPServer with HTTP transport", func() {
		By("creating MCPServer with HTTP transport")
		mcpServer := &acp.MCPServer{
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("%s-http", uniqueID),
				Namespace: namespace,
			},
			Spec: acp.MCPServerSpec{
				Transport: "http",
				URL:       "http://localhost:3000/mcp",
			},
		}

		err := client.Create(ctx, mcpServer)
		Expect(err).NotTo(HaveOccurred())

		By("verifying the MCPServer was created")
		createdMCP := &acp.MCPServer{}
		err = client.Get(ctx, types.NamespacedName{
			Name:      mcpServer.Name,
			Namespace: namespace,
		}, createdMCP)
		Expect(err).NotTo(HaveOccurred())
		Expect(createdMCP.Spec.Transport).To(Equal("http"))
		Expect(createdMCP.Spec.URL).To(Equal("http://localhost:3000/mcp"))

		By("cleaning up the MCPServer")
		err = client.Delete(ctx, mcpServer)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should create MCPServer with SSE transport", func() {
		By("creating MCPServer with SSE transport")
		mcpServer := &acp.MCPServer{
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("%s-sse", uniqueID),
				Namespace: namespace,
			},
			Spec: acp.MCPServerSpec{
				Transport: "sse",
				URL:       "http://localhost:3000/sse",
				Headers: map[string]string{
					"Authorization": "Bearer token123",
					"X-Custom":      "value",
				},
			},
		}

		err := client.Create(ctx, mcpServer)
		Expect(err).NotTo(HaveOccurred())

		By("verifying the MCPServer was created with headers")
		createdMCP := &acp.MCPServer{}
		err = client.Get(ctx, types.NamespacedName{
			Name:      mcpServer.Name,
			Namespace: namespace,
		}, createdMCP)
		Expect(err).NotTo(HaveOccurred())
		Expect(createdMCP.Spec.Transport).To(Equal("sse"))
		Expect(createdMCP.Spec.URL).To(Equal("http://localhost:3000/sse"))
		Expect(createdMCP.Spec.Headers).To(HaveKey("Authorization"))
		Expect(createdMCP.Spec.Headers["Authorization"]).To(Equal("Bearer token123"))
		Expect(createdMCP.Spec.Headers["X-Custom"]).To(Equal("value"))

		By("cleaning up the MCPServer")
		err = client.Delete(ctx, mcpServer)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should create MCPServer with streamable-http transport", func() {
		By("creating MCPServer with streamable-http transport")
		mcpServer := &acp.MCPServer{
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("%s-streamable", uniqueID),
				Namespace: namespace,
			},
			Spec: acp.MCPServerSpec{
				Transport: "streamable-http",
				URL:       "http://localhost:3000/stream",
				SessionID: "session-12345",
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			},
		}

		err := client.Create(ctx, mcpServer)
		Expect(err).NotTo(HaveOccurred())

		By("verifying the MCPServer was created with SessionId")
		createdMCP := &acp.MCPServer{}
		err = client.Get(ctx, types.NamespacedName{
			Name:      mcpServer.Name,
			Namespace: namespace,
		}, createdMCP)
		Expect(err).NotTo(HaveOccurred())
		Expect(createdMCP.Spec.Transport).To(Equal("streamable-http"))
		Expect(createdMCP.Spec.URL).To(Equal("http://localhost:3000/stream"))
		Expect(createdMCP.Spec.SessionID).To(Equal("session-12345"))
		Expect(createdMCP.Spec.Headers).To(HaveKey("Content-Type"))

		By("cleaning up the MCPServer")
		err = client.Delete(ctx, mcpServer)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should validate transport field requirements", func() {
		By("testing that transport field is required")
		mcpServer := &acp.MCPServer{
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("%s-invalid", uniqueID),
				Namespace: namespace,
			},
			Spec: acp.MCPServerSpec{
				// Missing transport field - should fail validation
				URL: "http://localhost:3000",
			},
		}

		err := client.Create(ctx, mcpServer)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("transport"))
	})
})
