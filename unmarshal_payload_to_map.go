package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	raw := []byte(`{
                  "places" : {
                     "london" : 65,
                     "stockholm" : 14,
                     "somecity" : 3,
		     "more": true
                   },
		   "a":500,
		   "xyz": "qwerty"	
                }`)

	dataMap := map[string]interface{}{}

	if err := json.Unmarshal(raw, &dataMap); err != nil {
		log.Fatalf("unmarsal failed: %v", err)
	}
	fmt.Println(dataMap)
	fmt.Printf("%+v\n", dataMap)

}
