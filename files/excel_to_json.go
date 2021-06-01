package main

import (
	"container/list"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/plandem/xlsx"
)

const file_path string = "./test.xlsx"

type sheet xlsx.Sheet

var (
	test      *string
	worksheet *string
)

func init() {

	test = flag.String("t", "1", "test to run")

	worksheet = flag.String("w", "all", "what worksheet to read")

	flag.Parse()

}

func main() {

	// Open the XLSX file using file name
	xl, err := xlsx.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	ReadXlsxFile(xl)

	//fmt.Println(data)

	//parsedJson, _ := json.Marshal(parsedData)

	//fmt.Println(string(parsedJson))

	getMainFields(xl)

	err = validateFile(xl)

	if err != nil {

		log.Fatalln("File is not valid:", err)

	}

}

func validateFile(xl *xlsx.Spreadsheet) error {

	err := checkSheetName(xl)

	if err != nil {
		return err
	}

	return nil

}

func contains(s []string, x string) bool {

	for _, v := range s {

		if x == v {

			return true

		}

	}

	return false

}

func checkSheetName(xl *xlsx.Spreadsheet) error {

	sheetNames := xl.SheetNames()

	if !contains(sheetNames, "main") {

		return errors.New("there is no \"main\" worksheet")

	}

	if !contains(sheetNames, "tax") {

		return errors.New("there is no \"tax\" worksheet")

	}

	return nil

}

func constMainFields() [3]string {

	return [...]string{"field1", "field2", "field3"}

}

//func ReadXlsxFile(filePath string) []map[string]interface{} {

func ReadXlsxFile(xlFile *xlsx.Spreadsheet) {

	switch *worksheet {

	case "t":

		//read tax items

		fmt.Println("----------------- reading tax----------------")

		t := xlFile.SheetByName("tax")

		readSheet(t, *test)

	default:

		log.Fatalln("uknown worksheet")

	}

}

func readSheet(s sheet, test string) {

	//  parsedData := make([]map[string]interface{}, 0)

	header_name := list.New()

	for rows := s.Rows(); rows.HasNext(); {

		row_counter, row := rows.Next()

		// column

		header_iterator := header_name.Front()

		var singleMap = make(map[string]interface{})

		//var sheetTest int

		for cells := row.Cells(); cells.HasNext(); {

			_, _, cell := cells.Next()

			if row_counter == 0 {

				text := cell.String()

				header_name.PushBack(text)

			} else {

				text := cell.String()

				//  if header_iterator.Value.(string) != "Test" {

				singleMap[header_iterator.Value.(string)] = text

				//  }

				header_iterator = header_iterator.Next()

			}

		}

		if row_counter != 0 && len(singleMap) > 0 {

			v := singleMap["#TEST"]

			if v != test {

				continue

			}

			transformValues(singleMap)

		}

		//for k, v := range singleMap {

		//      fmt.Println(k, " = ", v)

		//  }

		if len(singleMap) != 0 {

			f := "test" + test + "_" + *worksheet + ".json"

			writeFile(f, singleMap)

		}

	}

}

func writeFile(filename string, m map[string]interface{}) {

	parsedJson, err := json.MarshalIndent(m, "", " ")

	//fmt.Println(m)

	//parsedJson, err := json.Marshal(m)

	if err != nil {

		log.Fatal(err)

	}

	fmt.Println("JSON data:\n", string(parsedJson))

	err = ioutil.WriteFile(filename, parsedJson, 0644)

	if err != nil {

		log.Fatal(err)

	}

	//  fmt.Println("JSON data:\n", string(parsedJson))

}

func transformValues(m map[string]interface{}) {

	if m["isTest"].(string) == "1" {

		m["isTest"] = true

	} else {

		m["isTest"] = false

	}

	_, ok := m["amount"]

	if ok {

		m["amount"], _ = strconv.ParseFloat(m["amount"].(string), 64)

	}

}

func getMainFields(xl *xlsx.Spreadsheet) {

	/*

	   for sheets := xl.Sheets(); sheets.HasNext(); {

	       _, sheet := sheets.Next()



	       fmt.Println(sheet.Name())



	       // Get row by 0-based index

	       row := sheet.Row(0)

	       fmt.Println(strings.Join(row.Values(), ","))



	       l := len(row.Values())

	       var arr [l]string

	       copy(arr[:], row.Values())

	       fmt.Println(arr)



	   }

	*/

}

func iterate(xl *xlsx.Spreadsheet) {

	// Iterate sheets via iterator

	for sheets := xl.Sheets(); sheets.HasNext(); {

		_, sheet := sheets.Next()

		fmt.Println(sheet.Name())

		// Get cell by 0-based indexes

		cell := sheet.Cell(1, 1)

		fmt.Println(cell.Value())

		// Get row by 0-based index

		row := sheet.Row(0)

		fmt.Println(strings.Join(row.Values(), ","))

		// Get col by 0-based index

		col := sheet.Col(0)

		fmt.Println(strings.Join(col.Values(), ","))

		// Get cell by reference

		cell = sheet.CellByRef("B3")

		fmt.Println(cell.Value())

	}

	// Get sheet by 0-based index

	//sheet := xl.Sheet(0)

}
