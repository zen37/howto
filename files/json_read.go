// https://medium.com/@IndianGuru/consuming-json-apis-with-go-d711efc1dcf9

//https://www.soberkoder.com/consume-rest-api-go/

//https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7
// /https://tutorialedge.net/golang/consuming-restful-api-with-go/

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
}

type Switch struct {
	On bool `json:"on"`
}

func main() {

	jsonString := `{"name":"ann","age":33,"hobbies":["bike","soccer"]}`
	jsonByte := []byte(jsonString)

	p := person{}

	err := json.Unmarshal(jsonByte, &p)

	if err != nil {
		log.Fatal(err)
	}
	//fmt.Print("%+v\n", p)
	fmt.Printf("%+v\n", p)

	jsonString2 := `{"on": true}`
	jsonByte2 := []byte(jsonString2)

	s := Switch{}

	err = json.Unmarshal(jsonByte2, &s)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(jsonString2)
	fmt.Printf("%+v\n", s)
}
