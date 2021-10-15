package main

import (
	mapper "handler/subscriber/func"
	"log"
	"os"

	amqp "github.com/streadway/amqp"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Broker string `yaml:"broker"`
	} `yaml:"server"`
}

func configReader() Config {
	f, err := os.Open("config.yml")

	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	if err != nil {
		log.Panic(err)
	}

	return cfg
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {

	config := configReader()

	conn, err := amqp.Dial(config.Server.Broker)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"joiner", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			m := new(mapper.Mapper)
			stringBody := string(d.Body)
			m.Map(stringBody)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
