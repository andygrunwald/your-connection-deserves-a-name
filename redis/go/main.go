package main

import (
	"context"
	"log"
	"time"

	redis "github.com/go-redis/redis/v8"
)

const (
	RedisAddress        = "localhost:6379"
	RedisConnectionName = "your-connection-deserves-a-name-go"
)

func main() {
	// Building the connection
	log.Printf("Connecting to redis on %s\n", RedisAddress)
	r := redis.NewClient(&redis.Options{
		Addr: RedisAddress,
		OnConnect: func(ctx context.Context, conn *redis.Conn) error {

			// Doing all the magic: Setting a proper connection name.
			//
			// In redis this is supported via the command "CLIENT SETNAME <connection-name>".
			// The name of the _current_ connection can be retrieved via "CLIENT GETNAME".
			// To see all names connected to a redis instance, execute "CLIENT LIST".
			// The connection name value is visible in the `name` field.
			//
			// Redis docs:
			//	- CLIENT SETNAME: https://redis.io/commands/client-setname
			//	- CLIENT GETNAME: https://redis.io/commands/client-getname
			//	- CLIENT LIST: https://redis.io/commands/client-list
			log.Printf("Setting client name \"%s\"", RedisConnectionName)
			err := conn.ClientSetName(ctx, RedisConnectionName).Err()
			log.Printf("Setting client name \"%s\" ... Successful", RedisConnectionName)

			return err
		},
	})
	defer func() {
		r.Close() // Dropping err for simplicity
	}()

	// Sending a PING to ensure the connection is build up (lazy connection handling)
	ctx := context.Background()
	err := r.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}
	log.Printf("Connecting to redis on %s ... Successful\n", RedisAddress)

	log.Println("")
	log.Println("Playing PING/PONG to keep connection open ...")
	log.Println("You can connect to redis and execute \"CLIENT LIST\" to see the connection")
	log.Println("Hit CTRL + C or cancel the process to stop.")
	log.Println("")

	// Sending a ping to keep the connection open and provide time for
	// the user to check the CLIENT SETNAME behaviour.
	for {
		log.Println("PING ...")
		v, err := r.Ping(ctx).Result()
		if err != nil {
			panic(err)
		}
		log.Printf("     ... %s", v)

		time.Sleep(5 * time.Second)
	}
}
