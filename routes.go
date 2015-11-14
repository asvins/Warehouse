package main

import (
	"fmt"
	"net/http"

	"github.com/asvins/router"
	"github.com/asvins/router/errors"
	"github.com/asvins/router/logger"
)

func DefRoutes() *router.Router {
	r := router.NewRouter()

	// index - maybe discovery?
	r.Handle("/", router.GET, func(w http.ResponseWriter, apiRouter *http.Request) errors.Http {
		fmt.Fprint(w, "Request made to '/'")
		return nil
	}, []router.Interceptor{})

	// INVENTORY
	// product routes
	r.Handle("/api/inventory/product", router.GET, retreiveProduct, []router.Interceptor{})
	r.Handle("/api/inventory/product/:id", router.GET, retreiveProductById, []router.Interceptor{})
	r.Handle("/api/inventory/product", router.POST, insertProduct, []router.Interceptor{})
	r.Handle("/api/inventory/product/:id", router.PUT, updateProduct, []router.Interceptor{})
	r.Handle("/api/inventory/product", router.DELETE, deleteProduct, []router.Interceptor{})
	r.Handle("/api/inventory/product/:id/consume/:quantity", router.GET, consumeProduct, []router.Interceptor{})

	//order routes
	//r.Handle("/api/inventory/order", router.GET, GETOrderHandler, []router.Interceptor{})

	// interceptors
	r.AddBaseInterceptor("/", logger.NewLogger())

	return r
}
