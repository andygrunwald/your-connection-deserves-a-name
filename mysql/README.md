# your _MySQL_ connection deserves a name

Examples on how to assign a particular name to a [MySQL](https://www.mysql.com/) connection.

Programmming languages:

- [Go](./go)

## How it works

Sadly, MySQL **does not provide a dedicated feature to assign a particular name to a connection**.

However, there is a workaround:
Ensuring that each application operates on its own username.

For this, a bit of preparation is needed:
* Get to know which database operations the application is executing
* Create a new database user with particular permissions

The [`CREATE USER`](https://dev.mysql.com/doc/refman/8.0/en/create-user.html) and [GRANT](https://dev.mysql.com/doc/refman/8.0/en/grant.html) Statements provide instructions on how to create a user and assign particulr permissions.

Use this new user to build up the connection:

```go
// <user>:<password>@<ip>:<port>/<database-namee>
dsn := "stock-exchange-rates-app:newyork@/connection_name"
conn, err := sql.Open("mysql", dsn)
```

To see which clients are connected (incl. their username), you can query the [processlist](https://dev.mysql.com/doc/refman/8.0/en/show-processlist.html):

```sql
SHOW PROCESSLIST;

Id	User	                    Host	            db	            Command	Time	State	                Info
8	root	                    172.17.0.1:59426	connection_name	Query	0	    init	                SHOW PROCESSLIST
10	stock-exchange-rates-app	172.17.0.1:59434	connection_name	Sleep	5		                        NULL
```

## Don't know what this is all about?

Read the original blog post [_your connection deserves a name @ andygrunwald.com_](https://andygrunwald.com/blog/your-connection-deserves-a-name/ "Article your connection deserves a name at Andy Grunwalds blog").

Additionally, you can check out the [projects README](https://github.com/andygrunwald/your-connection-deserves-a-name#readme).