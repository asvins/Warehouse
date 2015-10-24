package main

import (
	"fmt"
	"net/http"

	"log"

	"github.com/asvins/common_interceptors/logger"
	"github.com/asvins/utils/config"
	"github.com/rcmgleite/router"
)

func main() {
	r := router.NewRouter()

	// index - maybe discovery?
	r.AddRoute("/", router.GET, func(w http.ResponseWriter, apiRouter *http.Request) {
		fmt.Fprint(w, "Request made to '/'")
	})

	// INVENTORY
	// product routes
	r.AddRoute("/api/inventory/product", router.GET, GETProductHandler)
	r.AddRoute("/api/inventory/product", router.POST, POSTProductHandler)
	r.AddRoute("/api/inventory/product", router.PUT, PUTProductHandler)
	r.AddRoute("/api/inventory/product", router.DELETE, DELETEProductHandler)

	//order routes
	r.AddRoute("/api/inventory/order", router.GET, GETOrderHandler)

	// SHOP/ACQUIRING

	// interceptors
	r.AddBaseInterceptor("/", logger.NewLogger())

	serverConfig := Config{}
	err := config.Load("warehouse_config.gcfg", &serverConfig)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[INFO] Server running on port:", serverConfig.Server.Port)
	http.ListenAndServe(":"+serverConfig.Server.Port, r)
}
