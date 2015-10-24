package main

import (
	"fmt"
	"net/http"

	"github.com/asvins/common_interceptors/logger"
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

	fmt.Println("[INFO] Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
