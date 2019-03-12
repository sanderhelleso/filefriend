//filefriend is a convenient wrapper around the os/filepath module for Goprovides a wide variety of functions that combines
//packages into one, single wrapper, allowing for flexible and easy modifications of the systems file and folders.
package filefriend

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

// File represent a file in a folder. Containing useful stats and information about the file.
//
// name: 		The name of the file without extension or path.
//
// extension:	The extension of the file.
//
// folder:		The relative folder path of the file.
//
// path:		The absolute filesystems path to the folder.
//
// size:		The size of the file in bytes in string format.
//
// lastChanged: The files last modified time in string format.
type File struct {
	name        string
	extension   string
	folder      string
	path        string
	size        string
	lastChanged string
}

// FilenameWithoutExt returns the cleaned filename without the folderpath or extension.
func FilenameWithoutExt(file string) string {
	return strings.TrimSuffix(filepath.Base(file), path.Ext(file))
}

// FullAbsPath returns the full absolute path from the passed in file or potensial error that occurred.
func FullAbsPath(file string) (string, error) {

	absPath, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// return the joined absolute path
	// and dirname of passed in file
	return filepath.Join(absPath, filepath.Dir(file)), err
}

// GetSize returns the size of the passed in file or potensial error that occurred.
func GetSize(file string) (string, error) {

	stats, err := os.Stat(file)
	if err != nil {
		return "", err
	}

	// return the file size from stats
	return strconv.Itoa(int(stats.Size())), err
}

// GetChangedTime returns the timestamp of the passed in file's last modified time, or potensial error that occurred.
func GetChangedTime(file string) (string, error) {

	stats, err := os.Stat(file)
	if err != nil {
		return "", err
	}

	// return the modified time as string
	return stats.ModTime().String(), err
}

// IsFolder returns a boolean (true) if the given path is a folder.
// If the path is a file, return boolean (false) and the error with the reason.
func IsFolder(path string) (bool, error) {

	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	// return info if path is folder
	return fileInfo.IsDir(), err
}

// SanitizePath formats the path into system navigatable format.
// If the path is invalid for system, format the path with starting './' and trailing '/'
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

// Move renames and moves the files from one directory,
// to another depending on folder structure input.
// Returns error or nil depending on successful move or rename.
func Move(from string, to string) error {
	err := os.Rename(from, to)
	return err
}

// Create creates 'n' amount of files, with extension, to passed in folder,
// with passed in content, that will be written to each single file.
//
// If file in folder exists, Create will skip to avoid remake of the file.
// Returns an error or nil depending on successful write or fail.
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

// PathExists checks if a given path exists or not.
// Returns a boolean (true/false) depending on case.
func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

// CreateFolder creates a new folder at a given path.
// If the folder does not exist create it.
// Returns an error if the creation failed.
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

// GetFileInfo retrieves the stats & information about a specific file.
// Returns a pointer to the file or potensial error that occurred.
func GetFileInfo(file string) (*File, error) {

	// get stats releated data from file and,
	// handle error occurred during stats opertation
	path, err := FullAbsPath(file)
	if err != nil {
		return nil, err
	}
	size, err := GetSize(file)
	if err != nil {
		return nil, err
	}
	lastChanged, err := GetChangedTime(file)
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
