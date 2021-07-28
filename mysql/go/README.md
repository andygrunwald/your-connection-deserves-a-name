# your _MySQL_ connection deserves a name: Go edition

An example on how to assign a name to a [MySQL](https://www.mysql.com/) connection in Go.

## Get it running

1. Start the MySQL docker container:
```sh
$ docker run --rm \
    --publish 3306:3306 \
    --env MYSQL_ROOT_PASSWORD=secret \
    --env MYSQL_DATABASE=connection_name \
    --env MYSQL_USER=stock-exchange-rates-app \
    --env MYSQL_PASSWORD=newyork \
    --name ycdan-mysql \
    --detach \
    mysql:8.0.26 mysqld --default-authentication-plugin=mysql_native_password
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
2021/07/28 17:33:43 Connecting to MySQL on stock-exchange-rates-app:newyork@/connection_name
2021/07/28 17:33:43 Connecting to MysQL on stock-exchange-rates-app:newyork@/connection_name ... Successful
2021/07/28 17:33:43
2021/07/28 17:33:43 Keeping the connection open ...
2021/07/28 17:33:43 You can connect to the MySQL database and execute the query:
2021/07/28 17:33:43 	SHOW PROCESSLIST;
2021/07/28 17:33:43 Hit CTRL + C or cancel the process to stop.
2021/07/28 17:33:43
```

4. Login into your database with username `root` and password `secret` and execute the query:
```sql
SHOW PROCESSLIST;
```

In the column `User` you can confirm your connection (aka username) name.

## Don't know what this is all about?

Read the original blog post [_your connection deserves a name @ andygrunwald.com_](https://andygrunwald.com/blog/your-connection-deserves-a-name/ "Article your connection deserves a name at Andy Grunwalds blog").

Additionally, you can check out the [projects README](https://github.com/andygrunwald/your-connection-deserves-a-name#readme).