package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("demo.nats.io", nats.Name("my-connection!"))
	if err != nil {
		log.Fatal(err)
	}

	nc.Subscribe("foo", func(msg *nats.Msg) {
		log.Println("[Received]", string(msg.Data))
	})

	log.Println("Connected to NATS!")

	select {}
}
