![PostgreSQL logo](../../images/postgresql-logo.png)

# your _PostgreSQL_ connection deserves a name: Go edition

An example on how to assign a name to a [PostgreSQL](https://www.postgresql.org/) connection in Go.

## Get it running

1. Start the PostgreSQL docker container:
```sh
$ docker run --rm \
    --publish 5432:5432 \
    --name ycdan-postgresql \
    --env POSTGRES_USER=postgres \
    --env POSTGRES_PASSWORD=postgrespassword \
    --detach \
    postgres:13.3
```

2. Compile the example program:
```sh
$ go build -o your-connection-deserves-a-name
```

3. Start the example program:
```sh
$ ./your-connection-deserves-a-name
```

You should see something like

```
2021/07/30 19:15:50 Connecting to PostgreSQL on postgres://postgres:postgrespassword@127.0.0.1/postgres
2021/07/30 19:15:50 Connecting to PostgreSQL on postgres://postgres:postgrespassword@127.0.0.1/postgres ... Successful
2021/07/30 19:15:50
2021/07/30 19:15:50 Keeping the connection open ...
2021/07/30 19:15:50 You can connect to the PostgreSQL database and execute the query:
2021/07/30 19:15:50 	SELECT pid, usename, application_name, client_addr, state, query, backend_type FROM pg_stat_activity;
2021/07/30 19:15:50 Hit CTRL + C or cancel the process to stop.
2021/07/30 19:15:50
```

4. Login into your database and execute the query:
```sql
postgres=# SELECT usename, application_name, client_addr, backend_type FROM pg_stat_activity;

 usename  |    application_name     | client_addr |  backend_type
----------+-------------------------+-------------+---------------
 postgres | currency-conversion-app | 172.17.0.1  | client backend
```

## Don't know what this is all about?

Read the original blog post [_your database connection deserves a name @ andygrunwald.com_](https://andygrunwald.com/blog/your-database-connection-deserves-a-name/ "Article your database connection deserves a name at Andy Grunwalds blog").

Additionally, you can check out the [projects README](https://github.com/andygrunwald/your-connection-deserves-a-name#readme).