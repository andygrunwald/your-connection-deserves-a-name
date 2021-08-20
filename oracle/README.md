![Oracle logo](../images/oracle-logo.png)

# your _Oracle_ connection deserves a name

Examples on how to assign a particular name to a [Oracle Database](https://www.oracle.com/de/database/technologies/appdev/xe.html) connection.

Programmming languages:

- [Go](./go)

## How it works

While executing a query on the Oracle database, you can provide a client name (and other client information) as query parameters.
This is called [`DBMS_APPLICATION_INFO`](https://docs.oracle.com/en/database/oracle/oracle-database/18/arpls/DBMS_APPLICATION_INFO.html#GUID-14484F86-44F2-4B34-B34E-0C873D323EAD).

Here is an example how it works in Go (using [github.com/godror/godror](https://github.com/godror/godror)):

```go
// Creating a connection to the oracle database
[...]

// Adding DBMS_APPLICATION_INFO
ctx := godror.ContextWithTraceTag(context.Background(), godror.TraceTag{
    ClientIdentifier: ConnectionName,
    ClientInfo:       "Demo showcase",
    DbOp:             "ping",
    Module:           "oracle/go",
    Action:           "main",
})

// Sending DBMS_APPLICATION_INFO
rows, err := client.QueryContext(ctx, "SELECT sysdate FROM dual")
```

To see which clients are connected (incl. client information) and executed query statement, you can ask the `v$session` and `v$sqlarea` tables:

```sql
SQL> SELECT sess.username, sess.client_identifier, sess.module, sess.action, area.sql_text
FROM v$session sess, v$sqlarea area
WHERE
    sess.sql_address = area.address
    AND sess.username = 'DEMO';

USERNAME        CLIENT_IDENTIFIER         MODULE          ACTION          SQL_TEXT
--------------- ------------------------- --------------- --------------- ---------------------------
DEMO            currency-conversion-app   oracle/go       main            SELECT sysdate FROM dual
```

## Don't know what this is all about?

Read the original blog post [_your database connection deserves a name @ andygrunwald.com_](https://andygrunwald.com/blog/your-database-connection-deserves-a-name/ "Article your database connection deserves a name at Andy Grunwalds blog").

Additionally, you can check out the [projects README](https://github.com/andygrunwald/your-connection-deserves-a-name#readme).