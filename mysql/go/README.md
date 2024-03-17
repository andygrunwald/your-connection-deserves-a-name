![MySQL logo](../../images/mysql-logo.png)

# your _MySQL_ connection deserves a name: Go edition (with github.com/go-sql-driver/mysql)

An example on how to assign a name to a [MySQL](https://www.mysql.com/) connection in Go.

## Get it running

1. Start the MySQL docker container:
```sh
$ docker run --rm \
    --publish 3306:3306 \
    --env MYSQL_ROOT_PASSWORD=secret \
    --env MYSQL_DATABASE=dummy \
    --name ycdan-mysql \
    --detach \
    mysql:8.0.36 mysqld --default-authentication-plugin=mysql_native_password
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
2024/03/17 09:04:05 Connecting to MySQL on root:secret@tcp(127.0.0.1:3306)/dummy?connectionAttributes=program_name:currency-conversion-app
2024/03/17 09:04:05 Connecting to MySQL on root:secret@tcp(127.0.0.1:3306)/dummy?connectionAttributes=program_name:currency-conversion-app ... Successful
2024/03/17 09:04:05
2024/03/17 09:04:05 Keeping the connection open ...
2024/03/17 09:04:05 You can connect to the MySQL database and execute the query:
2024/03/17 09:04:05 	SELECT
2024/03/17 09:04:05 	    session_connect_attrs.ATTR_VALUE AS program_name,
2024/03/17 09:04:05 	    processlist.*
2024/03/17 09:04:05 	FROM information_schema.processlist
2024/03/17 09:04:05 	LEFT JOIN  performance_schema.session_connect_attrs ON (
2024/03/17 09:04:05 	    processlist.ID = session_connect_attrs.PROCESSLIST_ID
2024/03/17 09:04:05 	    AND session_connect_attrs.ATTR_NAME = "program_name"
2024/03/17 09:04:05 	)
2024/03/17 09:04:05 Hit CTRL + C or cancel the process to stop.
2024/03/17 09:04:05
```

4. Login into your database and execute the query:
```sql
SELECT
    session_connect_attrs.ATTR_VALUE AS program_name,
    processlist.*
FROM information_schema.processlist
LEFT JOIN  performance_schema.session_connect_attrs ON (
    processlist.ID = session_connect_attrs.PROCESSLIST_ID
    AND session_connect_attrs.ATTR_NAME = "program_name"
)
```

The result should look similar to:

```
program_name        | ID | USER | HOST             | DB    | [...]
--------------------+----+------+------------------+-------+------
unit-conversion-app | 11 | root | 172.17.0.1:56382 | dummy | [...]
```

## Don't know what this is all about?

Read the original blog post [_your database connection deserves a name @ andygrunwald.com_](https://andygrunwald.com/blog/your-database-connection-deserves-a-name/ "Article your database connection deserves a name at Andy Grunwalds blog").

Additionally, you can check out the [projects README](https://github.com/andygrunwald/your-connection-deserves-a-name#readme).