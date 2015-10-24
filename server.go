package main

import (
	"fmt"
	"net/http"

	"log"

	"github.com/asvins/common_db/postgres"
	"github.com/asvins/common_interceptors/logger"
	"github.com/asvins/utils/config"
	"github.com/rcmgleite/router"
)

var ServerConfig *Config = new(Config)
var DatabaseConfig *postgres.Config

// function that will run before main
func init() {
	fmt.Println("[INFO] Initializing server")
	err := config.Load("warehouse_config.gcfg", ServerConfig)
	if err != nil {
		log.Fatal(err)
	}

	DatabaseConfig = postgres.NewConfig(ServerConfig.Database.User, ServerConfig.Database.DbName, ServerConfig.Database.SSLMode)
	fmt.Println("[INFO] Initialization Done!")
}

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

	fmt.Println("[INFO] Server running on port:", ServerConfig.Server.Port)
	http.ListenAndServe(":"+ServerConfig.Server.Port, r)
}
