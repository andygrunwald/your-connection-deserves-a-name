package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	MongoDBAddress = "mongodb://root:secret@127.0.0.1:27017/"
	ConnectionName = "currency-conversion-app"
)

func main() {
	// Building the connection
	log.Printf("Connecting to MongoDB on %s\n", MongoDBAddress)

	// Doing all the magic: Setting a proper application name.
	//
	// In MongoDB this is supported via the custom property `appName`.
	// This property can be added to the connection string / data source name (dsn).
	//
	// As a result, the appName can be seen in multiple places:
	//	- seerver logs
	//	- db.currentOp() query
	//	- system.profile.appName field in the database profiler output
	//
	// For more details please check the docs.
	//
	// MongoDB docs:
	//	- Miscellaneous Configuration - appName: https://docs.mongodb.com/manual/reference/connection-string/#mongodb-urioption-urioption.appName
	//
	dsn := fmt.Sprintf("%s?appName=%s", MongoDBAddress, ConnectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Establish the connection, because depending on the DB driver, we might only
	// validate the credentials, but don't built up a connection.
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	log.Printf("Connecting to MongoDB on %s ... Successful\n", MongoDBAddress)

	log.Println("")
	log.Println("Keeping the connection open ...")
	log.Println("You can connect to the MongoDB server and execute the query:")
	log.Println("	db.currentOp()")
	log.Println("Alternatively, you should see the `appName` in the logs of the MongoDB server.")
	log.Println("")
	log.Println("Hit CTRL + C or cancel the process to stop.")
	log.Println("")

	// Just looping to keep the process running.
	// In detail we do not keep the connection active, but rely on the standard timeout.
	// This way we provide time for the user to check the connection name behaviour.
	for {
		err = client.Ping(context.Background(), readpref.Primary())
		if err != nil {
			panic(err)
		}
		time.Sleep(5 * time.Second)
	}
}
