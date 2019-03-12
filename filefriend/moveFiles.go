package filefriend

import (
	"fmt"
	"os"
)

// MoveFiles moves all the files in given slice to destination
// - Uses the past in destination paramter to select folder to move to
//
// returns the updated slice containing all the files
// moved to the destination golder, if error return
// the potensial error that could occur during move
func MoveFiles(files []*File, dest string) ([]*File, error) {

	for _, file := range files {

		newPath := SanitizePath(dest) + file.name + file.extension
		fmt.Println(newPath)
		err := os.Rename(file.path, newPath)
		if err != nil {
			return nil, err
		}
	}

	return files, nil
}
