package main

import (
	"net/http"
	"os"

	. "github.com/pivotal-cf/pcf-metrics-trace-example-golang/middlewares"
	. "github.com/pivotal-cf/pcf-metrics-trace-example-golang/orders/handlers"
	"github.com/tedsuo/rata"
)

func main() {
	routes := rata.Routes{
		{Name: "ProcessOrder", Method: rata.GET, Path: "/process-order"},
	}

	routeHandlers := rata.Handlers{
		"ProcessOrder": NewProxyMiddleware("charge-card", os.Getenv("PAYMENTS_HOST"), NewOrdersHandler()),
	}

	router, err := rata.NewRouter(routes, routeHandlers)
	if err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), router); err != nil {
		panic(err)
	}
}
