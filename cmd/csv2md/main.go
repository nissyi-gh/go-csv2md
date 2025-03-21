package main

import (
	"bufio"
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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	header, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	headerRow := "|"
	for _, h := range header {
		headerRow += h + "|"
	}
	writer.WriteString(headerRow + "\n")

	delimiterRow := "|"
	for i := 0; i < len(header); i++ {
		delimiterRow += "---|"
	}
	writer.WriteString(delimiterRow + "\n")

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		dataRow := "|"
		for _, d := range record {
			dataRow += d + "|"
		}
		writer.WriteString(dataRow + "\n")
	}
}
