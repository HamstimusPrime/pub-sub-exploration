package main

import (
	"fmt"
	"os"
	"os/signal"
	pubsub "github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	routing "github.com/bootdotdev/learn-pub-sub-starter/internal/routing"

	amqp "github.com/rabbitmq/amqp091-go"
	
)

func main() {
	fmt.Println("Starting Peril server...")
	conctString := "amqp://guest:guest@localhost:5672/"
	connection, err := amqp.Dial(conctString)
	if err != nil {
		fmt.Println("unable to establish connection to server...")
	}
	defer connection.Close()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan //this is a technique to block the execution of the code below until a value has been passed into signalChan
	fmt.Println("exiting program...")

	channel,err := connection.Channel()
	if err != nil {
		fmt.Println("unable to establish create connection channel...")
		return
	}
	err = pubsub.PublishJSON(channel,
		routing.ExchangePerilDirect, 
		routing.PauseKey,
		routing.PlayingState{
			IsPaused: true,
		},
	)
	if err != nil {
		fmt.Println("unable to pusblish message")
		return
	}


}