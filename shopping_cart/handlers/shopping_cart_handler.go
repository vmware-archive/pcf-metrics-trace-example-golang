package handlers

import "net/http"

type shoppingCartHandler struct{}

func NewShoppingCartHandler() shoppingCartHandler {
	return shoppingCartHandler{}
}

func (t shoppingCartHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
