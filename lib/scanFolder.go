package lib

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	folderNotFound = "*ERROR*: Unable to read folder. *REASON*: Folder not found."
)

// File represent a file in a folder
// containing usefull stats and info
type File struct {
	name      string
	extension string
	folder    string
	path      string
	size      int
	fileNr    int
}

// ScanFolder scans the given directory and
// returns a slice containing all the files
func ScanFolder() []*File {

	// initialize empty slice to hold file data
	folderFiles := make([]*File, 0)

	// get files in target dir
	var dir = "./files/"
	var pattern = "*"

	// get all files matching pattern and
	// check for potensial errors while reading
	files, err := filepath.Glob(dir + pattern)
	if err != nil {
		panic(err)
	}

	// iterate over files in folder
	for i, file := range files {

		// add new File struct with file properties
		folderFiles = append(folderFiles, &File{
			name:      FilenameWithoutExt(file),
			extension: filepath.Ext(file),
			folder:    filepath.Dir(file),
			path:      FullAbsPath(file),
			size:      123,
			fileNr:    i + 1,
		})
	}

	for _, file := range folderFiles {
		fmt.Println(file)
	}

	return folderFiles
}

// FilenameWithoutExt returns the cleaned filename
// without the folderpath and extensipn
func FilenameWithoutExt(file string) string {
	return strings.TrimSuffix(filepath.Base(file), path.Ext(file))
}

// FullAbsPath returns the full absolute path
// from the passed in file
func FullAbsPath(file string) string {

	// get abs path from working dir
	absPath, err := os.Getwd()

	// if error, return 'N/A' as path
	if err != nil {
		return "N/A"
	}

	// return the joined absolute path
	// and dirname of passed in file
	return filepath.Join(absPath, filepath.Dir(file))
}