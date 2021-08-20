from datetime import datetime
import time
import cx_Oracle


def output(message:str):
    now = datetime.now()
    time = now.strftime("%Y/%m/%d %H:%M:%S")
    print(f"{time} {message}")


oracle_username = "demo"
oracle_password = "showcase"
oracle_address  = "127.0.0.1:1521/connection_showcase"
connection_name = "unit-conversion-app"

# Additional DBMS_APPLICATION_INFO
client_info = "Demo showcase"
db_op       = "ping"
module     = "oracle/go"
action     = "main"

# Building the connection
dsn = oracle_address
output(f"Connecting to Oracle on {dsn}");

connection = cx_Oracle.connect(
                user=oracle_username,
                password=oracle_password,
                dsn=dsn,
                encoding="UTF-8")

with connection:
    connection.client_identifier = connection_name
    connection.clientinfo = client_info
    connection.dbop = db_op
    connection.module = module
    connection.action = action

    with connection.cursor() as cursor:
        cursor.execute("SELECT sysdate FROM dual")

    connection.commit()
    output(f"Connecting to Oracle on {dsn}  ... Successful");

    output("")
    output("Keeping the connection open ...")
    output("You can connect to the Oracle database and execute the query:")
    output("	SELECT username, client_identifier, module, action FROM v$session WHERE username='DEMO';")
    output("")
    output("or with current query:")
    output("	SELECT sess.username, sess.client_identifier, sess.module, sess.action, area.sql_text")
    output("	FROM v$session sess, v$sqlarea area")
    output("	WHERE")
    output("		sess.sql_address = area.address")
    output("		AND sess.username = 'DEMO';")
    output("Hit CTRL + C or cancel the process to stop.")
    output("")

    # Just looping to keep the process running.
    # In detail we do not keep the connection active, but rely on the standard timeout.
    # This way we provide time for the user to check the connection name behaviour.
    with connection.cursor() as cursor:
        while True:
            sql = "SELECT sysdate FROM dual"
            cursor.execute(sql)
            time.sleep(5)
