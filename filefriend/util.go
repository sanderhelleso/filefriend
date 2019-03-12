package filefriend

import (
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
