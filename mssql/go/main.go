package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

const (
	SQLServerUsername = "sa"
	SQLServerPassword = "yourStrong(!)Password"
	SQLServerHostname = "127.0.0.1"
	SQLServerPort     = 1433
	ConnectionName    = "currency-conversion-app"
)

func main() {
	// Doing all the magic: Setting a proper connection/application name.
	//
	// In SQL-Server this is supported via the custom property `app name`.
	// This property can be added to the connection string.
	//
	// The application name will be maintained in the `sysprocesses` table of your SQL-Server.
	//
	// SQL-Server docs:
	//	- SqlConnection.ConnectionString Property: https://docs.microsoft.com/en-us/dotnet/api/system.data.sqlclient.sqlconnection.connectionstring
	//	- sys.dm_exec_sessions (Transact-SQL): https://docs.microsoft.com/en-us/sql/relational-databases/system-dynamic-management-views/sys-dm-exec-sessions-transact-sql?view=sql-server-ver15
	//
	query := url.Values{}
	query.Add("app name", ConnectionName)

	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(SQLServerUsername, SQLServerPassword),
		Host:     fmt.Sprintf("%s:%d", SQLServerHostname, SQLServerPort),
		Path:     "/",
		RawQuery: query.Encode(),
	}
	dsn := u.String()

	// Building the connection
	log.Printf("Connecting to SQL-Server on %s\n", dsn)
	client, err := sql.Open("sqlserver", dsn)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Close(); err != nil {
			panic(err)
		}
	}()

	// Establish the connection, because depending on the DB driver, we might only
	// validate the credentials, but don't built up a connection.
	err = client.Ping()
	if err != nil {
		panic(err)
	}
	log.Printf("Connecting to SQL-Server on %s ... Successful\n", dsn)

	log.Println("")
	log.Println("Keeping the connection open ...")
	log.Println("You can connect to the SQL-Server database and execute the query:")
	log.Println("	SELECT hostname, program_name, loginame, cmd FROM sys.sysprocesses WHERE program_name != \"\";")
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
