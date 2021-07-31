![Logo](images/assign-a-name-to-your-connection.png)

# your connection deserves a name

👀 When your app interacts with an external system, **assign a name to the connection**.
An external system in this context can be things like a **database**, a **cache**, a **message queue** or an **HTTP endpoint**.

🎯 The goal should be: the **external system can identify their clients**.

🔥 During an incident, it will **reduce the time to debug by multiple hours** and often save other applications from failing.

➡️ Read more about this at [_your connection deserves a name @ andygrunwald.com_](https://andygrunwald.com/blog/your-connection-deserves-a-name/ "Article your connection deserves a name at Andy Grunwalds blog").

## How to do it (with examples)

This repository provides you example code for various systems in different programming languages:

- [MongoDB](./mongodb/)
- [MySQL](./mysql/)
- [PostgreSQL](./postgresql/)
- [RabbitMQ](./rabbitmq/)
- [redis](./redis/)

## Missing a system or a programming language?

If you

* know a system that supports connection naming, and it is not listed here
* miss a programing language example in your favorite language

let us know in either of two ways:

1. [Create an Issue](https://github.com/andygrunwald/your-connection-deserves-a-name/issues/new) with all the details you have in mind
2. or [create a Pull Request](https://docs.github.com/en/desktop/contributing-and-collaborating-using-github-desktop/working-with-your-remote-repository-on-github-or-github-enterprise/creating-an-issue-or-pull-request#creating-a-pull-request) with the implementation.

We are happy to extend this project.