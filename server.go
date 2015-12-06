package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/asvins/common_db/postgres"
	"github.com/asvins/common_io"
	"github.com/asvins/utils/config"
	"github.com/jinzhu/gorm"
	"github.com/unrolled/render"
)

var (
	ServerConfig *Config        = new(Config)
	rend         *render.Render = render.New() // used to write into responses
	db           *gorm.DB
	producer     *common_io.Producer
	consumer     *common_io.Consumer
)

// function that will run before main
func init() {
	fmt.Println("[INFO] Initializing server")
	err := config.Load("warehouse_config.gcfg", ServerConfig)
	if err != nil {
		log.Fatal(err)
	}

	DatabaseConfig := postgres.NewConfig(ServerConfig.Database.User, ServerConfig.Database.DbName, ServerConfig.Database.SSLMode)
	db = postgres.GetDatabase(DatabaseConfig)
	fmt.Println("[INFO] Initialization Done!")

	/*
	*	Common io
	 */
	setupCommonIo()
}

func main() {
	router := DefRoutes()

	fmt.Println("[INFO] Server running on port:", ServerConfig.Server.Port)
	http.ListenAndServe(":"+ServerConfig.Server.Port, router)
}
