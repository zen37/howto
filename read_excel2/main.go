package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/plandem/xlsx"
)

type sheet xlsx.Sheet

func main() {
	// Create a new XLSX file
	xl := xlsx.New()

	// Open the XLSX file using file name
	xl, err := xlsx.Open("./test.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	sheetNames := xl.SheetNames()

	if !contains(sheetNames, "main") {
		log.Fatalln("there is no \"main\" worksheet")
	}

	if !contains(sheetNames, "tax") {
		log.Fatalln("there is no \"tax\" worksheet")
	}

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

	/*
		// Update the existing XLSX file
		err = xl.Save()
		if err != nil {
			log.Fatal(err)
		}

		// Save the XLSX file under different name
		err = xl.SaveAs("new_file.xlsx")
		if err != nil {
			log.Fatal(err)
		}
	*/
}

func checkSheetName(s *sheet) {

}

func contains(s []string, x string) bool {
	for _, v := range s {
		if x == v {
			return true
		}
	}
	return false
}
