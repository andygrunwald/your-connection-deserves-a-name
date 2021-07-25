# your _RabbitMQ_ connection deserves a name: PHP edition

An example on how to assign a name to a [RabbitMQ](https://www.rabbitmq.com/) connection in PHP.

## Get it running

1. Start the RabbitMQ docker container:
```sh
$ docker run --rm --publish 15672:15672 --publish 5672:5672 --name ycdan-rabbitmq --detach rabbitmq:3-management
```

2. Install dependencies:
```sh
$ composer install
```

3. Start the example program:
```sh
$ php connection-name.php
```

You should see something like

```
2021/07/22 09:17:37 Connecting to RabbitMQ on amqp://guest:guest@127.0.0.1:5672/
2021/07/22 09:17:37 Connecting to RabbitMQ on amqp://guest:guest@127.0.0.1:5672/  ... Successful
2021/07/22 09:17:37
2021/07/22 09:17:37 Keeping the connection open ...
2021/07/22 09:17:37 You can connect to the RabbitMQ management UI: http://127.0.0.1:15672/#/connections (Username: guest, Password: guest)
2021/07/22 09:17:37 	-> Check the Connections Tab and see the Connection Name
2021/07/22 09:17:37 Hit CTRL + C or cancel the process to stop.
2021/07/22 09:17:37
```

4. Visit the RabbitMQ Management UI at http://127.0.0.1:15672/#/connections (Username: guest, Password: guest) and confirm the connection name.

## Don't know what this is all about?

Read the original blog post [_your connection deserves a name @ andygrunwald.com_](https://andygrunwald.com/blog/your-connection-deserves-a-name/ "Article your connection deserves a name at Andy Grunwalds blog").

Additionally, you can check out the [projects README](https://github.com/andygrunwald/your-connection-deserves-a-name#readme).