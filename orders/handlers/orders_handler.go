package handlers

import "net/http"

type ordersHandler struct{}

func NewOrdersHandler() ordersHandler {
	return ordersHandler{}
}

func (t ordersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Order successfully placed!"))
}
