package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	// Doing all the magic: Using an own database user to apply a workaround for the connection name.
	//
	// In MySQL the connection name is not natively supported (e.g., in comparision to PostgreSQL).
	// However, there is a "kind of" workaround:
	//	Create an own database user peer application
	//
	// MySQL keeps track of all connections and queries.
	// Via `SHOW PROCESSLIST` you can query them.
	//
	// MySQL docs:
	//	- CREATE USER Statement: https://dev.mysql.com/doc/refman/8.0/en/create-user.html
	//	- SHOW PROCESSLIST Statement: https://dev.mysql.com/doc/refman/8.0/en/show-processlist.html
	MySQLAddress = "stock-exchange-rates-app:newyork@/connection_name"
)

func main() {
	// Building the connection
	log.Printf("Connecting to MySQL on %s\n", MySQLAddress)
	conn, err := sql.Open("mysql", MySQLAddress)
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
	log.Printf("Connecting to MysQL on %s ... Successful\n", MySQLAddress)

	log.Println("")
	log.Println("Keeping the connection open ...")
	log.Println("You can connect to the MySQL database and execute the query:")
	log.Println("	SHOW PROCESSLIST;")
	log.Println("Hit CTRL + C or cancel the process to stop.")
	log.Println("")

	// Just looping to keep the process running.
	// In detail we do not keep the connection active, but rely on the standard timeout.
	// This way we provide time for the user to check the connection name behaviour.
	for {
		time.Sleep(5 * time.Second)
	}
}
