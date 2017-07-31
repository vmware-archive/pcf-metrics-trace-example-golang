package tracing_test

import (
	"github.com/pivotal-cf/pcf-metrics-trace-example-golang/tracing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
)

var _ = Describe("Ids", func() {
	Describe("NewIds()", func() {
		It("converts http headers into trace ids", func() {
			headers := http.Header{}
			headers.Add("X-B3-TraceId", "some-trace-id")
			headers.Add("X-B3-SpanId", "some-span-id")
			headers.Add("X-B3-ParentSpanId", "some-parent-span-id")

			ids := tracing.NewIds(headers)
			Expect(ids.TraceId).To(Equal("some-trace-id"))
			Expect(ids.SpanId).To(Equal("some-span-id"))
			Expect(ids.ParentSpanId).To(Equal("some-parent-span-id"))
		})
	})
})
