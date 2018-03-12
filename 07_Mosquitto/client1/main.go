package main

import (
	"fmt"
	"log"
	"os"

	"github.com/yosssi/gmq/mqtt"
	"github.com/yosssi/gmq/mqtt/client"
)

//***********************************
// Publish
//***********************************

//Mqtt url
var MqttUrl = "localhost:1883"

func main() {
	clientID := "client1"
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

	// Publish a message.
	err := cli.Publish(&client.PublishOptions{
		QoS:       mqtt.QoS1,
		TopicName: []byte(mqttTopic),
		Message:   []byte("testMessage"),
	})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
