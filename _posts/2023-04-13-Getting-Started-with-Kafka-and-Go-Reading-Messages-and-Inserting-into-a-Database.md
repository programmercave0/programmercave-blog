---
layout: post
title: "Getting Started with Kafka and Go: Reading Messages and Inserting into a Database"
description: "In the below code, messages are read from a Kafka topic, parsed as JSON, and then inserted into a PostgreSQL database. This is a common use case for Kafka, where it is used as a data source for other applications or databases."
author: "Programmercave"
header-img: "/assets/Kafka/kgp.png"
tags:  [Go, Kafka, SQL, PostgreSQL, DataBase]
date: 2023-04-13
---
* toc
{:toc}

# Introduction:

Kafka is a popular distributed streaming platform that is used for building real-time data pipelines and streaming applications. It is designed to handle high-volume data streams in real-time and provide reliable and scalable data streaming. Kafka is used for a wide range of applications, including real-time analytics, monitoring, log aggregation, and messaging systems.

In this blog post, we will demonstrate how to read Kafka messages in Go and store them in a PostgreSQL database. This is a common use case for Kafka, where it is used as a data source for other applications or databases. We will use the kafka-go package to read messages from a Kafka topic, parse them as JSON, and then use the database/sql package to insert them into a PostgreSQL database.

By the end of this blog post, you will have a good understanding of how to set up a Kafka reader in Go, how to process Kafka messages in Go, and how to store them in a PostgreSQL database. This can be useful for building real-time data pipelines or for monitoring and analyzing high-volume data streams.

# Installing Kafka on a Machine:

To install Kafka on a machine, follow the below steps:

### Step 1: Download Kafka

