package payloads

type ContactResponse struct {
	Uid         string `json:"uid,omitempty"`
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	Facebook    string `json:"facebook,omitempty"`
	Twitter     string `json:"twitter,omitempty"`
}

type ChapterResponse struct {
	Uid     	string         `json:"uid,omitempty"`
	Name     	string         `json:"name,omitempty"`
	Location 	*LocationResponse`json:"location,omitempty"`
	Contact 	[]ContactResponse `json:"contact,omitempty"`
}

type WorkingGroupResponse struct {
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type LocationResponse struct {
	Type string `json:"type,omitempty"`
	Coordinates [2]float64  `json:"coordinates,omitempty"`
}




type EventResponse struct {
	Uid            string				`json:"uid,omitempty"`
	Name           string				`json:"name,omitempty"`
	Description    string				`json:"description,omitempty"`
	Location       *LocationResponse 	`json:"location,omitempty"`
	Time           string				`json:"time,omitempty"`
	WorkingGroup   []WorkingGroupResponse `json:"working_group,omitempty"`
	ChapterResponse []ChapterResponse		`json:"chapter,omitempty"`
}

type GithubUser struct {
	OrganizationsUrl	string	`json:"organizations_url,omitempty"`
}

type GithubOrg struct {
	Login	string	`json:"login,omitempty"`
}

type EventQuery struct {
	Event []EventResponse `json:"Event"`
}