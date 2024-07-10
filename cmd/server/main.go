package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		panic(envErr)
	}

	connString := os.Getenv("RABBITMQ_CONNECTION_STRING")

	conn, connErr := amqp.Dial(connString)

	if connErr != nil {
		panic(connErr)
	}

	defer conn.Close()

	fmt.Println("Successfully connected to RabbitMQ", connString)

	channel, channelErr := conn.Channel()

	if channelErr != nil {
		panic(channelErr)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
}
