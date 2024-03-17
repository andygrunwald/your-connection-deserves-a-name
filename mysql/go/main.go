package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	MySQLUsername           = "root"
	MySQLPassword           = "secret"
	MySQLHostname           = "127.0.0.1"
	MySQLPort               = 3306
	MySQLDatabase           = "dummy"
	ConnectionAttributeName = "program_name"
	ConnectionName          = "currency-conversion-app"
)

func main() {
	// Doing all the magic: Setting a proper connection/application name.
	//
	// Connection attributes are key-value pairs that application programs can pass to the server at connect time.
	// This property can be added to the connection string.
	//
	// The application name will be maintained in the `session_connect_attrs` table of your MySQL-Server.
	//
	// MySQL docs:
	//	- Performance Schema Connection Attribute Tables: https://dev.mysql.com/doc/refman/8.0/en/performance-schema-connection-attribute-tables.html
	//	- Connection Attribute Limits: https://dev.mysql.com/doc/refman/8.0/en/performance-schema-connection-attribute-tables.html#performance-schema-connection-attribute-limits

	dsn := fmt.Sprintf("%s@tcp(%s:%d)/%s?connectionAttributes=%s:%s", url.UserPassword(MySQLUsername, MySQLPassword), MySQLHostname, MySQLPort, MySQLDatabase, ConnectionAttributeName, ConnectionName)

	// Building the connection
	log.Printf("Connecting to MySQL on %s\n", dsn)
	client, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Close(); err != nil {
			panic(err)
		}
	}()

	// Dummy values for demo purpose
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	// Establish the connection, because depending on the DB driver, we might only
	// validate the credentials, but don't built up a connection.
	err = client.Ping()
	if err != nil {
		panic(err)
	}
	log.Printf("Connecting to MySQL on %s ... Successful\n", dsn)

	log.Println("")
	log.Println("Keeping the connection open ...")
	log.Println("You can connect to the MySQL database and execute the query:")
	log.Println("	SELECT")
	log.Println("	    session_connect_attrs.ATTR_VALUE AS program_name,")
	log.Println("	    processlist.*")
	log.Println("	FROM information_schema.processlist")
	log.Println("	LEFT JOIN  performance_schema.session_connect_attrs ON (")
	log.Println("	    processlist.ID = session_connect_attrs.PROCESSLIST_ID")
	log.Println("	    AND session_connect_attrs.ATTR_NAME = \"program_name\"")
	log.Println("	)")
	log.Println("Hit CTRL + C or cancel the process to stop.")
	log.Println("")

	// Just looping to keep the process running.
	// In detail we do not keep the connection active, but rely on the standard timeout.
	// This way we provide time for the user to check the connection name behaviour.
	for {
		err = client.Ping()
		if err != nil {
			panic(err)
		}
		time.Sleep(5 * time.Second)
	}
}
