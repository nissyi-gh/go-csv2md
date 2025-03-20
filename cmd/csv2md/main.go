package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("sample.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()
}
