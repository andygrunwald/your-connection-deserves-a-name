![redis logo](../../images/redis-logo.png)

# your _redis_ connection deserves a name: Go edition

An example on how to assign a name to a [redis](https://redis.io/) connection in Go.

## Get it running

1. Start the redis docker container:
```sh
$ docker run --rm \
    --publish 127.0.0.1:6379:6379/tcp \
    --name ycdan-redis \
    --detach \
    redis:6.2.4
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
2021/07/30 19:23:59 Connecting to redis on localhost:6379
2021/07/30 19:23:59 Setting client name "currency-conversion-app"
2021/07/30 19:23:59 Setting client name "currency-conversion-app" ... Successful
2021/07/30 19:23:59 Connecting to redis on localhost:6379 ... Successful
2021/07/30 19:23:59
2021/07/30 19:23:59 Playing PING/PONG to keep connection open ...
2021/07/30 19:23:59 You can connect to redis and execute "CLIENT LIST" to see the connection
2021/07/30 19:23:59 Hit CTRL + C or cancel the process to stop.
2021/07/30 19:23:59
2021/07/30 19:23:59 PING ...
2021/07/30 19:23:59      ... PONG
2021/07/30 19:24:04 PING ...
2021/07/30 19:24:04      ... PONG
...
```

## Don't know what this is all about?

Read the original blog post [_your database connection deserves a name @ andygrunwald.com_](https://andygrunwald.com/blog/your-database-connection-deserves-a-name/ "Article your database connection deserves a name at Andy Grunwalds blog").

Additionally, you can check out the [projects README](https://github.com/andygrunwald/your-connection-deserves-a-name#readme).