![redis logo](../../images/redis-logo.png)

# your _redis_ connection deserves a name: PHP edition

An example on how to assign a name to a [redis](https://redis.io/) connection in PHP.

## Get it running

1. Start the redis docker container:
```sh
$ docker run --rm \
    --publish 127.0.0.1:6379:6379/tcp \
    --name ycdan-redis \
    --detach \
    redis:6.2.4
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
2021/07/30 17:25:19 Connecting to Redis on tcp://127.0.0.1:6379
2021/07/30 17:25:19 Connecting to Redis on tcp://127.0.0.1:6379  ... Successful
2021/07/30 17:25:19 Setting client name "stock-exchange-rates-app"
2021/07/30 17:25:19 Setting client name "stock-exchange-rates-app" ... Successful
2021/07/30 17:25:19
2021/07/30 17:25:19 Playing PING/PONG to keep connection open ...
2021/07/30 17:25:19 You can connect to redis and execute "CLIENT LIST" to see the connection
2021/07/30 17:25:19 Hit CTRL + C or cancel the process to stop.
2021/07/30 17:25:19
2021/07/30 17:25:19 PING ...
2021/07/30 17:25:19      ... PONG
2021/07/30 17:25:24 PING ...
2021/07/30 17:25:24      ... PONG
...
```

## Don't know what this is all about?

Read the original blog post [_your database connection deserves a name @ andygrunwald.com_](https://andygrunwald.com/blog/your-database-connection-deserves-a-name/ "Article your database connection deserves a name at Andy Grunwalds blog").

Additionally, you can check out the [projects README](https://github.com/andygrunwald/your-connection-deserves-a-name#readme).