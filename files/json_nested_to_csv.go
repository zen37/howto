package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
)

type structRec struct {
	Nested1 Nested1 `json:"nested1Record"`
	Nested2 Nested2 `json:"nested2Record"`
}

type Nested1 struct {
	Field1 string  `json:"adjustmentReasonCode"`
	Field2 float64 `json:"amount"`
	Field3 int64   `json:"cm_receipt_response_ts"`
	Field4 bool    `json:"cm_receipt_sent"`
}

type Nested2 struct {
	Field1 string `json:"cm_eventhub_consumer_group"`
	Field2 int64  `json:"cm_receipt_response_ts"`
	Field3 bool   `json:"cm_receipt_sent"`
}

const filename string = "samples/employees.csv"

var (
	nested1s  []Nested1
	nested2es []Nested2
)

func main() {

	byteValue, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	//	fmt.Println(string(byteValue))

	s := string(byteValue)

	events := strings.Split(s, "\n")

	for i, v := range events {

		if v != "" {
			if json.Valid([]byte(v)) {

				// Parse JSON string into data object
				var data structRec
				err := json.Unmarshal([]byte(v), &data)
				if err != nil {
					log.Fatalln("Error parsing JSON string - %s", err)
				}

				if (Nested1{} != data.Nested1) { // not empty so we have a nested1 record
					//fmt.Printf("Event type is %s\n", data.Nested1.Version)
					//	fmt.Println(data.Nested1.Isreverted, data.Nested1.Invoicelineitemid)
					nested1s = append(nested1s, data.Nested1)
					//	fmt.Println(nest1)

				}

				if (Nested2{} != data.Nested2) { //not empty so we have a nested2 record
					//fmt.Printf("Event type is %s\n", data.Nested2.Version)
					nested2es = append(nested2es, data.Nested2)
				}

			} else {
				fmt.Println("json not valid for event: ", i)
			}
		}

	}

	if nested1s != nil {
		save("nest1")
	}

	if nested2es != nil {
		save("nest2")
	}

}

//func saveNested1(Nested1s []Nested1) {
func save(what string) {
	//	fmt.Println(nest1)

	var f string
	var nest1 Nested1
	var nested2 Nested2

	var elem reflect.Value

	var r []interface{}

	switch what {
	case "nest1":
		f = "nest1.csv"
		//	nest1 = Nested1{}
		elem = reflect.ValueOf(&nest1).Elem()
		for _, v := range nested1s {
			r = append(r, v)
		}
	case "nest2":
		f = "nest2.csv"
		//nested2 = Nested2{}
		elem = reflect.ValueOf(&nested2).Elem()
		for _, v := range nested2es {
			r = append(r, v)
		}
	}
	csvFile, err := os.Create(f)

	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)

	header := []string{}

	//	nest1 := Nested1{}
	//elem = reflect.ValueOf(&nest1).Elem()

	for i := 0; i < elem.NumField(); i++ {
		structFieldName := elem.Type().Field(i).Name
		header = append(header, structFieldName)
	}

	writer.Write(header)

	for _, v := range r { //nested1s {
		var row []string

		if val, ok := v.(Nested1); ok {
			nest1 = val
		}

		if val, ok := v.(Nested2); ok {
			nested2 = val
		}

		//	fmt.Println("v is", v)

		for i := 0; i < elem.NumField(); i++ {
			//varName := elem.Type().Field(i).Name

			//varType := elem.Type().Field(i).Type
			varValue := elem.Field(i).Interface()
			//	fmt.Printf("%v %v %v\n", varName, varType, varValue)
			//	fmt.Printf("%v", varValue)
			s := fmt.Sprintf("%v", varValue) //we need a string
			row = append(row, s)

		}

		//	row = append(row, v.Adjustmentreasoncode)

		writer.Write(row)
	}

	// remember to flush!
	writer.Flush()
}
