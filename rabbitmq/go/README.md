# Your RabbitMQ connection deserves a name: Go edition

An example on how to assign a name to a RabbitMQ connection in Go(lang).

## Get it running

1. Start the RabbitMQ docker container:
```sh
$ docker run --rm --publish 15672:15672 --publish 5672:5672 --name ycdan-rabbitmq --detach rabbitmq:3-management
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

4. Visit the RabbitMQ Management UI at http://127.0.0.1:15672/#/connections (Username: guest, Password: guest) and confirm the connection name.

## Don't know what this is all about?

Checkout the [projects README](https://github.com/andygrunwald/your-connection-deserves-a-name#readme) to find out more.
