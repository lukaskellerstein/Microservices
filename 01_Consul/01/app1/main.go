package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	log.Fatal(http.ListenAndServe(":10011", myRouter))
}

func main() {
	fmt.Println("App1 Running . . .")
	handleRequests()
}
