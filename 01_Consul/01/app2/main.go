package main

import (
	"fmt"
	"net/http"
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
	http.HandleFunc("/", handler)
	http.HandleFunc("/ping", ping)
	http.ListenAndServe(":10012", nil)
}
