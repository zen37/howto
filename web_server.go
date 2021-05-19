package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":8080"

func main() {
	http.HandleFunc("/", rootHandler)
	log.Fatalln(http.ListenAndServe(port, nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	fmt.Println(r.UserAgent())
}
