package phabricatortools

// GetMyOpenTasks retrieves open tasks
func GetMyOpenTasks() ([]Task, error) {
	connection, err := dialViaCmdLine()

	if err != nil {
		return nil, err
	}

	searchParams := maniphestTaskSearch{}
	var first = true
	var tasks = []Task{}
	after := ""

	user, err := WhoAmI()
	if err != nil {
		return nil, err
	}

	searchParams.Constraints.Assigned = []string{user.PHID}

	for first || after != "" {
		searchResponse := maniphestTaskSearchResults{Data: []maniphestTaskSearchData{}}

		err := connection.Call("maniphest.search", &searchParams, &searchResponse)
		if err != nil {
			return nil, err
		}

		first = false
		if after != "" {
			searchParams.After = after
		}

		for _, task := range searchResponse.Data {
			if task.Task.DateClosed == 0 {
				thisTask := task.Task
				thisTask.PHID = task.PHID
				thisTask.ID = task.ID
				tasks = append(tasks, thisTask)
			}
		}

		after = searchResponse.After

	}

	return tasks, nil
}
