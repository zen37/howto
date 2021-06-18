package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	// import our encoding/json package
	// “encoding/json”
)

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {

	byteValue, err := ioutil.ReadFile("files/users.json")

	if err != nil {
		log.Fatal(err)
	}

	// we initialize our Users array
	var users Users

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example

	for _, user := range users.Users {
		fmt.Println(user.Name)
	}

}
