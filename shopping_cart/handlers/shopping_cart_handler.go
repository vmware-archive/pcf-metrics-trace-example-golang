package handlers

import (
	"net/http"
	"github.com/pivotal-cf/pcf-metrics-trace-example-golang/tracing"
	"github.com/pivotal-cf/pcf-metrics-trace-example-golang/logging"
)

type shoppingCartHandler struct{}

func NewShoppingCartHandler() shoppingCartHandler {
	return shoppingCartHandler{}
}

func (t shoppingCartHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	traceIds := tracing.NewIds(r.Header)
	logger := logging.NewLogger(traceIds)

	logger.Println("Added items to the shopping cart")
}
