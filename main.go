package main

import (
	"fmt"

	"./lib"
)

func main() {

	// get files in target folder maching pattern
	var folder = "./files/"
	var pattern = "*"

	files, err := lib.ScanFolder(folder, pattern, false)
	if err != nil {
		panic(err)
	}

	fmt.Println(files)
}
