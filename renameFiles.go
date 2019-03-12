package filefriend

import (
	"strconv"
)

// RenameFiles renames all the files in given slice to newName.
// - newName defines a string passed in as param that will replace the name,
// if slice contains more than 1 file, an index will added to the end of name.
//
// - returns the updated slice containing all the files with the new name or potensial error.
func RenameFiles(files []*File, newName string) error {

	for i, file := range files {

		// get new and old names
		newName := file.path + "\\" + newName
		oldName := file.path + "\\" + file.name + file.extension

		// check if file already exists, if true, add counter
		// if not, append file extension to path and change
		if i > 0 {
			newName = newName + strconv.Itoa(i) + file.extension
		} else {
			newName = newName + file.extension
		}

		// rename file (from -> to)
		err := Move(oldName, newName)
		if err != nil {
			return err
		}

		// if no errors, get new updated file info
		updatedFileInfo, err := GetFileInfo(newName)
		if err != nil {
			return err
		}

		// set updated file info at old file
		*file = *updatedFileInfo
	}

	return nil
}
