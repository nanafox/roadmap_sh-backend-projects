package task

// ListAll lists all the tasks available in the storage. A status can be
// provided to filter the returned results.
//
// todo => For tasks not yet started.
// in-progress => For tasks that has been started but not completed.
// done => For completed tasks.
func ListAll(cmd []string) (int, error) {
	var status string
	if len(cmd) == 2 {
		status = cmd[1]
	} else {
		status = ""
	}

	tasks, err := getAll(status)

	if err != nil {
		return 1, err
	}

	tasksPrinter(tasks, status)
	return 0, nil
}
