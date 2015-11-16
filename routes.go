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

	// product routes
	r.Handle("/api/inventory/product", router.GET, retreiveProduct, []router.Interceptor{})
	r.Handle("/api/inventory/product/:id", router.GET, retreiveProductById, []router.Interceptor{})
	r.Handle("/api/inventory/product", router.POST, insertProduct, []router.Interceptor{})
	r.Handle("/api/inventory/product/:id", router.PUT, updateProduct, []router.Interceptor{})
	r.Handle("/api/inventory/product/:id", router.DELETE, deleteProduct, []router.Interceptor{})
	r.Handle("/api/inventory/product/:id/consume/:quantity", router.GET, consumeProduct, []router.Interceptor{})

	// order routes
	r.Handle("/api/inventory/order", router.GET, retreiveOrder, []router.Interceptor{})
	r.Handle("/api/inventory/order/open", router.GET, retreiveOpenOrder, []router.Interceptor{})
	r.Handle("/api/inventory/order/:id", router.GET, retreiveOrderById, []router.Interceptor{})
	r.Handle("/api/inventory/order/:id/approve", router.PUT, approveOrder, []router.Interceptor{})
	r.Handle("/api/inventory/order/:id/cancel", router.PUT, cancelOrder, []router.Interceptor{})

	// purchase
	r.Handle("/api/inventory/purchase", router.GET, retreivePurchase, []router.Interceptor{})
	r.Handle("/api/inventory/purchase/:id", router.GET, retreivePurchaseById, []router.Interceptor{})
	r.Handle("/api/inventory/purchase/:id/confirm", router.PUT, confirmPurchase, []router.Interceptor{})
	r.Handle("/api/inventory/purchase/:id/conclude", router.PUT, concludePurchase, []router.Interceptor{})

	// purchase products
	r.Handle("/api/inventory/purchaseProducts", router.GET, retreivePurchaseProducts, []router.Interceptor{})
	r.Handle("/api/inventory/purchaseProducts/:product_id", router.GET, retreivePurchaseProductsByProductId, []router.Interceptor{})
	r.Handle("/api/inventory/purchaseProducts/:id", router.GET, retreivePurchaseProductsById, []router.Interceptor{})
	r.Handle("/api/inventory/purchaseProducts/:id/updateQuantity/:quantity", router.PUT, updatePurchaseProductOnQuantity, []router.Interceptor{})
	r.Handle("/api/inventory/purchaseProducts/:id/updateValue/:quantity", router.PUT, updatePurchaseProductOnValue, []router.Interceptor{})

	// withdrawal
	r.Handle("/api/inventory/withdrawal", router.GET, retreiveWithdrawal, []router.Interceptor{})

	// interceptors
	r.AddBaseInterceptor("/", logger.NewLogger())

	return r
}
