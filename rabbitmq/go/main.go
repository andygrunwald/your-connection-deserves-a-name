package main

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

const (
	RabbitMQAddress        = "amqp://guest:guest@127.0.0.1:5672/"
	RabbitMQManagementUI   = "http://127.0.0.1:15672/#/connections"
	RabbitMQConnectionName = "your-connection-deserves-a-name-go"
)

func main() {
	// Building the connection
	log.Printf("Connecting to RabbitMQ on %s\n", RabbitMQAddress)

	// Doing all the magic: Setting a proper connection name.
	//
	// In RabbitMQ this is supported via the custom properties you can set via the AMQP config.
	// This is part of the AMQP specification.
	//
	// The RabbitMQ Management UI displays the connection name in the Connection listing.
	//
	// RabbitMQ docs:
	//	- Connections: Client-Provided Connection Name: https://www.rabbitmq.com/connections.html#client-provided-names
	config := amqp.Config{
		Properties: amqp.Table{
			"connection_name": RabbitMQConnectionName,
		},
	}

	conn, err := amqp.DialConfig(RabbitMQAddress, config)
	if err != nil {
		panic(err)
	}
	defer func() {
		conn.Close() // Dropping err for simplicity
	}()
	log.Printf("Connecting to RabbitMQ on %s ... Successful\n", RabbitMQAddress)

	log.Println("")
	log.Println("Keeping the connection open ...")
	log.Printf("You can connect to the RabbitMQ management UI: %s (Username: guest, Password: guest)\n", RabbitMQManagementUI)
	log.Println("	-> Check the Connections Tab and see the Connection Name")
	log.Println("Hit CTRL + C or cancel the process to stop.")
	log.Println("")

	// Just looping to keep the process running.
	// In detail we do not keep the connection active, but rely on the standard timeout.
	// This way we provide time for the user to check the connection name behaviour.
	for {
		time.Sleep(5 * time.Second)
	}
}
