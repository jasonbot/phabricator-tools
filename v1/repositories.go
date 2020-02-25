package phabricatortools

// GetRepositories returns a list of repositories
func GetRepositories() ([]DiffusionRepository, error) {
	connection, err := dialViaCmdLine()

	if err != nil {
		return nil, err
	}

	first := true
	request := diffusionRepositorySearch{Attachments: map[string]bool{"uris": true}}
	results := []DiffusionRepository{}
	after := ""

	for first || after != "" {
		searchResponse := diffusionRepositorySearchResult{}

		if !first {
			request.After = after
		}

		err := connection.Call("diffusion.repository.search", &request, &searchResponse)
		if err != nil {
			return nil, err
		}
		after = searchResponse.Cursor.After

		results = append(results, searchResponse.Data...)
		first = false
	}

	return results, nil
}
