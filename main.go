package main

import (
	"time"

	zmq "github.com/pebbe/zmq4"

	"github.com/FaridehGhani/broker/broker"
	"github.com/FaridehGhani/broker/destination"
	"github.com/FaridehGhani/broker/infra/repository"
	"github.com/FaridehGhani/broker/source"
)

func init() {
	// connect mongodb
	repository.NewMongoDBClient()
}

func main() {
	go source.Publish()
	go destination.Subscribe()
	go broker.Listen()

	time.Sleep(100 * time.Millisecond)

	subscriber, _ := zmq.NewSocket(zmq.XSUB)
	subscriber.Connect("tcp://localhost:6000")
	publisher, _ := zmq.NewSocket(zmq.XPUB)
	publisher.Bind("tcp://*:6001")
	listener, _ := zmq.NewSocket(zmq.PAIR)
	listener.Connect("inproc://pipe")
	zmq.Proxy(subscriber, publisher, listener)
}
