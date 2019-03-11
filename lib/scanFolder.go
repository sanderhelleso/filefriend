package lib

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
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

// - ScanFolder scans the given folder passed in
// - Uses the passed in pattern to include/exclude files
// - Recur (true/false) decides if scann follows nested folders
//
// returns a slice containing all the files scanned
// matching the given pattern, if any error occured
// the returned values will be nil and occured error
func ScanFolder(folder string, pattern string, recur bool) ([]*File, error) {

	// initialize empty slice to hold file data
	folderFiles := make([]*File, 0)

	// get all files matching pattern and
	// check for potensial errors while reading
	files, err := filepath.Glob(folder + pattern)
	if err != nil {
		return nil, err
	}

	// iterate over files in folder
	for i, file := range files {

		// check if file is folder
		if ()

		// get stats releated data from file
		path, err := FullAbsPath(file)
		size, err := GetSize(file)
		lastChanged, err := GetChangedTime(file)

		// habdle error occured during stats opretation
		if err != nil {
			return nil, err
		}

		// add new File struct with file properties
		folderFiles = append(folderFiles, &File{
			name:        FilenameWithoutExt(file),
			extension:   filepath.Ext(file),
			folder:      filepath.Dir(file),
			path:        path,
			size:        size,
			lastChanged: lastChanged,
			fileNr:      i + 1,
		})
	}

	for _, file := range folderFiles {
		fmt.Println(file)
	}

	return folderFiles, err
}

// FilenameWithoutExt returns the cleaned filename
// without the folderpath and extensipn
func FilenameWithoutExt(file string) string {
	return strings.TrimSuffix(filepath.Base(file), path.Ext(file))
}

// FullAbsPath returns the full absolute path
// from the passed in file or error
func FullAbsPath(file string) (string, error) {

	absPath, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// return the joined absolute path
	// and dirname of passed in file
	return filepath.Join(absPath, filepath.Dir(file)), err
}

// GetSize returns the size of the
// passed in file or potensial error
func GetSize(file string) (string, error) {

	stats, err := os.Stat(file)
	if err != nil {
		return "", err
	}

	// return the file size from stats
	return strconv.Itoa(int(stats.Size())), err
}

// GetChangedTime returns the timestamp
// of the passed in files last changed time
// or potensial error that occured
func GetChangedTime(file string) (string, error) {

	stats, err := os.Stat(file)
	if err != nil {
		return "", err
	}

	// return the modified time as string
	return stats.ModTime().String(), err
}

// IsFolder returns a boolean (true) if
// the given path is a folder, if the
// path is a file, return boolean (false)
// and the error that occured
func IsFolder(path string) (bool, error) {

	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	// return info if path is folder
	return fileInfo.IsDir(), err
}
