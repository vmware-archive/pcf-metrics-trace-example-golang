package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/tedsuo/rata"
)

var (
	traceIdHeader = "X-B3-TraceId"
	spanIdHeader  = "X-B3-SpanId"
)

type proxyMiddleware struct {
	reqs      []string
	targetUrl string
	next      http.Handler
}

func NewProxyMiddleware(proxyReqs []string, targetUrl string, next http.Handler) http.Handler {
	return proxyMiddleware{
		reqs:      proxyReqs,
		targetUrl: targetUrl,
		next:      next,
	}
}

func (p proxyMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	delayParam := rata.Param(r, "delay")
	delay, err := strconv.ParseFloat(delayParam, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, path := range p.reqs {
		err := p.proxy(path, delay, r)
		if err != nil {
			fmt.Printf("Failed to forward request %s", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Failed to forward request"))
			return
		}
	}

	p.next.ServeHTTP(w, r)
}

func (p proxyMiddleware) proxy(path string, delay float64, proxyReq *http.Request) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/%s/%f", p.targetUrl, path, delay), nil)

	if err != nil {
		fmt.Printf("Cannot create proxy request %s", err.Error())
		return err
	}

	traceId := proxyReq.Header.Get(traceIdHeader)
	spanId := proxyReq.Header.Get(spanIdHeader)

	req.Header.Add(traceIdHeader, traceId)
	req.Header.Add(spanIdHeader, spanId)

	client := &http.Client{
		Timeout: time.Duration(30 * time.Second),
	}

	_, err = client.Do(req)

	return err
}
