package destination

import (
	"log"

	zmq "github.com/pebbe/zmq4"
)

func Subscribe() {
	subscriber, err := zmq.NewSocket(zmq.SUB)
	if err != nil {
		log.Fatalf("%s: %s", "subscriber socket error", err)
	}

	err = subscriber.Connect("tcp://localhost:6001")
	if err != nil {
		log.Fatalf("%s: %s", "subscriber connect error", err)
	}

	err = subscriber.SetSubscribe("random")
	if err != nil {
		log.Fatalf("%s: %s", "subscriber connect error", err)
	}
	defer subscriber.Close()

	var receivedMessages int
	for {
		_, err := subscriber.RecvMessage(0)

		if err != nil {
			log.Fatalf("%s: %s", "subscriber recieve message error", err)
		}
		receivedMessages++
		log.Println("received messages :: ", receivedMessages)
	}
}
