package main

import (
	"fmt"
	"log"
	"net/http"
)

const address string = ":8080"

func hello(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {

	case "GET":
		http.ServeFile(w, r, "form.html")

		printRequestHeaders(w, r)

	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n\n", address)

		printRequestHeaders(w, r)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")

		printRequestHeaders(w, r)
	}
}

func main() {

	http.HandleFunc("/", hello)

	log.Fatalln(http.ListenAndServe(address, nil))

}

func printRequestHeaders(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
			fmt.Printf("%v: %v\n", name, h)
		}
	}
}
