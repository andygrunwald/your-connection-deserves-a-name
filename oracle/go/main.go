package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/godror/godror"
)

const (
	OracleUsername = "demo"
	OraclePassword = "showcase"
	OracleAddress  = "127.0.0.1:1521/connection_showcase"
	ConnectionName = "currency-conversion-app"
)

func main() {
	// Required to get godror running.
	// See https://godror.github.io/godror/doc/installation.html
	libraryPath := os.Getenv("ORACLE_LIB_DIR")
	libraryPath, err := filepath.Abs(libraryPath)
	if err != nil {
		panic(err)
	}

	// Building the connection
	log.Printf("Connecting to Oracle on %s (with library path %s)\n", OracleAddress, libraryPath)

	dsn := fmt.Sprintf(`user="%s" password="%s" connectString="%s" libDir="%s"`, OracleUsername, OraclePassword, OracleAddress, libraryPath)
	client, err := sql.Open("godror", dsn)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Close(); err != nil {
			panic(err)
		}
	}()

	// Doing all the magic: Setting a proper connection name.
	//
	// In Oracle this is supported via the DBMS_APPLICATION_INFO feature set.
	// This structure supports information about the particular client.
	// Even things like `module` and `action` are supported.
	//
	// These settings will be maintained in the `v$session` and `v$sqlarea` table of your Oracle database.
	//
	// Oracle docs:
	//	- DBMS_APPLICATION_INFO: https://docs.oracle.com/en/database/oracle/oracle-database/18/arpls/DBMS_APPLICATION_INFO.html#GUID-14484F86-44F2-4B34-B34E-0C873D323EAD
	//
	ctx := godror.ContextWithTraceTag(context.Background(), godror.TraceTag{
		ClientIdentifier: ConnectionName,
		ClientInfo:       "Demo showcase",
		DbOp:             "ping",
		Module:           "oracle/go",
		Action:           "main",
	})

	// Establish the connection, because depending on the DB driver, we might only
	// validate the credentials, but don't built up a connection.
	err = client.PingContext(ctx)
	if err != nil {
		panic(err)
	}
	log.Printf("Connecting to Oracle on %s (with library path %s) ... Successful\n", OracleAddress, libraryPath)

	log.Println("")
	log.Println("Keeping the connection open ...")
	log.Println("You can connect to the Oracle database and execute the query:")
	log.Println("	SELECT username, client_identifier, module, action FROM v$session WHERE username='DEMO';")
	log.Println("")
	log.Println("or with current query:")
	log.Println("	SELECT sess.username, sess.client_identifier, sess.module, sess.action, area.sql_text")
	log.Println("	FROM v$session sess, v$sqlarea area")
	log.Println("	WHERE")
	log.Println("		sess.sql_address = area.address")
	log.Println("		AND sess.username = 'DEMO';")
	log.Println("Hit CTRL + C or cancel the process to stop.")
	log.Println("")

	// Just looping to keep the process running.
	// In detail we do not keep the connection active, but rely on the standard timeout.
	// This way we provide time for the user to check the connection name behaviour.
	for {
		rows, err := client.QueryContext(ctx, "SELECT sysdate FROM dual")
		if err != nil {
			panic(err)
		}
		rows.Close()

		time.Sleep(5 * time.Second)
	}
}
