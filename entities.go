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
	Special string `"special,omitempty"`
	Color   string `json:"color,omitempty"`
}

// Priority represents the priority of a task
type Priority struct {
	Value uint   `json:"value"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type taskDescription struct {
	Raw string `json:"raw"`
}

// Task represents a maniphest task
type Task struct {
	PHID         string          `json:"phid,omitempty"`         // Not actually in the struct, will need to be populated by the consumer
	ID           uint            `json:"id,omitempty"`           // Not actually in the struct, will need to be populated by the consumer
	Name         string          `json:"name,omitempty"`         // string - The title of the task.
	Description  taskDescription `json:"description,omitempty"`  // remarkup - The task description.
	AuthorPHID   string          `json:"authorPHID,omitempty"`   // phid - Original task author.
	OwnerPHID    string          `json:"ownerPHID,omitempty"`    // phid? - Current task owner, if task is assigned.
	Status       Status          `json:"status,omitempty"`       // map<string, wild> - Information about task status.
	Priority     Priority        `json:"priority,omitempty"`     // map<string, wild> - Information about task priority.
	Points       string          `json:"points,omitempty"`       // points - Point value of the task.
	Subtype      string          `json:"subtype,omitempty"`      // string - Subtype of the task.
	CloserPHID   string          `json:"closerPHID,omitempty"`   // phid? - User who closed the task, if the task is closed.
	DateClosed   uint            `json:"dateClosed,omitempty"`   // int? - Epoch timestamp when the task was closed.
	SpacePHID    string          `json:"spacePHID,omitempty"`    // phid? - PHID of the policy space this object is part of.
	DateCreated  uint            `json:"dateCreated,omitempty"`  // int - Epoch timestamp when the object was created.
	DateModified uint            `json:"dateModified,omitempty"` // int - Epoch timestamp when the object was last updated.
}

type emptyRequest struct{}

type cursoredRequest struct {
	After string `json:"after,omitempty"`
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

type maniphestStatusSearchResponse struct {
	Statuses []Status `json:"data"`
}

type maniphestTaskSearchConstraints struct {
	Assigned []string `json:"assigned,omitempty"`
}

type maniphestTaskSearch struct {
	Constraints maniphestTaskSearchConstraints `json:"constraints,omitempty"`
	cursoredRequest
}

type maniphestTaskSearchData struct {
	ID   uint   `json:"id"`
	PHID string `json:"phid"`
	Task Task   `json:"fields"`
}

type maniphestTaskSearchResults struct {
	Data []maniphestTaskSearchData `json:"data"`
	cursoredRequest
}

type diffusionRepositorySearchAttachments map[string]bool

type diffusionRepositorySearch struct {
	Attachments diffusionRepositorySearchAttachments `json:"attachments"`
	cursoredRequest
}

// DiffusionRepositoryFields holds basic information about a repository
type DiffusionRepositoryFields struct {
	Name          string `json:"name"`
	VCS           string `json:"vcs"`
	Callsign      string `json:"callsign"`
	ShortName     string `json:"shortName"`
	DefaultBranch string `json:"defaultBranch"`
}

// DiffusionRepoURI holds the clone URIs for the repo
type DiffusionRepoURI struct {
	Raw       string `json:"raw"`
	Effective string `json:"effective"`
}

// DiffusionRepoBuiltin holds what this URI is for
type DiffusionRepoBuiltin struct {
	Protocol   string `json:"protocol"`
	Identifier string `json:"identifier"`
}

// DiffusionURIFields holds the actual URI info
type DiffusionURIFields struct {
	URI     DiffusionRepoURI     `json:"uri"`
	Builtin DiffusionRepoBuiltin `json:"builtin"`
}

// DiffusionURI is the actual URI object, doubly nested for some reason
type DiffusionURI struct {
	Fields DiffusionURIFields `json:"fields"`
}

// DiffusionURIsAttachments is because the JSON struct is way too deep
type DiffusionURIsAttachments struct {
	URIs []DiffusionURI `json:"uris"`
}

// DiffusionAttachments holds the requested attachments
type DiffusionAttachments struct {
	URIs DiffusionURIsAttachments `json:"uris"`
}

// DiffusionRepository represents a repository
type DiffusionRepository struct {
	ID          uint64                    `json:"id"`
	Fields      DiffusionRepositoryFields `json:"fields"`
	Attachments DiffusionAttachments      `json:"attachments"`
}

type diffusionRepositorySearchResult struct {
	Data  []DiffusionRepository `json:"data"`
	After cursoredRequest       `json:"cursor"`
}
