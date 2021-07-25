# your _PostgreSQL_ connection deserves a name: Go edition

An example on how to assign a name to a [PostgreSQL](https://www.postgresql.org/) connection in Go.

## Get it running

1. Start the PostgreSQL docker container:
```sh
$ docker run --rm --publish 5432:5432 --name ycdan-postgresql --env POSTGRES_USER=postgres --env POSTGRES_PASSWORD=postgrespassword --detach postgres:13.3
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
2021/07/22 10:48:51 Connecting to RabbitMQ on amqp://guest:guest@127.0.0.1:5672/
2021/07/22 10:48:51 Connecting to RabbitMQ on amqp://guest:guest@127.0.0.1:5672/ ... Successful
2021/07/22 10:48:51
2021/07/22 10:48:51 Keeping the connection open ...
2021/07/22 10:48:51 You can connect to the RabbitMQ management UI: http://127.0.0.1:15672/#/connections (Username: guest, Password: guest)
2021/07/22 10:48:51 	-> Check the Connections Tab and see the Connection Name
2021/07/22 10:48:51 Hit CTRL + C or cancel the process to stop.
2021/07/22 10:48:51
```

4. Login into your database and execute the query:
```sql
SELECT pid, usename, application_name, client_addr, state, query, backend_type
FROM pg_stat_activity;
```

In the column `application_name` you can confirm your connection name.

## Don't know what this is all about?

Read the original blog post [_your connection deserves a name @ andygrunwald.com_](https://andygrunwald.com/blog/your-connection-deserves-a-name/ "Article your connection deserves a name at Andy Grunwalds blog").

Additionally, you can check out the [projects README](https://github.com/andygrunwald/your-connection-deserves-a-name#readme).