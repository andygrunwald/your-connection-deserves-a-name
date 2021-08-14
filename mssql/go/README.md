![MSSQL / SQL-Server logo](../../images/sql-server-logo.png)

# your _MSSQL / SQL-Server_ connection deserves a name: Go edition

An example on how to assign a name to a [MSSQL / SQL-Server](https://www.microsoft.com/en-us/sql-server/sql-server-2019) connection in Go.

## Get it running

1. Start the SQL-Server docker container:
```sh
$ docker run --rm \
    --publish 1433:1433 \
    --name ycdan-sql-server \
    --env ACCEPT_EULA=Y \
    --env MSSQL_PID=Developer \
    --env 'SA_PASSWORD=yourStrong(!)Password' \
    --detach \
    mcr.microsoft.com/mssql/server:2019-CU12-ubuntu-20.04
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
2021/08/14 10:47:55 Connecting to SQL-Server on sqlserver://sa:yourStrong%28%21%29Password@127.0.0.1:1433/?app+name=currency-conversion-app
2021/08/14 10:47:55 Connecting to SQL-Server on sqlserver://sa:yourStrong%28%21%29Password@127.0.0.1:1433/?app+name=currency-conversion-app ... Successful
2021/08/14 10:47:55
2021/08/14 10:47:55 Keeping the connection open ...
2021/08/14 10:47:55 You can connect to the SQL-Server database and execute the query:
2021/08/14 10:47:55 	SELECT hostname, program_name, loginame, cmd FROM sys.sysprocesses WHERE program_name != "";
2021/08/14 10:47:55 Hit CTRL + C or cancel the process to stop.
2021/08/14 10:47:55
```

4. Login into your database and execute the query:
```sql
1> SELECT hostname, program_name, loginame, cmd FROM sys.sysprocesses WHERE program_name != "";
2> Go

hostname       program_name                  loginame       cmd
-------------- ----------------------------- -------------- ----------------------
lap-dev        currency-conversion-app       sa             AWAITING COMMAND
```

## Don't know what this is all about?

Read the original blog post [_your database connection deserves a name @ andygrunwald.com_](https://andygrunwald.com/blog/your-database-connection-deserves-a-name/ "Article your database connection deserves a name at Andy Grunwalds blog").

Additionally, you can check out the [projects README](https://github.com/andygrunwald/your-connection-deserves-a-name#readme).