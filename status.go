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
