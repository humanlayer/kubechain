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
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var (
	testFramework *TestFramework
)

func TestE2EFramework(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2E Framework Suite")
}

var _ = BeforeSuite(func() {
	logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))

	By("creating test framework")
	testFramework = NewTestFramework()

	By("starting test framework with all controllers")
	err := testFramework.Start()
	Expect(err).NotTo(HaveOccurred())

	// Configure timeouts for integration tests
	SetDefaultEventuallyTimeout(30 * time.Second)
	SetDefaultEventuallyPollingInterval(1 * time.Second)
	SetDefaultConsistentlyDuration(5 * time.Second)
	SetDefaultConsistentlyPollingInterval(200 * time.Millisecond)
})

var _ = AfterSuite(func() {
	By("stopping test framework")
	if testFramework != nil {
		err := testFramework.Stop()
		Expect(err).NotTo(HaveOccurred())
	}
})
