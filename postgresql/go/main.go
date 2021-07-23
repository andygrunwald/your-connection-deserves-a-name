package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	PostgreSQLAddress        = "postgres://postgres:postgrespassword@127.0.0.1/postgres"
	PostgreSQLConnectionName = "your-connection-deserves-a-name-go"
)

func main() {
	// Building the connection
	log.Printf("Connecting to PostgreSQL on %s\n", PostgreSQLAddress)

	// Doing all the magic: Setting a proper connection name.
	//
	// In PostgreSQL this is supported via the custom property `application_name`.
	// This property can be added to the connection string.
	//
	// The application name will be maintained in the `pg_stat_activity` table of your PostgreSQL database.
	//
	// Side note:
	//	sslmode is not part of the connection name.
	//	This is only needed due to the showcase docker setup.
	//
	// PostgreSQL docs:
	//	- Runtime config: application_name: https://www.postgresql.org/docs/9.0/runtime-config-logging.html#GUC-APPLICATION-NAME
	//	- libpq connect: https://www.postgresql.org/docs/9.0/libpq-connect.html
	connStr := fmt.Sprintf("%s?sslmode=disable&application_name=%s", PostgreSQLAddress, PostgreSQLConnectionName)
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer func() {
		conn.Close() // Dropping err for simplicity
	}()

	// Establish the connection, because depending on the DB driver, we might only
	// validate the credentials, but don't built up a connection.
	err = conn.Ping()
	if err != nil {
		panic(err)
	}
	log.Printf("Connecting to PostgreSQL on %s ... Successful\n", PostgreSQLAddress)

	log.Println("")
	log.Println("Keeping the connection open ...")
	log.Println("You can connect to the PostgreSQL database and execute the query:")
	log.Println("	SELECT pid, usename, application_name, client_addr, state, query, backend_type FROM pg_stat_activity;")
	log.Println("Hit CTRL + C or cancel the process to stop.")
	log.Println("")

	// Just looping to keep the process running.
	// In detail we do not keep the connection active, but rely on the standard timeout.
	// This way we provide time for the user to check the connection name behaviour.
	for {
		time.Sleep(5 * time.Second)
	}
}
