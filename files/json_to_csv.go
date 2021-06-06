package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Employee struct {
	ID             string `json:"id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary string `json:"employee_salary"`
	EmployeeAge    string `json:"employee_age"`
}
type Employees struct {
	Employee []Employee `json:"employees"`
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("samples/employees.json")
	// Handle error then handle it
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("successfully opened employees.json")
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	//fmt.Print(byteValue)
	//initialize struct
	var employees []Employee
	// jsonFile's content into 'employees' which we defined above
	err = json.Unmarshal(byteValue, &employees)
	if err != nil {
		fmt.Println(err)
	}

	csvFile, err := os.Create("samples/employees.csv")

	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)

	header := []string{"name", "salary", "age"}
	writer.Write(header)

	for _, empl := range employees {
		var row []string
		row = append(row, empl.EmployeeName)
		row = append(row, empl.EmployeeSalary)
		row = append(row, empl.EmployeeAge)
		writer.Write(row)
	}

	// remember to flush!
	writer.Flush()
}
