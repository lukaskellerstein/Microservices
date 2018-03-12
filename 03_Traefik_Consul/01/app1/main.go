package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	consulapi "github.com/hashicorp/consul/api"
)

type Article struct {
	Id      int    `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	fmt.Println("Endpoint Hit: returnAllArticles")

	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w, "Key: "+key)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Everything seems fine!")
}

//***********************************************
// GORILLA MUX package
//***********************************************

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/ping", ping)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(port(), myRouter))
}

func main() {
	fmt.Println("App1 Running . . .")

	registerServiceWithConsul()

	handleRequests()
}

func registerServiceWithConsul() {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println(err)
	}

	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "myapp1"
	registration.Name = "myapp1"

	address := hostname()
	registration.Address = address
	port, _ := strconv.Atoi(port()[1:len(port())])
	registration.Port = port

	consul.Agent().ServiceRegister(registration)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "10011"
	}
	return ":" + port
}

func hostname() string {
	hostname, _ := os.Hostname()
	return hostname
}
