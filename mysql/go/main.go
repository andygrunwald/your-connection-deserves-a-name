package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	MySQLAddress   = "root:secret@/dummy"
	ConnectionName = "currency-conversion-app"
)

func main() {
	// Building the connection
	log.Printf("Connecting to MySQL on %s\n", MySQLAddress)

	// Doing all the magic: Setting a proper connection attribute.
	//
	// In MySQL this is supported via a custom connection attribute.
	// You can choose a attribute on your own.
	// However, the docs say `program_name` is the common standard.
	// This property can be added to the connection string.
	//
	// The application name will be maintained in the `performance_schema.session_connect_attrs` table of your MySQL database.
	//
	// MySQL docs:
	//	- Performance Schema Connection Attribute Tables: https://dev.mysql.com/doc/refman/8.0/en/performance-schema-connection-attribute-tables.html
	//
	dsn := fmt.Sprintf("%s?connectAttrs=program_name:%s", MySQLAddress, ConnectionName)
	conn, err := sql.Open("mysql", dsn)
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
	log.Println("	SELECT")
	log.Println("	    session_connect_attrs.ATTR_VALUE AS program_name,")
	log.Println("	    processlist.*")
	log.Println("	FROM information_schema.processlist")
	log.Println("	LEFT JOIN performance_schema.session_connect_attrs ON (")
	log.Println("	    processlist.ID = session_connect_attrs.PROCESSLIST_ID")
	log.Println("	    AND session_connect_attrs.ATTR_NAME = \"program_name\"")
	log.Println("	)")
	log.Println("Hit CTRL + C or cancel the process to stop.")
	log.Println("")

	// Just looping to keep the process running.
	// In detail we do not keep the connection active, but rely on the standard timeout.
	// This way we provide time for the user to check the connection name behaviour.
	for {
		time.Sleep(5 * time.Second)
	}
}
