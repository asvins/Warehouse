package main

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
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

	/*
	*	Dead letters
	 */
	initDeadLettersReader()

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

	if err := p.Save(db); err != nil {
		producer.Publish("product_created_dead_letter", msg)
	}
}

/*
*	Dead Letter
 */
func handleProductCreatedDeadLetter(done chan bool) func(msg []byte) {
	stillWorking := make(chan bool)

	go func() {
		for {
			select {
			case <-stillWorking:
				fmt.Println("[INFO] DeadLetter Reader still has work to do")
				break

			case <-time.After(time.Minute * time.Duration(ServerConfig.Deadletter.Interval) / 2):
				fmt.Println("[INFO] DeadLetter Reader is Done")
				done <- true
				return
			}
		}
	}()

	return func(msg []byte) {
		stillWorking <- true
		fmt.Println("[INFO] Received Kafka message from topic 'product_created_dead_letter'")
		p := models.Product{}
		if err := json.Unmarshal(msg, &p); err != nil {
			fmt.Println("[ERROR] Unable to Unmarshal json from message 'product_created_dead_letter'", err.Error())
			return
		}

		if err := p.Save(db); err != nil {
			producer.Publish("product_created_dead_letter", msg)
		}
	}
}

func initDeadLettersReader() {
	cfg := common_io.Config{}

	err := config.Load("common_io_config.gcfg", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	flisten := func() {
		dlc := common_io.NewConsumer(cfg)

		done := make(chan bool)
		dlc.HandleTopic("product_created_dead_letter", handleProductCreatedDeadLetter(done))

		if err := dlc.StartListening(); err != nil {
			fmt.Println("[ERROR] ", err.Error())
		}
		<-done
		fmt.Println("[INFO] Will TearDown()")
		dlc.TearDown()
	}

	for {
		select {
		case <-time.After(time.Minute * time.Duration(ServerConfig.Deadletter.Interval)):
			fmt.Println("[INFO] Will execute fliste()")
			go flisten()
			break
		}
		fmt.Println("[INFO] Number of current active goroutines: ", runtime.NumGoroutine())
	}
}
