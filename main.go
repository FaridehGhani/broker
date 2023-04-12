package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	zmq "github.com/pebbe/zmq4"

	"github.com/FaridehGhani/broker/broker"
	"github.com/FaridehGhani/broker/destination"
	"github.com/FaridehGhani/broker/infra/repository"
	"github.com/FaridehGhani/broker/source"
)

func init() {
	// load project Env variables
	if err := godotenv.Load(); err != nil {
		log.Print("Error loading .env file")
	}

	// connect mongodb
	repository.NewMongoDBClient()
}

func main() {
	go source.Publish()
	go destination.Subscribe()
	go broker.Listen()

	time.Sleep(100 * time.Millisecond)

	subscriber, _ := zmq.NewSocket(zmq.XSUB)
	subscriber.Connect(os.Getenv("SUBSCRIBER_ENDPOINT"))
	publisher, _ := zmq.NewSocket(zmq.XPUB)
	publisher.Bind(os.Getenv("PUBLISHER_ENDPOINT"))
	listener, _ := zmq.NewSocket(zmq.PAIR)
	listener.Connect(os.Getenv("LISTENER_ENDPOINT"))
	zmq.Proxy(subscriber, publisher, listener)
}
