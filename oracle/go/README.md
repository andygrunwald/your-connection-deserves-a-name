![Oracle logo](../../images/oracle-logo.png)

# your _Oracle_ connection deserves a name: Go edition

An example on how to assign a name to a [Oracle Database](https://www.oracle.com/de/database/technologies/appdev/xe.html) connection in Go.

## Get it running

0. Install the [`ODPI-C` driver](https://oracle.github.io/odpi/doc/installation.html).

1. Start the Oracle docker container:
```sh
$ docker run --rm \
    --publish 1521:1521 \
    --name ycdan-oracle \
    --env ORACLE_PASSWORD=secret \
    --env ORACLE_DATABASE=connection_showcase \
    --env APP_USER=demo \
    --env APP_USER_PASSWORD=showcase \
    --detach \
    gvenzl/oracle-xe:18.4.0-slim
```

2. Compile the example program:
```sh
$ go build -o your-connection-deserves-a-name
```

3. Start the example program:
```sh
$ ORACLE_LIB_DIR=/Users/agrunwald/Downloads/instantclient_19_8 ./your-connection-deserves-a-name
```

You should see something like
```
2021/08/20 16:16:49 Connecting to Oracle on 127.0.0.1:1521/connection_showcase (with library path /Users/agrunwald/Downloads/instantclient_19_8)
2021/08/20 16:16:49 Connecting to Oracle on 127.0.0.1:1521/connection_showcase (with library path /Users/agrunwald/Downloads/instantclient_19_8) ... Successful
2021/08/20 16:16:49
2021/08/20 16:16:49 Keeping the connection open ...
2021/08/20 16:16:49 You can connect to the Oracle database and execute the query:
2021/08/20 16:16:49 	SELECT username, client_identifier, module, action FROM v$session WHERE username='DEMO';
2021/08/20 16:16:49
2021/08/20 16:16:49 or with current query:
2021/08/20 16:16:49 	SELECT sess.username, sess.client_identifier, sess.module, sess.action, area.sql_text
2021/08/20 16:16:49 	FROM v$session sess, v$sqlarea area
2021/08/20 16:16:49 	WHERE
2021/08/20 16:16:49 		sess.sql_address = area.address
2021/08/20 16:16:49 		AND sess.username = 'DEMO';
2021/08/20 16:16:49 Hit CTRL + C or cancel the process to stop.
2021/08/20 16:16:49
```

4. Login into your database and execute the query (e.g., with _sqlplus_):
```sql
SQL> SELECT sess.username, sess.client_identifier, sess.module, sess.action, area.sql_text
FROM v$session sess, v$sqlarea area
WHERE
    sess.sql_address = area.address
    AND sess.username = 'DEMO';

USERNAME	    CLIENT_IDENTIFIER	      MODULE	      ACTION	      SQL_TEXT
--------------- ------------------------- --------------- --------------- ----------------------------------------
DEMO		    currency-conversion-app   oracle/go	      main		      SELECT sysdate FROM dual
```

## Don't know what this is all about?

Read the original blog post [_your database connection deserves a name @ andygrunwald.com_](https://andygrunwald.com/blog/your-database-connection-deserves-a-name/ "Article your database connection deserves a name at Andy Grunwalds blog").

Additionally, you can check out the [projects README](https://github.com/andygrunwald/your-connection-deserves-a-name#readme).