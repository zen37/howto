package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type result struct {
	token string `json:"page"`
}

func main() {
	endpoint := "https://url/oauth2/token"
	//	form := &url.Values{} //works
	form := url.Values{} //works
	form.Set("client_id", "...")
	form.Set("client_secret", "...")
	form.Set("grant_type", "client_credentials")
	form.Set("resource", "https://...")

	//data := "client_=...&client_=..."

	client := &http.Client{}
	request, err := http.NewRequest("POST", endpoint, strings.NewReader(form.Encode())) // URL-encoded payload
	//r, err := http.NewRequest("POST", endpoint, strings.NewReader(data) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//r.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	log.Println(response.Status)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))

	var r result
	err = json.Unmarshal(body, &r)

	fmt.Println(r)
	fmt.Println(r.token)
}
