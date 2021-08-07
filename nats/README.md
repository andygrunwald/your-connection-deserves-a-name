![NATS.io logo](../images/nats-logo.png)

# your _NATS_ connection deserves a name

Examples on how to assign a particular name to a [NATS.io](https://nats.io/) connection.

Programmming languages:

- [Go](./go)

## How it works

While creating a connection to NATS, you can provide [client connection name](https://docs.nats.io/developing-with-nats/connecting/name "Connection Name @ NATS docs").

This is how it looks like in Go:

```go
nc, err := nats.Connect("demo.nats.io", nats.Name("currency-conversion-app"))
```

Via the [NATS monitoring endpoint](https://docs.nats.io/nats-server/configuration/monitoring#connection-information "Monitoring endpoint @ NATS docs"), you can see all connected clients, including their names.

PS: The NATS team offers a [demo server at `demo.nats.io`](https://docs.nats.io/nats-protocol/nats-protocol-demo "NATS demo servier"). See [http://demo.nats.io:8222/connz](http://demo.nats.io:8222/connz).

## How it looks like

### Before

```sh
$ curl http://127.0.0.1:8222/connz
{
  [...]
  "connections": [
    {
      "ip": "172.17.0.1",
      "port": 57046,
      [...]
      "lang": "go",
      "version": "1.11.0"
    }
  ]
}
```

## After

```sh
$ curl http://127.0.0.1:8222/connz
{
  [...]
  "connections": [
    {
      "ip": "172.17.0.1",
      "port": 57054,
      [...]
      "name": "currency-conversion-app",
      "lang": "go",
      "version": "1.11.0"
    }
  ]
}
```

## Don't know what this is all about?

Read the original blog post [_your database connection deserves a name @ andygrunwald.com_](https://andygrunwald.com/blog/your-database-connection-deserves-a-name/ "Article your database connection deserves a name at Andy Grunwalds blog").

Additionally, you can check out the [projects README](https://github.com/andygrunwald/your-connection-deserves-a-name#readme).