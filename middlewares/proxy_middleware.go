package middlewares

import (
	"fmt"
	"net/http"
	"time"
	"math/rand"
	"github.com/pivotal-cf/pcf-metrics-trace-example-golang/logging"
	"github.com/pivotal-cf/pcf-metrics-trace-example-golang/tracing"
)

type proxyMiddleware struct {
	path      string
	targetUrl string
	next      http.Handler
}

func NewProxyMiddleware(path, targetUrl string, next http.Handler) http.Handler {
	return proxyMiddleware{
		path:      path,
		targetUrl: targetUrl,
		next:      next,
	}
}

func (p proxyMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	traceIds := tracing.NewIds(r.Header)
	logger := logging.NewLogger(traceIds)

	err := p.proxy(r, traceIds)
	if err != nil {
		logger.Printf(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to forward request"))
		return
	}

	p.next.ServeHTTP(w, r)
}

func (p proxyMiddleware) proxy(proxyReq *http.Request, traceIds *tracing.Ids) error {
	time.Sleep(time.Duration(500) * time.Millisecond)

	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/%s", p.targetUrl, p.path), nil)
	if err != nil {
		return fmt.Errorf("Cannot create proxy request: %s", err.Error())
	}

	previousSpanId := traceIds.SpanId

	req.Header.Add(tracing.TraceIdHeader, traceIds.TraceId)
	req.Header.Add(tracing.ParentSpanIdHeader, previousSpanId)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	req.Header.Add(tracing.SpanIdHeader, fmt.Sprintf("%x", r1.Uint64()))

	client := &http.Client{
		Timeout: time.Duration(30 * time.Second),
	}

	_, err = client.Do(req)
	if err != nil {
		return fmt.Errorf("Failed to forward request: %s", err.Error())
	}

	return nil
}
