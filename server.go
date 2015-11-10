package main

import (
	"fmt"
	"net/http"

	"log"

	"github.com/asvins/common_db/postgres"
	"github.com/asvins/utils/config"
	"github.com/unrolled/render"
)

var ServerConfig *Config = new(Config)
var DatabaseConfig *postgres.Config
var rend *render.Render = render.New() // used to write into responses

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
	router := DefRoutes()

	fmt.Println("[INFO] Server running on port:", ServerConfig.Server.Port)
	http.ListenAndServe(":"+ServerConfig.Server.Port, router)
}
