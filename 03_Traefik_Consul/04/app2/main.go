package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	consulapi "github.com/hashicorp/consul/api"
)

// OPEN http://localhost:10012/monkeys

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Everything seems fine!")
}

func main() {
	fmt.Println("App2 Running . . .")

	getConfigurationFromConsul()
	registerServiceWithConsul()

	http.HandleFunc("/", handler)
	http.HandleFunc("/ping", ping)

	fmt.Println(":" + port)
	http.ListenAndServe(":"+port, nil)
}

func registerServiceWithConsul() {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println(err)
	}

	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = microserviceID
	registration.Name = microserviceName

	address := hostname()
	registration.Address = address
	portN, _ := strconv.Atoi(port)
	registration.Port = portN

	consul.Agent().ServiceRegister(registration)
}

var port = ""
var microserviceID = ""
var microserviceName = ""

func getConfigurationFromConsul() {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println(err)
	}

	kv := consul.KV()

	//port
	kvp, _, err := kv.Get("app2/port", nil)
	if err != nil {
		fmt.Println(err)
	} else {
		port = string(kvp.Value)
		fmt.Println(port)
	}

	//microserviceID
	kvid, _, err := kv.Get("app2/id", nil)
	if err != nil {
		fmt.Println(err)
	} else {
		microserviceID = string(kvid.Value)
		fmt.Println(microserviceID)
	}

	//microserviceName
	kvname, _, err := kv.Get("app2/name", nil)
	if err != nil {
		fmt.Println(err)
	} else {
		microserviceName = string(kvname.Value)
		fmt.Println(microserviceName)
	}
}

// func port() string {
// 	port := os.Getenv("PORT")
// 	if len(port) == 0 {
// 		port = "10012"
// 	}
// 	return ":" + port
// }

func hostname() string {
	hostname, _ := os.Hostname()
	return hostname
}
