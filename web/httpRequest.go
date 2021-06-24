package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Message struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

const address string = ":8080"

// curl localhost:8080/v1/test -d '{"name":"Hello"}'
func useProperName(w http.ResponseWriter, r *http.Request) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	fmt.Println("body is:", string(b))

	// Unmarshal
	var msg Message
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
	fmt.Fprintf(w, "output %s\n ", output)

}

func main() {
	http.HandleFunc("/v1/test", useProperName)

	log.Println("starting server on address", address)
	log.Fatalln(http.ListenAndServe(address, nil))
}
