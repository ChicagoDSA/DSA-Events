package payloads

type Location struct {
	Type        string    `json:"type,omitempty"`
	Coordinates []float64 `json:"coordinates,omitempty"`
}

type ContactRequest struct {
	UID         string `json:"uid,omitempty"`
	Name        string `json:"name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Email       string `json:"email,omitempty"`
	Facebook    string `json:"facebook,omitempty"`
	Twitter     string `json:"twitter,omitempty"`
}

type ChapterRequest struct {
	UID      string          `json:"uid,omitempty"`
	Name     string          `json:"name,omitempty"`
	Location Location        `json:"location,omitempty"`
	Contact  *ContactRequest `json:"contact,omitempty"`
}

type WorkingGroupRequest struct {
	UID         string `json:"uid,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type EventRequest struct {
	UID          string               `json:"uid,omitempty"`
	Name         string               `json:"name,omitempty"`
	Time         string               `json:"time,omitempty"`
	Description  string               `json:"description,omitempty"`
	Location     Location             `json:"location,omitempty"`
	WorkingGroup *WorkingGroupRequest `json:"working_group,omitempty"`
	Chapter      *ChapterRequest      `json:"chapter,omitempty"`
}
