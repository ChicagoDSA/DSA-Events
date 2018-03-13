package payloads

type EventQuery struct {
	Event []EventResponse `json:"Event"`
}

type ContactResponse struct {
	UID         string `json:"uid,omitempty"`
	Name        string `json:"name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Email       string `json:"email,omitempty"`
	Facebook    string `json:"facebook,omitempty"`
	Twitter     string `json:"twitter,omitempty"`
}

type ChapterResponse struct {
	UID      string            `json:"uid,omitempty"`
	Name     string            `json:"name,omitempty"`
	Location Location          `json:"location,omitempty"`
	Contact  []ContactResponse `json:"contact,omitempty"`
}

type WorkingGroupResponse struct {
	UID         string `json:"uid,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type EventResponse struct {
	UID          string                 `json:"uid,omitempty"`
	Name         string                 `json:"name,omitempty"`
	Time         string                 `json:"time,omitempty"`
	Description  string                 `json:"description,omitempty"`
	Location     Location               `json:"location,omitempty"`
	WorkingGroup []WorkingGroupResponse `json:"working_group,omitempty"`
	Chapter      []ChapterResponse      `json:"chapter,omitempty"`
}
