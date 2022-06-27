package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Greeting struct {
	Id      uint32 `json:"id"`
	Content string `json:"content"`
}

func greetingController(w http.ResponseWriter, r *http.Request) {
	name := "World!"
	val := r.URL.Query()["name"]
	if len(val) > 0 {
		name = val[0]
	}
	reply := Greeting{getCount(), "Hello, " + name}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reply)
	return
}

func main() {
	handler := http.HandlerFunc(greetingController)
	http.HandleFunc("/greeting", handler)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
