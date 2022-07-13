package source

import (
	"fmt"
	"log"
	"time"

	zmq "github.com/pebbe/zmq4"
	"github.com/thanhpk/randstr"
)

var randomNumber = 10

func Publish() {
	publisher, err := zmq.NewSocket(zmq.PUB)
	if err != nil {
		log.Fatalf("%s: %s", "publisher socket error", err)
	}
	err = publisher.Bind("tcp://*:6000")
	if err != nil {
		log.Fatalf("%s: %s", "publisher bind error", err)
	}

	// RandomNumber can change
	for i := 0; i < randomNumber; i++ {
		token := randstr.Hex(16)
		s := fmt.Sprintf("%s-%s", "random", token)
		_, err := publisher.SendMessage(s)
		if err != nil {
			log.Fatalf("%s: %s", "publisher send error", err)
		}
		time.Sleep(1 * time.Second)
	}
}
