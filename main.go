package main

import (
	"fmt"

	"./filefriend"
)

func main() {

	// get files in target folder maching pattern
	var folder = "./files/"
	var pattern = "*"
	fmt.Println("\n")

	files, err := filefriend.ScanFolder(folder, pattern, true)
	if err != nil {
		panic(err)
	}

	/*for _, file := range files {
		fmt.Println(file)
	}*/

	movedFiles, err := filefriend.MoveFiles(files, "target")
	if err != nil {
		panic(err)
		fmt.Println(movedFiles)
	}

	//fmt.Println(filefriend.SanitizePath("target"))
}
