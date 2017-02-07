package main

import (
	"net/http"
	"os"

	"github.com/pivotal-cf/pcf-metrics-trace-example-golang/payments/handlers"
	"github.com/tedsuo/rata"
)

func main() {
	routes := rata.Routes{
		{Name: "ChargeCard", Method: rata.GET, Path: "/charge-card"},
	}

	routeHandlers := rata.Handlers{
		"ChargeCard": handlers.NewPaymentsHandler(),
	}

	router, err := rata.NewRouter(routes, routeHandlers)
	if err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), router); err != nil {
		panic(err)
	}
}
