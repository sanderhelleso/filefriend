package main

import (
	"./filefriend"
)

func main() {

	/*// get files in target folder maching pattern
	var folder = "./files/test/"
	var pattern = "*"
	fmt.Println("\n")

	files, err := filefriend.ScanFolder(folder, pattern, true)
	if err != nil {
		panic(err)
	}

	err = filefriend.RenameFiles(files, "ilovecake")
	if err != nil {
		panic(err)
	}*/

	var testFolder = "files/writeTest"
	var pattern = "*"
	files, err := filefriend.ScanFolder(testFolder, pattern, true)
	if err != nil {
		panic(err)
	}

	err = filefriend.MoveFiles(files, "test/mikefolder", true)
	if err != nil {
		panic(nil)
	}

	/*for _, file := range files {
		fmt.Println(file)
	}
	/*err := filefriend.Create("mikes", "txt", testFolder, "MIKE IS A BAT FILE", 1000)
	if err != nil {
		panic(err)
	}*/

	/*err = filefriend.RenameFiles(files, "ashishi")
	if err != nil {
		panic(err)
	}*/
}
