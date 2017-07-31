package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/pivotal-cf/pcf-metrics-trace-example-golang/middlewares"
)

var _ = Describe("Proxy Middleware", func() {
	var (
		previousTraceId = "abc123"
		previousSpanId  = "def123"
		nextHandler     mockHandler
		destHandler     mockHandler
		mockProxyServer *httptest.Server
		mockDestServer  *httptest.Server
	)

	BeforeEach(func() {
		nextHandler = newMockHandler()

		destHandler = newMockHandler()
		mockDestServer = httptest.NewServer(destHandler)

		destUrl := mockDestServer.URL
		mockProxyServer = httptest.NewServer(middlewares.NewProxyMiddleware("dest", destUrl[7:], nextHandler))

		req, err := http.NewRequest("GET", mockProxyServer.URL, nil)
		Expect(err).ToNot(HaveOccurred())

		req.Header.Add("X-B3-TraceId", previousTraceId)
		req.Header.Add("X-B3-SpanId", previousSpanId)

		client := &http.Client{
			Timeout: time.Duration(30 * time.Second),
		}
		_, err = client.Do(req)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		mockProxyServer.Close()
		mockDestServer.Close()
	})

	It("sends the correct headers", func() {
		Eventually(destHandler.requests).Should(HaveLen(1))

		var receivedTraceId string
		Eventually(destHandler.traceId).Should(HaveLen(1))
		Eventually(destHandler.traceId).Should(Receive(&receivedTraceId))
		Expect(receivedTraceId).To(Equal(previousTraceId))

		var receivedSpanId string
		Eventually(destHandler.spanId).Should(HaveLen(1))
		Eventually(destHandler.spanId).Should(Receive(&receivedSpanId))
		Expect(receivedSpanId).NotTo(Equal(previousSpanId))
		Expect(receivedSpanId).NotTo(Equal(""))

		var receivedParentSpanId string
		Eventually(destHandler.parentSpanId).Should(HaveLen(1))
		Eventually(destHandler.parentSpanId).Should(Receive(&receivedParentSpanId))
		Expect(receivedParentSpanId).To(Equal(previousSpanId))
	})

	It("calls next handler", func() {
		Eventually(nextHandler.requests).Should(HaveLen(1))
	})
})

type mockHandler struct {
	requests     chan int
	traceId      chan string
	spanId       chan string
	parentSpanId chan string
}

func newMockHandler() mockHandler {
	requests := make(chan int, 5)
	spanId := make(chan string, 5)
	parentSpanId := make(chan string, 5)
	traceId := make(chan string, 5)
	return mockHandler{
		requests:     requests,
		traceId:      traceId,
		spanId:       spanId,
		parentSpanId: parentSpanId,
	}
}

func (m mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.requests <- 1
	m.traceId <- r.Header.Get("X-B3-TraceId")
	m.spanId <- r.Header.Get("X-B3-SpanId")
	m.parentSpanId <- r.Header.Get("X-B3-ParentSpanId")
}
