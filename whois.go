package phabricatortools

import (
	"errors"
)

// WhoIs calls the conduit user.search method with a single user PHID
func WhoIs(PHID string) (User, error) {
	var user User
	var response whoisResponse

	if PHID == "" {
		return user, errors.New("No user specified")
	}

	connection, err := dialViaCmdLine()

	if err != nil {
		return user, err
	}

	err = connection.Call("user.search", &whoisRequest{Constraints: whoisRequestConstraints{PHIDS: []string{PHID}}}, &response)
	if err != nil {
		return user, err
	}

	if len(response.Data) != 1 {
		return user, errors.New("Didn't get 1 result")
	}
	user = response.Data[0].User

	return user, nil
}
