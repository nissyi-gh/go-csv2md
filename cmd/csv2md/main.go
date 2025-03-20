package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: csv2md <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	header := records[0]
	headerRow := "|"
	for _, h := range header {
		headerRow += h + "|"
	}

	delimiterRow := "|"
	for i := 0; i < len(header); i++ {
		delimiterRow += "---|"
	}

	dataRows := ""
	for _, record := range records[1:] {
		dataRows += "|"
		for _, d := range record {
			dataRows += d + "|"
		}
		dataRows += "\n"
	}

	fmt.Println(headerRow)
	fmt.Println(delimiterRow)
	fmt.Print(dataRows)
}
