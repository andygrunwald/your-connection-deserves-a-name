![RabbitMQ logo](../images/rabbitmq-logo.png)

# your _RabbitMQ_ connection deserves a name

Examples on how to assign a particular name to a [RabbitMQ](https://www.rabbitmq.com/) connection.

Programmming languages:

- [Go](./go)
- [PHP](./php)

## How it works

While creating a connection to RabbitMQ, you can provide AMQP options.
One of the options is [`connection_name`](https://www.rabbitmq.com/connections.html#client-provided-names "AMQP Client-Provided Connection Name @ RabbitMQ docs").

This is how it looks like in Go:

```go
config := amqp.Config{
    Properties: amqp.Table{
        "connection_name": "currency-conversion-app",
    },
}
conn, err := amqp.DialConfig("amqp://guest:guest@127.0.0.1:5672/", config)
```

In the UI of Rabbit under the Connection tab, you can see all connected clients, including their names.

## How it looks like

### Before

![RabbitMQ: Client connections without name](../images/rabbitmq-connections-without-name.png)

## After

![RabbitMQ: Client connections with name](../images/rabbitmq-connections-with-name.png)

## Don't know what this is all about?

Read the original blog post [_your connection deserves a name @ andygrunwald.com_](https://andygrunwald.com/blog/your-connection-deserves-a-name/ "Article your connection deserves a name at Andy Grunwalds blog").

Additionally, you can check out the [projects README](https://github.com/andygrunwald/your-connection-deserves-a-name#readme).