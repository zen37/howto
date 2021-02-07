package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	byteValue, err := ioutil.ReadFile("files/arbitrary.json")
	if err != nil {
		log.Fatal(err)
	}

	var j map[string]interface{}

	// we unmarshal our byteValue which contains our jsonFile's content into the var defined above
	err = json.Unmarshal(byteValue, &j)
	if err != nil {
		log.Fatalln("json malformed:", err)
	}

	fmt.Println(j)
	fmt.Println()

	for k, v := range j {

		switch v := v.(type) {

		case []interface{}:

			for _, v := range v {

				//	fmt.Println(v)

				switch v := v.(type) {

				case map[string]interface{}:

					for i, k := range v {

						fmt.Println(i, ":", k)

					}

				}

			}

		default:

			fmt.Println(k, ":", v)
			//fmt.Println(v)

		}

	}

}
