package filefriend

// MoveFiles moves all the files in given slice to destination
// - Uses the past in destination paramter to select folder to move to
//
// returns the updated slice containing all the files
// moved to the destination golder, if error return
// the potensial error that could occur during move
func MoveFiles(files []*File, dest string) ([]*File, error) {
	return files, nil
}
