package lib

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	folderNotFound = "*ERROR*: Unable to read folder. *REASON*: Folder not found."
)

// File represent a file in a folder
// containing usefull stats and info
type File struct {
	name        string
	extension   string
	folder      string
	path        string
	size        string
	lastChanged string
	fileNr      int
}

// ScanFolder scans the given directory and
// returns a slice containing all the files
func ScanFolder() ([]*File, error) {

	// initialize empty slice to hold file data
	folderFiles := make([]*File, 0)

	// get files in target dir
	var dir = "./files/"
	var pattern = "*"

	// get all files matching pattern and
	// check for potensial errors while reading
	files, err := filepath.Glob(dir + pattern)
	if err != nil {
		return nil, err
	}

	// iterate over files in folder
	for i, file := range files {

		// add new File struct with file properties
		folderFiles = append(folderFiles, &File{
			name:        FilenameWithoutExt(file),
			extension:   filepath.Ext(file),
			folder:      filepath.Dir(file),
			path:        FullAbsPath(file),
			size:        GetFileSize(file),
			lastChanged: GetFileChangedTime(file),
			fileNr:      i + 1,
		})
	}

	for _, file := range folderFiles {
		fmt.Println(file)
	}

	return folderFiles, nil
}

// FilenameWithoutExt returns the cleaned filename
// without the folderpath and extensipn
func FilenameWithoutExt(file string) string {
	return strings.TrimSuffix(filepath.Base(file), path.Ext(file))
}

// FullAbsPath returns the full absolute path
// from the passed in file, if error return 'N/A'
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

// GetFileSize returns the size of the
// passed in file, if error return 'N/A
func GetFileSize(file string) string {

	// get stats from passed in file
	stats, err := os.Stat(file)

	// handle potensial errors
	if err != nil {
		return "N/A"
	}

	// return the file size from stats
	return strconv.Itoa(int(stats.Size()))
}

// GetFileChangedTime returns the timestamp
// of the passed in files last changed time
// if error, return 'N/A'
func GetFileChangedTime(file string) string {

	// get stats from passed in file
	stats, err := os.Stat(file)

	// handle potensial errors
	if err != nil {
		return "N/A"
	}

	// return the modified time as string
	return stats.ModTime().String()
}
