package phabricatortools

// GetRepositories returns a list of repositories
func GetRepositories() ([]DiffusionRepository, error) {
	connection, err := dialViaCmdLine()

	if err != nil {
		return nil, err
	}

	first := true
	request := diffusionRepositorySearch{Attachments: map[string]bool{"uris": true}}
	searchResponse := diffusionRepositorySearchResult{}
	results := []DiffusionRepository{}

	for first || searchResponse.After.After != request.After {

		if !first && searchResponse.After.After != "" {
			request.After = searchResponse.After.After
		}

		err := connection.Call("diffusion.repository.search", &request, &searchResponse)
		if err != nil {
			return nil, err
		}

		results = append(results, searchResponse.Data...)
		first = false
	}

	return results, nil
}
