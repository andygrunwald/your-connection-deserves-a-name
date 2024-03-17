![MySQL logo](../../images/mysql-logo.png)

# your _MySQL_ connection deserves a name: Python edition (with PyMySQL)

An example on how to assign a name to a [MySQL](https://www.mysql.com/) connection in Python.

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

2. Install dependencies:
```sh
$ pip install -r requirements.txt
```

3. Start the example program:
```sh
$ python main.py
```

You should see something like

```
2021/07/31 08:39:14 Connecting to MySQL on mysql://root:secret@127.0.0.1/dummy
2021/07/31 08:39:14 Connecting to MySQL on mysql://root:secret@127.0.0.1/dummy  ... Successful
2021/07/31 08:39:14
2021/07/31 08:39:14 Keeping the connection open ...
2021/07/31 08:39:14 You can connect to the MySQL database and execute the query:
2021/07/31 08:39:14 	SELECT
2021/07/31 08:39:14 	    session_connect_attrs.ATTR_VALUE AS program_name,
2021/07/31 08:39:14 	    processlist.*
2021/07/31 08:39:14 	FROM information_schema.processlist
2021/07/31 08:39:14 	LEFT JOIN  performance_schema.session_connect_attrs ON (
2021/07/31 08:39:14 	    processlist.ID = session_connect_attrs.PROCESSLIST_ID
2021/07/31 08:39:14 	    AND session_connect_attrs.ATTR_NAME = "program_name"
2021/07/31 08:39:14 	)
2021/07/31 08:39:14 Hit CTRL + C or cancel the process to stop.
2021/07/31 08:39:14
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