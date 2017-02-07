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
		traceId         = "abc123"
		spanId          = "def123"
		proxyHandler    mockHandler
		destHandler     mockHandler
		mockProxyServer *httptest.Server
		mockDestServer  *httptest.Server
	)

	BeforeEach(func() {
		proxyHandler = newMockHandler()

		destHandler = newMockHandler()
		mockDestServer = httptest.NewServer(destHandler)

		destUrl := mockDestServer.URL
		mockProxyServer = httptest.NewServer(middlewares.NewProxyMiddleware("dest", destUrl[7:len(destUrl)], proxyHandler))

		req, err := http.NewRequest("GET", mockProxyServer.URL, nil)
		Expect(err).ToNot(HaveOccurred())

		req.Header.Add("X-B3-TraceId", traceId)
		req.Header.Add("X-B3-SpanId", spanId)

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

	It("requests destination", func() {
		Expect(destHandler.requests).To(HaveLen(1))
		Eventually(destHandler.traceId).Should(HaveLen(1))
		Eventually(destHandler.spanId).Should(HaveLen(1))
	})

	It("calles next handler", func() {
		Eventually(proxyHandler.requests).Should(HaveLen(1))
	})
})

type mockHandler struct {
	requests chan int
	traceId  chan string
	spanId   chan string
}

func newMockHandler() mockHandler {
	requests := make(chan int, 5)
	spanId := make(chan string, 5)
	traceId := make(chan string, 5)
	return mockHandler{
		requests: requests,
		traceId:  traceId,
		spanId:   spanId,
	}
}

func (m mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.requests <- 1
	m.traceId <- r.Header.Get("X-B3-TraceId")
	m.spanId <- r.Header.Get("X-B3-SpanId")
}