Go to the official Kafka website [https://kafka.apache.org/downloads](https://kafka.apache.org/downloads) and download the latest stable release of Kafka. Extract the downloaded tarball using the following command:

```shell
tar -xzf kafka_2.13-2.8.0.tgz
```

Note: The version number may vary depending on the latest release.

### Step 2: Start ZooKeeper

Kafka uses ZooKeeper to manage its distributed state. To start ZooKeeper, open a new terminal window and navigate to the Kafka installation directory. Then, run the following command:

```shell
cd kafka_2.13-2.8.0
bin/zookeeper-server-start.sh config/zookeeper.properties
```

![Read Kafka Message]({{ site.url }}/assets/Kafka/kgo1.png){:class="img-responsive"}

This command starts the ZooKeeper service, which is required by Kafka to run.

### Step 3: Start Kafka

To start Kafka, open a new terminal window and navigate to the Kafka installation directory. Then, run the following command:

```shell
bin/kafka-server-start.sh config/server.properties
```

This command starts the Kafka broker service, which is responsible for handling incoming and outgoing data streams.

### Step 4: Create a Topic

To create a topic in Kafka, open a new terminal window and navigate to the Kafka installation directory. Then, run the following command:

```shell
bin/kafka-topics.sh --create --topic mytopic --bootstrap-server localhost:9092
```

![Read Kafka Message]({{ site.url }}/assets/Kafka/kgo2.png){:class="img-responsive"}

Note: Replace mytopic with the name of your topic.

### Step 5: Check if Kafka is installed

To check if Kafka is installed and running properly, open a new terminal window and navigate to the Kafka installation directory. Then, run the following command:

```shell
bin/kafka-topics.sh --list --bootstrap-server localhost:9092
```

![Read Kafka Message]({{ site.url }}/assets/Kafka/kgo3.png){:class="img-responsive"}

This command should display the list of topics that are available in Kafka. If you see your topic in the list, then Kafka is installed and running properly.

# Reading Kafka Response in Go:

To read Kafka response in Go, follow the below steps:

### Step 1: Install Kafka Go client

To use Kafka in Go, we need to install the Kafka Go client. To install it, run the following command:

```shell
go get github.com/segmentio/kafka-go
```

### Step 2: Create a Kafka Reader

To read messages from a Kafka topic, we need to create a Kafka reader. 

The `kafka.ReaderConfig` struct is used to define the configuration options for the Kafka reader.

```go
reader := kafka.NewReader(kafka.ReaderConfig{
    Brokers:  []string{"localhost:9092"},
    Topic:    "mytopic",
    GroupID:  "mygroup",
    MinBytes: 10e3, // 10KB
    MaxBytes: 10e6, // 10MB
})
```

The `Brokers` field is used to specify a list of Kafka broker addresses that the reader will connect to. In this case, we have specified only one broker running on the local machine on port `9092`.

The `Topic` field specifies the Kafka topic that the reader will be reading from. The reader can only consume messages from a single topic at a time.

The `GroupID` field specifies the consumer group that the reader belongs to. A consumer group is a logical grouping of consumers that work together to consume messages from one or more Kafka topics. Each message in a topic can only be consumed by one consumer within a group.

The `MinBytes` and `MaxBytes` fields specify the minimum and maximum number of bytes that the reader will attempt to fetch in a single request to the Kafka broker. These options allow you to tune the performance of the reader based on the characteristics of your messages and network.

Once the reader is set up, we can use the `ReadMessage` method to read messages from the Kafka topic.

### Step 3: Read Kafka Messages

To read messages from Kafka, we use the `ReadMessage()` method of the Kafka reader. The following code reads messages from Kafka:

```go
m, err := reader.ReadMessage(context.Background())
if err != nil {
    log.Printf("Error reading message: %v", err)
    continue
}
```

The `ReadMessage` method blocks until a new message is available on the topic. When a new message is available, it returns a `kafka.Message` struct containing the message key, value, and other metadata. We can then parse the message value as JSON and insert it into a PostgreSQL database, as we have done in the code snippet provided.

# Setting up the PostgreSQL Database Connection:

First setup postgres and create database, user and table. Visit [Complete Guide: Installing, Setting Up, and Removing PostgreSQL in Linux]({{site.url}}/blog/2023/02/23/Complete-Guide-Installing-Setting-Up-and-Removing-PostgreSQL-in-Linux.md) to know more.

I have created table `kafka_mytopic` in database `kafka_responses` and the user is `user_mytopic`.

![Read Kafka Message]({{ site.url }}/assets/Kafka/kgo4.png){:class="img-responsive"}

In the code snippet provided, we see that the PostgreSQL database connection is set up using the `database/sql` package.

```go
db, err := sql.Open("postgres", "postgres://user_mytopic:password@localhost/kafka_responses?sslmode=disable")
if err != nil {
    log.Fatalf("Failed to connect to database: %v", err)
}
defer db.Close()
```

The `sql.Open` function is used to create a new database connection. The first argument to the function specifies the database driver to use, in this case `postgres`.

The second argument to `sql.Open` is the connection string. The connection string specifies the details of the database connection, such as the database name, username, and password. In this case, we are connecting to a PostgreSQL database running on the local machine with the database name `mydatabase`, using a username of `user` and a password of `password`.

The `sslmode=disable` option is used to disable SSL encryption for the database connection. This option should only be used for development and testing purposes, as it can make your database connection less secure.

If the connection is successful, a new `*sql.DB` object is returned. We can use this object to execute SQL queries and transactions against the database.

```go
_, err = db.Exec("INSERT INTO kafka_mytopic (key, value) VALUES ($1, $2)", msg.Key, msg.Value)
if err != nil {
    log.Printf("Error inserting record: %v", err)
    continue
}
```

In the code snippet above, we are using the `db.Exec` method to execute an `INSERT` query against the `kafka_mytopic` table in the PostgreSQL database. The `$1` and `$2` placeholders are used to pass in the `msg.Key` and `msg.Value` values as parameters to the query.

Finally, we use the `defer` keyword to ensure that the database connection is closed when the program exits.

```go
defer db.Close()
```

Closing the database connection is important to ensure that database resources are properly released and to avoid leaking connections.

# Processing Kafka Messages:

The main loop of the program is responsible for reading messages from Kafka, parsing them as JSON, and inserting them into the PostgreSQL database.

```go
for {
    m, err := reader.ReadMessage(context.Background())
    if err != nil {
        log.Printf("Error reading message: %v", err)
        continue
    }

    var msg Message
    err = json.Unmarshal(m.Value, &msg)
    if err != nil {
        log.Printf("Error parsing message: %v", err)
        continue
    }

    _, err = db.Exec("INSERT INTO kafka_mytopic (key, value) VALUES ($1, $2)", msg.Key, msg.Value)
    if err != nil {
        log.Printf("Error inserting record: %v", err)
        continue
    }

    fmt.Printf("Processed message: %v\n", msg)
}
```

In the loop, we use the `reader.ReadMessage` method to read a single message from the Kafka topic. The `context.Background()` parameter is used to create a new context for the message read operation.

If there is an error reading the message, we log the error and continue to the next iteration of the loop.

```go
if err != nil {
    log.Printf("Error reading message: %v", err)
    continue
}
```

Once we have a message, we use the `json.Unmarshal` method to parse the JSON-encoded message into a `Message` struct.

```go
var msg Message
err = json.Unmarshal(m.Value, &msg)
if err != nil {
    log.Printf("Error parsing message: %v", err)
    continue
}
```

If there is an error parsing the message, we log the error and continue to the next iteration of the loop.

We then use the `db.Exec` method to insert the message into the PostgreSQL database.

```go
_, err = db.Exec("INSERT INTO kafka_mytopic (key, value) VALUES ($1, $2)", msg.Key, msg.Value)
if err != nil {
    log.Printf("Error inserting record: %v", err)
    continue
}
```

If there is an error inserting the message into the database, we log the error and continue to the next iteration of the loop.

Finally, we log that we have processed the message.

```go
fmt.Printf("Processed message: %v\n", msg)
```

If there are no more messages to read from the Kafka topic, the loop will continue running, waiting for new messages to arrive.

In summary, the main loop of the program reads messages from Kafka, parses them as JSON, and inserts them into the PostgreSQL database. The program handles errors during message processing by logging them and continuing to the next iteration of the loop.

# Sending Messages to Kafka Topic

In addition to reading messages from a Kafka topic, it's also possible to send messages to a topic using the kafka-go package. Here's an example of how to send a message to a topic using the kafka-console-producer command-line tool:

```shell
$ bin/kafka-console-producer.sh --broker-list localhost:9092 --topic mytopic
> {"key": "foo", "value": "bar"}
> {"key": "baz", "value": "qux"}
```

This command opens an interactive console where you can enter messages to send to the specified Kafka topic. Each message should be a JSON object with "key" and "value" fields.

![Read Kafka Message]({{ site.url }}/assets/Kafka/Kafka-.gif){:class="img-responsive"}

In a Go program, you can use the kafka-go package to send messages to a Kafka topic. Here's an example:

```go
package main

import (
    "context"
    "fmt"

    "github.com/segmentio/kafka-go"
)

func main() {
    writer := kafka.NewWriter(kafka.WriterConfig{
        Brokers:  []string{"localhost:9092"},
        Topic:    "mytopic",
        Balancer: &kafka.LeastBytes{},
    })

    message := kafka.Message{
        Key:   []byte("foo"),
        Value: []byte("bar"),
    }

    err := writer.WriteMessages(context.Background(), message)
    if err != nil {
        fmt.Printf("Error writing message: %v\n", err)
    } else {
        fmt.Println("Message written successfully")
    }

    writer.Close()
}
```

This program creates a new Kafka writer and uses it to send a message to the "mytopic" topic. The message consists of a key and a value, both of which are byte arrays.

The `WriteMessages` method is used to send the message to Kafka. This method takes a context object and one or more `kafka.Message` objects as arguments. If any errors occur during message sending, they are returned by this method.

Finally, the writer is closed to free up any resources it was using.

If you'd like to see the complete code, you can find it on my **Github** repository [here](https://github.com/{{site.github_username}}/Read-Kafka-Response).

# Conclusion

In this blog post, we covered the basics of reading messages from a Kafka topic in Go. We saw how to set up a Kafka reader using the kafka-go package, connect to a PostgreSQL database using the database/sql package, and process messages by parsing them as JSON and inserting them into the database.

We also saw how to send messages to a Kafka topic using the kafka-go package and the kafka-console-producer command-line tool.

With these tools at your disposal, you should be able to build robust, scalable applications that can read and write data to Kafka topics.