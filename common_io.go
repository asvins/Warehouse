package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/asvins/common_io"
	"github.com/asvins/utils/config"
	"github.com/asvins/warehouse/models"
)

func setupCommonIo() {
	cfg := common_io.Config{}

	err := config.Load("common_io_config.gcfg", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	/*
	*	Producer
	 */
	producer, err = common_io.NewProducer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	/*
	*	Consumer
	 */
	consumer = common_io.NewConsumer(cfg)
	consumer.HandleTopic("product_created", handleProductCreated)

	if err = consumer.StartListening(); err != nil {
		log.Fatal(err.Error())
	}
}

/*
*	Handlers
 */
func handleProductCreated(msg []byte) {
	fmt.Println("[INFO] Received Kafka message from topic 'product_created'")
	p := models.Product{}
	if err := json.Unmarshal(msg, &p); err != nil {
		fmt.Println("[ERROR] Unable to Unmarshal json from message 'product_created'", err.Error())
		return
	}
	p.CurrQuantity = 100000
	p.MinQuantity = 90

	rand.Seed(time.Now().UTC().UnixNano())
	p.CurrentValue = float64(rand.Intn(10)) / 10.0

	if err := p.Save(db); err != nil {
		producer.Publish("product_created_dead_letter", msg)
	}
}
