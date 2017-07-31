package logging_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/pcf-metrics-trace-example-golang/tracing"
	"github.com/pivotal-cf/pcf-metrics-trace-example-golang/logging"
	"bytes"
)

var _ = Describe("Logging", func() {
	var logOutput bytes.Buffer

	BeforeEach(func() {
		logging.LogToWriter(&logOutput)
	})

	Describe("Println()", func() {
		It("prints the desired text with tracing information", func() {
			traceIds := &tracing.Ids{
				TraceId: "some-trace-id",
				SpanId: "some-span-id",
				ParentSpanId: "some-parent-span-id",
			}

			logger := logging.NewLogger(traceIds)

			logger.Println("some message")
			Expect(logOutput.String()).To(ContainSubstring("some message"))
			Expect(logOutput.String()).To(ContainSubstring("[some-parent-span-id,some-trace-id,some-span-id,true]"))
		})
	})

	Describe("Printf()", func() {
		It("formats and prints the desired text with tracing information", func() {
			traceIds := &tracing.Ids{
				TraceId: "some-trace-id",
				SpanId: "some-span-id",
				ParentSpanId: "some-parent-span-id",
			}

			logger := logging.NewLogger(traceIds)

			logger.Printf("some message %v", 47)
			Expect(logOutput.String()).To(ContainSubstring("some message 47"))
			Expect(logOutput.String()).To(ContainSubstring("[some-parent-span-id,some-trace-id,some-span-id,true]"))
		})
	})
})
