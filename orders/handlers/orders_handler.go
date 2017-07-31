package handlers

import (
	"net/http"
	"github.com/pivotal-cf/pcf-metrics-trace-example-golang/tracing"
	"github.com/pivotal-cf/pcf-metrics-trace-example-golang/logging"
)

type ordersHandler struct{}

func NewOrdersHandler() ordersHandler {
	return ordersHandler{}
}

func (t ordersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	traceIds := tracing.NewIds(r.Header)
	logger := logging.NewLogger(traceIds)

	logger.Println("Order successfully placed!")
	w.Write([]byte("Order successfully placed!"))
}
