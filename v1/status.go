package phabricatortools

// GetStatuses returns a list of statuses from maniphest.status.search
func GetStatuses() ([]Status, error) {
	connection, err := dialViaCmdLine()

	if err != nil {
		return nil, err
	}

	var response maniphestStatusSearchResponse

	err = connection.Call("maniphest.status.search", &emptyRequest{}, &response)
	if err != nil {
		return nil, err
	}

	return response.Statuses, nil
}

// GetStatusMap returns a map of status types -> statuses from maniphest.status.search
func GetStatusMap() (map[string]Status, error) {
	statuses, err := GetStatuses()

	if err != nil {
		return nil, err
	}

	statusmap := make(map[string]Status)
	for _, item := range statuses {
		statusmap[item.Value] = item
	}

	return statusmap, nil
}
