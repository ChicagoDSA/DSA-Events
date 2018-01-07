package payloads

type EventQuery struct {
	Event []EventResponse `json:"Event"`
}

type ContactResponse struct {
	Uid         string `json:"uid,omitempty"`
	Name        string `json:"name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Email       string `json:"email,omitempty"`
	Facebook    string `json:"facebook,omitempty"`
	Twitter     string `json:"twitter,omitempty"`
}

type HostingChapterResponse struct {
	Uid     string            `json:"uid,omitempty"`
	Title   string            `json:"title,omitempty"`
	State   string            `json:"state,omitempty"`
	City    string            `json:"city,omitempty"`
	Contact []ContactResponse `json:"contact,omitempty"`
}

type LocationResponse struct {
	Uid     string `json:"uid,omitempty"`
	Name    string `json:"name,omitempty"`
	State   string `json:"state,omitempty"`
	City    string `json:"city,omitempty"`
	ZipCode string `json:"zip_code,omitempty"`
}

type EventResponse struct {
	Uid             string                   `json:"uid,omitempty"`
	HostingChapters []HostingChapterResponse `json:"hosting_chapter,omitempty"`
	Name            string                   `json:"name,omitempty"`
	Time            string                   `json:"time,omitempty"`
	Description     string                   `json:"description,omitempty"`
	Date            string                   `json:"date,omitempty"`
	Locations       []LocationResponse       `json:"location,omitempty"`
}
