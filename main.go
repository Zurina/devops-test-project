package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Simple struct {
	Name        string
	Description string
	Url         string
}

func SimpleFactory(host string) Simple {
	return Simple{"Hello", "Mathias", host}
}

func handler(w http.ResponseWriter, r *http.Request) {
	simple := Simple{"Hello", "Mathias", r.Host}

	jsonOutput, _ := json.Marshal(simple)

	fmt.Fprintln(w, string(jsonOutput))
}

func main() {
	fmt.Println("Server started")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
