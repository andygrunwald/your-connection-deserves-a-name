![MongoDB logo](../images/mongodb-logo.png)

# your _MongoDB_ connection deserves a name

Examples on how to assign a particular name to a [MongoDB](https://www.mongodb.com/) connection.

Programmming languages:

- [Go](./go)


## How it works

While creating a connection to MongoDB, you can provide an application name in the connection string.
The property is called [`appName`](https://docs.mongodb.com/manual/reference/connection-string/#mongodb-urioption-urioption.appName).

```go
dsn := "mongodb://root:secret@127.0.0.1:27017/?appName=currency-conversion-app"
client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
```

To see which clients are connected (incl. their application name), you can query `db.currentOp()`:

![Screenshot of the db.currentOp() query.](../images/mongodb-db.currentOp.png)

or check the server logs:

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

Read the original blog post [_your connection deserves a name @ andygrunwald.com_](https://andygrunwald.com/blog/your-connection-deserves-a-name/ "Article your connection deserves a name at Andy Grunwalds blog").

Additionally, you can check out the [projects README](https://github.com/andygrunwald/your-connection-deserves-a-name#readme).