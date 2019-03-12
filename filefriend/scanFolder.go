package filefriend

import (
	"path/filepath"
)

// ScanFolder scans the given folder passed in
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

				// append files from nesten folder to
				// main 'fileFolder' slice to be returned
				folderFiles = append(folderFiles, filesFromFolder...)
			}

		} else {

			// if file is not a dir, get information
			// and add file to slice of files to be returned

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
	}

	return folderFiles, err
}
