package main

import (
	"encoding/json"
	"fmt"
	"log"

	"usertest.com/broker"
	"usertest.com/config"
	"usertest.com/event"
)

func main() {
	config := config.GetConfig()
	log.Printf("Successful %v", config)

	bs, err := broker.NewRabbitConnectionForDomain(config.MessageBroker)
	if err != nil {
		log.Fatalf("ERROR: can't initialize the broker: %s\n", err)
	}
	defer bs.Close()

	consumer, err := bs.Consumer(broker.DomainQueue)
	if err != nil {
		log.Fatal("ERROR: Can't connect to broker")
	}

	forever := make(chan bool)
	go func() {
		for msg := range consumer {
			e := event.DomainEvent{}
			err := json.Unmarshal(msg.Body, &e)
			if err != nil {
				log.Printf("ERROR: unmarshalling event %s\n", msg.Body)
				continue
			}
			
			fmt.Printf("Received Message: %#v\n", e)
		}
	}()

	fmt.Println("Waiting for messages...")
	<-forever
}
