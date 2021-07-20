# Your redis connection deserves a name: PHP edition

An example on how to assign a name to a redis connection in PHP.

## Get it running

1. Start the redis docker container:
```sh
$ docker run --rm --publish 127.0.0.1:6379:6379/tcp --name ycdan-redis --detach redis:6.2.4
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
2021/07/20 06:05:32 Connecting to Redis on tcp://127.0.0.1:6379
2021/07/20 06:05:32 Connecting to Redis on tcp://127.0.0.1:6379  ... Successful
2021/07/20 06:05:32 Setting client name "your-connection-deserves-a-name-php"
2021/07/20 06:05:32 Setting client name "your-connection-deserves-a-name-php" ... Successful
2021/07/20 06:05:32
2021/07/20 06:05:32 Playing PING/PONG to keep connection open ...
2021/07/20 06:05:32 You can connect to redis and execute "CLIENT LIST" to see the connection
2021/07/20 06:05:32 Hit CTRL + C or cancel the process to stop.
2021/07/20 06:05:32
2021/07/20 06:05:32 PING ...
2021/07/20 06:05:32      ... PONG
2021/07/20 06:05:37 PING ...
2021/07/20 06:05:37      ... PONG
...
```

## Don't know what this is all about?

Checkout the [projects README](https://github.com/andygrunwald/your-connection-deserves-a-name#readme) to find out more.
