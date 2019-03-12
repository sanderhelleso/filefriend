package filefriend

import "os"

// MoveFiles moves all the files in given slice to destination
// - Uses the past in destination paramter to select folder to move to
//
// returns the updated slice containing all the files
// moved to the destination golder, if error return
// the potensial error that could occur during move
func MoveFiles(files []*File, dest string) ([]*File, error) {

	// check if folder exists, create if nt
	if !PathExists(dest) {
		err := os.MkdirAll(dest, 0755)
		if err != nil {
			return nil, err
		}
	}

	for _, file := range files {

		newPath := SanitizePath(dest) + file.name + file.extension
		oldPath := file.path + "\\" + file.name + file.extension
		moved := Move(oldPath, newPath)
		if moved != nil {
			return nil, moved
		}
	}

	return files, nil
}

// Move moves the fies from one directory
// to another, returns err or nil depending
// on successfull move or not
func Move(from string, to string) error {
	err := os.Rename(from, to)
	return err
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
