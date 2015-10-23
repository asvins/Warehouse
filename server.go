package main

import (
	"fmt"
	"net/http"

	"github.com/asvins/warehouse/interceptors"
	"github.com/asvins/warehouse/inventory"
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
	r.AddRoute("/api/inventory/product", router.GET, inventory.GETProductHandler)
	r.AddRoute("/api/inventory/product", router.POST, inventory.POSTProductHandler)
	r.AddRoute("/api/inventory/product", router.PUT, inventory.PUTProductHandler)
	r.AddRoute("/api/inventory/product", router.DELETE, inventory.DELETEProductHandler)

	//order routes
	r.AddRoute("/api/inventory/order", router.GET, inventory.GETOrderHandler)

	// SHOP/ACQUIRING

	// interceptors
	r.AddBaseInterceptor("/", &interceptors.Logger{})

	fmt.Println("[INFO] Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
