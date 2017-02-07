package main

import (
	"net/http"
	"os"

	. "github.com/pivotal-cf/pcf-metrics-trace-example-golang/middlewares"
	. "github.com/pivotal-cf/pcf-metrics-trace-example-golang/shopping_cart/handlers"
	"github.com/tedsuo/rata"
)

func main() {
	routes := rata.Routes{
		{Name: "CheckOut", Method: rata.GET, Path: "/checkout"},
	}

	routeHandlers := rata.Handlers{
		"CheckOut": NewProxyMiddleware("process-order", os.Getenv("ORDERS_HOST"), NewShoppingCartHandler()),
	}

	router, err := rata.NewRouter(routes, routeHandlers)
	if err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), router); err != nil {
		panic(err)
	}
}
