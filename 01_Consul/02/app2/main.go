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

	registerServiceWithConsul()

	http.HandleFunc("/", handler)
	http.HandleFunc("/ping", ping)
	http.ListenAndServe(port(), nil)
}

func registerServiceWithConsul() {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println(err)
	}

	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "myapp2"
	registration.Name = "myapp2"

	address := hostname()
	registration.Address = address
	port, _ := strconv.Atoi(port()[1:len(port())])
	registration.Port = port

	consul.Agent().ServiceRegister(registration)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "10012"
	}
	return ":" + port
}

func hostname() string {
	hostname, _ := os.Hostname()
	return hostname
}
