package main

import (
	"net/http"

	"github.com/asvins/router"
	"github.com/asvins/router/errors"
	"github.com/asvins/router/logger"
)

func DefRoutes() *router.Router {
	r := router.NewRouter()

	// discovery
	r.Handle("/api/discovery", router.GET, func(w http.ResponseWriter, rq *http.Request) errors.Http {
		discoveryMap := make(map[string]map[string]string)

		// product
		discoveryMap["retreive_product"] = map[string]string{"GET": "/api/inventory/product"}
		discoveryMap["retreive_product_by_id"] = map[string]string{"GET": "/api/inventory/product/:id"}
		discoveryMap["insert_product"] = map[string]string{"POST": "/api/inventory/product"}
		discoveryMap["update_product"] = map[string]string{"PUT": "/api/inventory/product/:id"}
		discoveryMap["delete_product"] = map[string]string{"DELETE": "/api/inventory/product/:id"}
		discoveryMap["consume_product"] = map[string]string{"GET": "/api/inventory/product/:id/consume/:quantity"}

		// order
		discoveryMap["retreive_order"] = map[string]string{"GET": "/api/inventory/order"}
		discoveryMap["retreive_open_order"] = map[string]string{"GET": "/api/inventory/order/open"}
		discoveryMap["retreive_order_by_id"] = map[string]string{"GET": "/api/inventory/order/:id"}
		discoveryMap["approve_order"] = map[string]string{"PUT": "/api/inventory/order/:id/approve"}
		discoveryMap["cancel_order"] = map[string]string{"PUT": "/api/inventory/order/:id/cancel"}

		// purchase
		discoveryMap["retreive_purchase"] = map[string]string{"GET": "/api/inventory/purchase"}
		discoveryMap["retreive_purchase_by_id"] = map[string]string{"GET": "/api/inventory/purchase/:id"}
		discoveryMap["confirm_purchase"] = map[string]string{"PUT": "/api/inventory/purchase/:id/confirm"}
		discoveryMap["conclude_purchase"] = map[string]string{"PUT": "/api/inventory/purchase/:id/conclude"}

		// purchase products
		discoveryMap["retreive_purchase_product"] = map[string]string{"GET": "/api/inventory/purchaseProduct"}
		discoveryMap["retreive_purcahse_product_by_product_id"] = map[string]string{"GET": "/api/inventory/purchaseProduct/product/:product_id"}
		discoveryMap["retreive_purchase_product_by_id"] = map[string]string{"GET": "/api/inventory/purchaseProduct/:id"}
		discoveryMap["update_purchase_product_quantity"] = map[string]string{"PUT": "/api/inventory/purchaseProduct/:id/updateQuantity/:quantity"}
		discoveryMap["update_purchase_product_value"] = map[string]string{"PUT": "/api/inventory/purchaseProduct/:id/updateValue/:value"}

		// withdrawal
		discoveryMap["retreive_withdrawl"] = map[string]string{"GET": "/api/inventory/withdrawal"}

		rend.JSON(w, http.StatusOK, discoveryMap)
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
	r.Handle("/api/inventory/purchaseProduct", router.GET, retreivePurchaseProducts, []router.Interceptor{})
	r.Handle("/api/inventory/purchaseProduct/product/:product_id", router.GET, retreivePurchaseProductsByProductId, []router.Interceptor{})
	r.Handle("/api/inventory/purchaseProduct/:id", router.GET, retreivePurchaseProductsById, []router.Interceptor{})
	r.Handle("/api/inventory/purchaseProduct/:id/updateQuantity/:quantity", router.PUT, updatePurchaseProductOnQuantity, []router.Interceptor{})
	r.Handle("/api/inventory/purchaseProduct/:id/updateValue/:value", router.PUT, updatePurchaseProductOnValue, []router.Interceptor{})

	// withdrawal
	r.Handle("/api/inventory/withdrawal", router.GET, retreiveWithdrawal, []router.Interceptor{})

	// interceptors
	r.AddBaseInterceptor("/", logger.NewLogger())

	return r
}
