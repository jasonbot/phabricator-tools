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

// Status represents a maniphest/differential task status
type Status struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Closed  bool   `json:"closed"`
	Special string `"special",omitempty`
}

type emptyRequest struct{}

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

type maniphestStatusSearchResponse struct {
	Statuses []Status `json:"data"`
}
