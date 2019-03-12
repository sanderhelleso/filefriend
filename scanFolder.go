// Package filefriend provides a wide variety of functions that combines the os & filepath
// packages into one, single wrapper, that allows for flexible and easy modifications of
// the systems file and folders
package filefriend

import (
	"path/filepath"
)

// ScanFolder scans the given folder passed in.
// Recur (true/false) decides if scan will follows nested folders.
// Pattern matches files to include/exclude:
//   pattern:
//	{ term }
//	term:
//		'*'         matches any sequence of non-Separator characters
//		'?'         matches any single non-Separator character
//		'[' [ '^' ] { character-range } ']'
//					character class (must be non-empty)
//		c           matches character c (c != '*', '?', '\\', '[')
//		'\\' c      matches character c
//
//	character-range:
//		c           matches character c (c != '\\', '-', ']')
//		'\\' c      matches character c
//		lo '-' hi   matches character c for lo <= c <= hi
//
//
// Returns a slice containing all the files scanned, matching the given pattern, or potensial error that occurred
func ScanFolder(folder string, pattern string, recur bool) ([]*File, error) {

	// initialize empty slice to hold file data
	folderFiles := make([]*File, 0)
	folder = SanitizePath(folder)

	// get all files matching pattern and
	// check for potensial errors while reading
	files, err := filepath.Glob(folder + pattern)
	if err != nil {
		return nil, err
	}

	// iterate over files in folder
	for _, file := range files {

		// check if file is folder
		isFolder, _ := IsFolder(file)
		if isFolder {

			// if folder and recur flag is true
			// contiune to scan nesten folder
			if recur {

				// get files from scanned folder by recursion
				filesFromFolder, err := ScanFolder(file+"/", pattern, recur)

				// handle potensial error from nested folder scann
				if err != nil {
					return nil, err
				}

				// append files from nested folder to
				// main 'fileFolder' slice to be returned
				folderFiles = append(folderFiles, filesFromFolder...)
			}

		} else {

			// if file is not a dir, get information
			// and add file to slice of files to be returned
			// if error, return error and nil slice
			fileInfo, err := GetFileInfo(file)
			if err != nil {
				return nil, err
			}

			folderFiles = append(folderFiles, fileInfo)
		}
	}

	return folderFiles, err
}
