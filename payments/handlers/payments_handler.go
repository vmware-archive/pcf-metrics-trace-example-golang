package handlers

import "net/http"

type paymentsHandler struct{}

func NewPaymentsHandler() paymentsHandler {
	return paymentsHandler{}
}

func (t paymentsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("card successfully charged!"))
}
