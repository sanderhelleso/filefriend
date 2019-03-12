package main

import (
	"fmt"

	"./lib"
)

func main() {

	// get files in target folder maching pattern
	var folder = "./files/"
	var pattern = "*"
	fmt.Println("\n")

	files, err := lib.ScanFolder(folder, pattern, true)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}
}
