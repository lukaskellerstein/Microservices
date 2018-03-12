package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/yosssi/gmq/mqtt"
	"github.com/yosssi/gmq/mqtt/client"
)

//***********************************
// Subscribe
//***********************************

//Mqtt url
var MqttUrl = "localhost:1883"

func main() {
	// Set up channel on which to send signal notifications.
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill)

	clientID := "client2"
	mqttTopic := "test/topic1"

	//--------------------------------------------------
	//MQTT ---------------------------------------------
	//--------------------------------------------------

	// Create an MQTT Client.
	cli := client.New(&client.Options{
		// Define the processing of the error handler.
		ErrorHandler: func(err error) {
			//low-level exception logging
			fmt.Println(err)
			log.Fatalln(err)
			os.Exit(1) // Exit a program
		},
	})

	// Terminate the Client.
	defer cli.Terminate()

	// Connect to the MQTT Server.
	err2 := cli.Connect(&client.ConnectOptions{
		Network:  "tcp",
		Address:  MqttUrl,
		ClientID: []byte(clientID),
	})
	if err2 != nil {
		//low-level exception logging
		fmt.Println(err2)
		os.Exit(1) // Exit a program
	}

	// Subscribe to topics.
	err2 = cli.Subscribe(&client.SubscribeOptions{
		SubReqs: []*client.SubReq{
			&client.SubReq{
				TopicFilter: []byte(mqttTopic),
				QoS:         mqtt.QoS1,
				// Define the processing of the message handler.
				Handler: processMessage,
			},
		},
	})
	if err2 != nil {
		//low-level exception logging
		fmt.Println(err2)
		os.Exit(1) // Exit a program
	}

	// Wait for receiving a signal.
	<-sigc
}

func processMessage(topicName, message []byte) {

	value := string(message)
	fmt.Println(value)

}
