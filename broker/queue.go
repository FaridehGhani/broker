package broker

import (
	"github.com/FaridehGhani/broker/message"
	"log"

	zmq "github.com/pebbe/zmq4"

	"github.com/FaridehGhani/broker/infra/repository"
)

func Listen() {
	pipe, err := zmq.NewSocket(zmq.PAIR)
	if err != nil {
		log.Fatalf("%s: %s", "listener socket error", err)
	}
	err = pipe.Bind("inproc://pipe")
	if err != nil {
		log.Fatalf("%s: %s", "listener bind error", err)
	}

	db := repository.NewDB()

	for {
		msgs, err := pipe.RecvMessage(0)
		if err != nil {
			log.Fatalf("%s: %s", "listener recieve message error", err)
		}

		err = db.InsertManyMessages(message.ListMessages(msgs))
		if err != nil {
			log.Println("insert message error: ", err)
		}
	}
}
