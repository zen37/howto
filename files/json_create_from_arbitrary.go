package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	m := make(map[string]interface{})

	m["int"] = 1
	m["string"] = "arbitrary types, there is no struct for json creation"
	m["boolean"] = true

	for k, v := range m {
		fmt.Println(k, " = ", v)
	}

	jsonData, err := json.Marshal(m)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(jsonData))

	err = ioutil.WriteFile("file_from_map.json", jsonData, 0644)

}
