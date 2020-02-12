package phabricatortools

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

type whoisRequestConstraints struct {
	PHIDS []string `json:"phids"`
}

type whoisRequest struct {
	Constraints whoisRequestConstraints `json:"constraints"`
}

type whoisResponseData struct {
	User `json:"fields"`
}

type whoisResponse struct {
	Data []whoisResponseData `json:"data"`
}
