# your _PostgreSQL_ connection deserves a name

Examples on how to assign a particular name to a [PostgreSQL](https://www.postgresql.org/) connection.

Programmming languages:

- [Go](./go)

## How it works

While creating a connection to PostgreSQL, you can provide a client name in the connection string.
The property is called [`application_name`](https://www.postgresql.org/docs/9.0/runtime-config-logging.html#GUC-APPLICATION-NAME) and is part of [libpq](https://www.postgresql.org/docs/9.0/libpq-connect.html).

```go
dsn := "postgres://user:pass@127.0.0.1/database?application_name=your-connection-deserves-a-name-go"
conn, err := sql.Open("postgres", dsn)
```

To see which clients are connected (incl. their application name), you can query the `pg_stat_activity` table:

```sql
postgres=# SELECT usename, application_name, client_addr, backend_type FROM pg_stat_activity;

 usename  |          application_name          | client_addr |         backend_type
----------+------------------------------------+-------------+------------------------------
 postgres | your-connection-deserves-a-name-go | 172.17.0.1  | client backend
```

## How it looks like

### Before

TODO

## After

TODO

## Don't know what this is all about?

Read the original blog post [_your connection deserves a name @ andygrunwald.com_](https://andygrunwald.com/blog/your-connection-deserves-a-name/ "Article your connection deserves a name at Andy Grunwalds blog").

Additionally, you can check out the [projects README](https://github.com/andygrunwald/your-connection-deserves-a-name#readme).