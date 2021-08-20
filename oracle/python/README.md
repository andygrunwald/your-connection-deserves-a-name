![Oracle logo](../../images/oracle-logo.png)

# your _Oracle_ connection deserves a name: Python edition

An example on how to assign a name to a [Oracle Database](https://www.oracle.com/de/database/technologies/appdev/xe.html) connection in Python.

## Get it running

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
2021/08/20 17:12:08 Connecting to Oracle on 127.0.0.1:1521/connection_showcase
2021/08/20 17:12:09 Connecting to Oracle on 127.0.0.1:1521/connection_showcase  ... Successful
2021/08/20 17:12:09
2021/08/20 17:12:09 Keeping the connection open ...
2021/08/20 17:12:09 You can connect to the Oracle database and execute the query:
2021/08/20 17:12:09 	SELECT username, client_identifier, module, action FROM v$session WHERE username='DEMO';
2021/08/20 17:12:09
2021/08/20 17:12:09 or with current query:
2021/08/20 17:12:09 	SELECT sess.username, sess.client_identifier, sess.module, sess.action, area.sql_text
2021/08/20 17:12:09 	FROM v$session sess, v$sqlarea area
2021/08/20 17:12:09 	WHERE
2021/08/20 17:12:09 		sess.sql_address = area.address
2021/08/20 17:12:09 		AND sess.username = 'DEMO';
2021/08/20 17:12:09 Hit CTRL + C or cancel the process to stop.
2021/08/20 17:12:09
```

4. Login into your database and execute the query (e.g., with _sqlplus_):
```sql
SQL> SELECT username, client_identifier, module, action FROM v$session WHERE username='DEMO';

USERNAME        CLIENT_IDENTIFIER         MODULE          ACTION
--------------- ------------------------- --------------- ---------------
DEMO            unit-conversion-app       oracle/go       main
```

## Don't know what this is all about?

Read the original blog post [_your database connection deserves a name @ andygrunwald.com_](https://andygrunwald.com/blog/your-database-connection-deserves-a-name/ "Article your database connection deserves a name at Andy Grunwalds blog").

Additionally, you can check out the [projects README](https://github.com/andygrunwald/your-connection-deserves-a-name#readme).