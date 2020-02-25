package phabricatortools

// WhoAmI calls the conduit user.whoami method
func WhoAmI() (User, error) {
	var user User
	connection, err := dialViaCmdLine()

	if err != nil {
		return user, err
	}

	err = connection.Call("user.whoami", &emptyRequest{}, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}
