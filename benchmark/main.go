package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	Yob    int    `json:"yob"`
	Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func ReadJson() (result Users) {

	byteValue, err := ioutil.ReadFile("users.json")

	if err != nil {
		log.Fatal(err)
	}

	// we initialize our Users array
	var users Users

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)

	/*
		for _, user := range users.Users {
			result = result + user.Name
		}

	*/

	return users //result

}

func ReadCsvFile() [][]string {
	f, err := os.Open("users.csv")
	if err != nil {
		log.Fatal("Unable to read input file ", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for ", err)
	}

	return records
}

func main() {
	fmt.Println(ReadJson())
	fmt.Println(ReadCsvFile())
}
