package handlers

import (
	"net/http"
	"github.com/pivotal-cf/pcf-metrics-trace-example-golang/tracing"
	"github.com/pivotal-cf/pcf-metrics-trace-example-golang/logging"
)

type paymentsHandler struct{}

func NewPaymentsHandler() paymentsHandler {
	return paymentsHandler{}
}

func (t paymentsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	traceIds := tracing.NewIds(r.Header)
	logger := logging.NewLogger(traceIds)

	logger.Println("Card successfully charged!")
	w.Write([]byte("Card successfully charged!"))
}
