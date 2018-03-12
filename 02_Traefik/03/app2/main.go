package main

import (
	"fmt"
	"net/http"
)

// OPEN http://localhost:10012/monkeys

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":10012", nil)
}
