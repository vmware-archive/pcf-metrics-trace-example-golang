package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pivotal-cf/pcf-metrics-trace-example-golang/handlers"
	"github.com/tedsuo/rata"
)

func main() {
	routes := rata.Routes{
		{Name: "ViewItem", Method: rata.GET, Path: "/view-item/:delay"},
		{Name: "ItemInfo", Method: rata.GET, Path: "/item-info/:delay"},
		{Name: "ItemReview", Method: rata.GET, Path: "/item-review/:delay"},
	}

	targetUrl := fmt.Sprintf("%s.%s", os.Getenv("TARGET_APP"), os.Getenv("TARGET_DOMAIN"))

	routeHandlers := rata.Handlers{
		"ViewItem":   handlers.NewProxyMiddleware([]string{"item-info", "item-review"}, targetUrl, handlers.NewTraceHandler()),
		"ItemInfo":   handlers.NewTraceHandler(),
		"ItemReview": handlers.NewTraceHandler(),
	}

	router, err := rata.NewRouter(routes, routeHandlers)
	if err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), router); err != nil {
		panic(err)
	}
}
