package payloads

type ContactRequest struct {
	Uid         string `json:"uid,omitempty"`
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	Facebook    string `json:"facebook,omitempty"`
	Twitter     string `json:"twitter,omitempty"`
}

type ChapterRequest struct {
	Uid     	string         `json:"uid,omitempty"`
	Name     	string         `json:"name,omitempty"`
	Location 	*LocationRequest`json:"location,omitempty"`
	Contact 	*ContactRequest `json:"contact,omitempty"`
}

type WorkingGroupRequest struct {
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type LocationRequest struct {
	Type string `json:"type,omitempty"`
	Coordinates [2]float64  `json:"coordinates,omitempty"`
}




type EventRequest struct {
	Uid            string				`json:"uid,omitempty"`
	Name           string				`json:"name,omitempty"`
	Description    string				`json:"description,omitempty"`
	Location       *LocationRequest 	`json:"location,omitempty"`
	Time           string				`json:"time,omitempty"`
	WorkingGroup   *WorkingGroupRequest `json:"working_group,omitempty"`
	ChapterRequest *ChapterRequest		`json:"chapter,omitempty"`
}