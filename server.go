package main

import (
	"fmt"
	"net/http"

	"github.com/asvins/warehouse/interceptors"
	"github.com/rcmgleite/router"
)

func main() {
	r := router.NewRouter()

	// index - maybe discovery?
	r.AddRoute("/", router.GET, func(w http.ResponseWriter, apiRouter *http.Request) {
		fmt.Fprint(w, "Request made to '/'")
	})

	// product routes
	r.AddRoute("/api/product", router.GET, GETProductHandler)
	r.AddRoute("/api/product", router.POST, POSTProductHandler)
	r.AddRoute("/api/product", router.PUT, PUTProductHandler)
	r.AddRoute("/api/product", router.DELETE, DELETEProductHandler)

	//order routes
	// - TODO

	// interceptors
	r.AddBaseInterceptor("/", &interceptors.Logger{})

	fmt.Println("[INFO] Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
