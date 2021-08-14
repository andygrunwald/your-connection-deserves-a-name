![MSSQL / SQL-Server logo](../images/sql-server-logo.png)

# your _MSSQL / SQL-Server_ connection deserves a name

Examples on how to assign a particular name to a [MSSQL / SQL-Server](https://www.microsoft.com/en-us/sql-server/sql-server-2019) connection.

Programmming languages:

- [Go](./go)

## How it works (TODO)

While creating a [connection to SQL-Server, you can set an `Application Name`](https://docs.microsoft.com/en-us/dotnet/api/system.data.sqlclient.sqlconnection.connectionstring?view=dotnet-plat-ext-5.0).
This is (mostly) part part of the data source name (dsn)/connection string.

Here is an example in Go:

```go
query := url.Values{}
query.Add("app name", "currency-conversion-app")

u := &url.URL{
    Scheme:   "sqlserver",
    User:     url.UserPassword("sa", "yourStrong(!)Password"),
    Host:     fmt.Sprintf("%s:%d", "127.0.0.1", 1433),
    Path:     "/",
    RawQuery: query.Encode(),
}
dsn := u.String()

client, err := sql.Open("sqlserver", dsn)
```

To see which clients are connected (incl. their application name), you can query the `sys.sysprocesses` table:

```sql
SELECT
    hostname,
    program_name,
    loginame,
    cmd
FROM sys.sysprocesses
WHERE program_name != \"\";
```

The result should look similar to:

```
hostname       program_name                  loginame       cmd
-------------- ----------------------------- -------------- ----------------------
lap-dev        currency-conversion-app       sa             AWAITING COMMAND
```

## Don't know what this is all about?

Read the original blog post [_your database connection deserves a name @ andygrunwald.com_](https://andygrunwald.com/blog/your-database-connection-deserves-a-name/ "Article your database connection deserves a name at Andy Grunwalds blog").

Additionally, you can check out the [projects README](https://github.com/andygrunwald/your-connection-deserves-a-name#readme).