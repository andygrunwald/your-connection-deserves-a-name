![MongoDB logo](../../images/mongodb-logo.png)

# your _MongoDB_ connection deserves a name: Go edition

An example on how to assign a name to a [MongoDB](https://www.mongodb.com/) connection in Go.

## Get it running

1. Start the MongoDB docker container:
```sh
$ docker run --rm \
    --publish 27017:27017 \
    --name ycdan-mongodb \
    --env MONGO_INITDB_ROOT_USERNAME=root \
    --env MONGO_INITDB_ROOT_PASSWORD=secret \
    --detach \
    mongo:5.0.1
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
2021/07/30 18:46:37 Connecting to MongoDB on mongodb://root:secret@127.0.0.1:27017/
2021/07/30 18:46:37 Connecting to MongoDB on mongodb://root:secret@127.0.0.1:27017/ ... Successful
2021/07/30 18:46:37
2021/07/30 18:46:37 Keeping the connection open ...
2021/07/30 18:46:37 You can connect to the MongoDB server and execute the query:
2021/07/30 18:46:37 	db.currentOp()
2021/07/30 18:46:37 Alternatively, you should see the `appName` in the logs of the MongoDB server.
2021/07/30 18:46:37
2021/07/30 18:46:37 Hit CTRL + C or cancel the process to stop.
2021/07/30 18:46:37
```

4. Login into your mongodb instance and execute the query:
```sql
db.currentOp()
```

In the key `appName` you can confirm your connection name:

![Screenshot of the db.currentOp() query.](../../images/mongodb-db.currentOp.png)

Alternatively, you can also check the server logs:

```json
{
    "t": {
        "$date": "2021-07-30T16:30:18.042+00:00"
    },
    "attr": {
            "application": {
                "name": "currency-conversion-app"
            }
        }
    }
}
```

## Don't know what this is all about?

Read the original blog post [_your database connection deserves a name @ andygrunwald.com_](https://andygrunwald.com/blog/your-database-connection-deserves-a-name/ "Article your database connection deserves a name at Andy Grunwalds blog").

Additionally, you can check out the [projects README](https://github.com/andygrunwald/your-connection-deserves-a-name#readme).