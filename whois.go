package phabricatortools

import (
	"errors"
)

// User represents a user in the Phabricator system
type User struct {
	PHID     string   `json:"phid"`
	UserName string   `json:"userName"`
	RealName string   `json:"realName"`
	Email    string   `json:"primaryEmail"`
	Image    string   `json:"image"`
	URI      string   `json:"uri"`
	Roles    []string `json:"roles"`
}

type requestConstraints struct {
	PHIDS []string `json:"phids"`
}

type whoisRequest struct {
	Constraints requestConstraints `json:"constraints"`
}

type responseData struct {
	User `json:"fields"`
}

type whoisResponse struct {
	Data []responseData `json:"data"`
}

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

	err = connection.Call("user.search", &whoisRequest{Constraints: requestConstraints{PHIDS: []string{PHID}}}, &response)
	if err != nil {
		return user, err
	}

	if len(response.Data) != 1 {
		return user, errors.New("Didn't get 1 result")
	}
	user = response.Data[0].User

	return user, nil
}
