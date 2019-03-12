package filefriend

import (
	"io/ioutil"
	"os"
)

// MoveFiles moves all the files in given slice to destination
// - Uses the past in destination paramter to select folder to move to
//
// - returns the updated slice containing all the files
// - moved to the destination golder, if error return
// - cleanup will delete all trailing folders that are empty after move
// the potensial error that could occur during move
func MoveFiles(files []*File, dest string, cleanup bool) error {

	// check if folder exists, create if nt
	if !PathExists(dest) {
		err := os.MkdirAll(dest, 0755)
		if err != nil {
			return err
		}
	}

	for _, file := range files {

		// get new and old paths
		newPath := SanitizePath(dest) + file.name + file.extension
		oldPath := file.path + "\\" + file.name + file.extension

		// move path (from -> to)
		moved := Move(oldPath, newPath)

		// handle potensial error occuring during move
		if moved != nil {
			return moved
		} else {

			// if 'clenup' flag is set to true
			// check if old folder is empty
			// if its empty, remove it
			dirFiles, err := ioutil.ReadDir(file.folder)
			if err != nil {
				return err
			}

			// delete folder if empty after move
			if len(dirFiles) == 0 {
				os.Remove(file.folder)
			}

			// if no errors, get new updated file info
			updatedFileInfo, err := GetFileInfo(newPath)
			if err != nil {
				return err
			}

			// set updated file info at old file
			*file = *updatedFileInfo
		}
	}

	return nil
}
