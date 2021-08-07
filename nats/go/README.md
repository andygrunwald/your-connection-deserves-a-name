![NATS.io logo](../../images/nats-logo.png)

# your _NATS_ connection deserves a name: Go edition

An example on how to assign a name to a [NATS.io](https://nats.io/) connection in Go.

## Get it running

1. Start the NATS docker container:
```sh
$ docker run --rm \
    --publish 4222:4222 \
    --publish 8222:8222 \
    --name ycdan-nats \
    --detach \
    nats:2.3.4
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
2021/08/07 14:17:17 Connecting to NATS on nats://127.0.0.1:4222
2021/08/07 14:17:17 Connecting to NATS on nats://127.0.0.1:4222 ... Successful
2021/08/07 14:17:17
2021/08/07 14:17:17 Keeping the connection open ...
2021/08/07 14:17:17 You can connect to the NATS monitoring endpoint: http://127.0.0.1:8222/connz
2021/08/07 14:17:17 Hit CTRL + C or cancel the process to stop.
2021/08/07 14:17:17
```

4. Visit the NATS monitoring endpoint at http://127.0.0.1:8222/connz and confirm the connection name.

## Don't know what this is all about?

Read the original blog post [_your database connection deserves a name @ andygrunwald.com_](https://andygrunwald.com/blog/your-database-connection-deserves-a-name/ "Article your database connection deserves a name at Andy Grunwalds blog").

Additionally, you can check out the [projects README](https://github.com/andygrunwald/your-connection-deserves-a-name#readme).