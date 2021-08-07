package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

const (
	NATSAddress            = "nats://127.0.0.1:4222"
	NATSMonitoringEndpoint = "http://127.0.0.1:8222/connz"
	ConnectionName         = "currency-conversion-app"
)

func main() {
	// Building the connection
	log.Printf("Connecting to NATS on %s\n", NATSAddress)

	// Doing all the magic: Setting a proper connection name.
	//
	// In NATS this is supported via the connection name you can set via during the connection creation process.
	// This is part of the NATS protocol.
	//
	// The NATS monitoring endpointdisplays the connection name for all connected clients.
	//
	// NATS docs:
	//	- Connection Name: https://docs.nats.io/developing-with-nats/connecting/name
	client, err := nats.Connect(NATSAddress, nats.Name(ConnectionName))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	log.Printf("Connecting to NATS on %s ... Successful\n", NATSAddress)

	log.Println("")
	log.Println("Keeping the connection open ...")
	log.Printf("You can connect to the NATS monitoring endpoint: %s\n", NATSMonitoringEndpoint)
	log.Println("Hit CTRL + C or cancel the process to stop.")
	log.Println("")

	// Just looping to keep the process running.
	// In detail we do not keep the connection active, but rely on the standard timeout.
	// This way we provide time for the user to check the connection name behaviour.
	for {
		time.Sleep(5 * time.Second)
	}
}
