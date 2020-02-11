package phabricatortools

import (
	"github.com/uber/gonduit/entities"
	"github.com/uber/gonduit/requests"
)

// WhoAmI calls the conduit user.whoami method
func WhoAmI() (entities.User, error) {
	connection, err := dialViaCmdLine()

	var user entities.User
	if err != nil {
		return user, err
	}

	err = connection.Call("user.whoami", &requests.Request{}, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}
