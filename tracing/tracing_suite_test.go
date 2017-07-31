package tracing_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTracing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tracing Suite")
}
