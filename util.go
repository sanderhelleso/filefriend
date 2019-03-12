package filefriend

import (
	"io/ioutil"
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

// SanitizePath formats the path if
// the path is invalid for dir. Paths without
// starting './' will be added and trailing '/'
func SanitizePath(path string) string {

	// check first char
	if path[0] != '/' && path[0] != '\\' {
		path = "./" + path
	}

	// check last char
	if path[len(path)-1] != '/' && path[len(path)-1] != '\\' {
		path = path + "/"
	}

	return path
}

// Move renames and moves the files from one directory
// to another depending on folder structure input,
// returns err or nil depending on successfull move/rename
func Move(from string, to string) error {
	err := os.Rename(from, to)
	return err
}

// Create creates 'n' amount of files with ext to passed in folder
// with passed in content that will be written to each file
// if file in folder exists, Create will skip to avoid remake
// returns err or nil depending on successfull write
func Create(name string, ext string, folder string, content string, amount int) error {

	// check if folder exists, create if not exists
	folder = SanitizePath(folder)
	err := CreateFolder(folder)
	if err != nil {
		return err
	}

	fileName := name
	for i := 0; i < amount; i++ {

		// update file index name
		if i > 0 {
			fileName = name + strconv.Itoa(i)
		}

		// check if path to file that should be created exists
		file := folder + fileName + "." + ext
		exists := PathExists(file)

		// if it does not exists, continue to create file
		if !exists {

			// attempt to create file
			err := ioutil.WriteFile(file, []byte(content), 0755)

			// handle potensial error
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// PathExists checks if a path exists or not
// returns a boolean (true/false) depending
// if path exists or not
func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

// CreateFolder creates a new folder
// if the folder does not exist create it
// if not return the error that folder exist
func CreateFolder(path string) error {

	// check if folder exists, create if not
	if !PathExists(path) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetFileInfo retrieves the stats
// & information about a specific file,
// returns a pointer to the file
func GetFileInfo(file string) (*File, error) {

	// get stats releated data from file
	path, err := FullAbsPath(file)
	size, err := GetSize(file)
	lastChanged, err := GetChangedTime(file)

	// handle error occured during stats opertation
	if err != nil {
		return nil, err
	}

	// add new File struct with file properties
	return &File{
		name:        FilenameWithoutExt(file),
		extension:   filepath.Ext(file),
		folder:      filepath.Dir(file),
		path:        path,
		size:        size,
		lastChanged: lastChanged,
	}, nil
}
