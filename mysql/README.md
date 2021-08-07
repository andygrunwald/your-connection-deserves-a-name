![MySQL logo](../images/mysql-logo.png)

# your _MySQL_ connection deserves a name

Examples on how to assign a particular name to a [MySQL](https://www.mysql.com/) connection.

Programmming languages:

- [Python (with PyMySQL)](./python-PyMySQL)
- [Python (with mysql-connector-python)](./python-mysql-connector-python)
- Go: Not supported yet, see [Support for sending connection attributes #737 @ go-sql-driver/mysql](https://github.com/go-sql-driver/mysql/pull/737)
- PHP: Not supported yet, see [Add possibility to add MySQL Connection Attributes (incl. custom ones) @ PHP Bugtracker](https://bugs.php.net/bug.php?id=81314)

## How it works

While creating a [connection to MySQL, you can set connection attributes](https://dev.mysql.com/doc/refman/8.0/en/performance-schema-connection-attribute-tables.html).
Depending on the client library this is either part of the data source name (dsn)/connection string or provided via a function call (mostly setting a kind of option).

For the application name, the connection attribute `program_name` is suggested.
Even if this is not a strict rule, it has been established as best practice.
Many applications follow this suggestion.

Here is an example in Python with the [PyMySQL library](https://pypi.org/project/PyMySQL/):

```python
connection = pymysql.connect(
                host='127.0.0.1',
                user='root',
                password='secret',
                database='dummy',
                charset='utf8mb4',
                cursorclass=pymysql.cursors.DictCursor,
                program_name='unit-conversion-app',
            )
```

To see which clients are connected (incl. their application name), you can query the `performance_schema` and `information_schema` schema:

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