package payloads

type ContactRequest struct {
	Uid         string `json:"uid,omitempty"`
	Name        string `json:"name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Email       string `json:"email,omitempty"`
	Facebook    string `json:"facebook,omitempty"`
	Twitter     string `json:"twitter,omitempty"`
}

type HostingChapterRequest struct {
	Uid     string         `json:"uid,omitempty"`
	Title   string         `json:"title,omitempty"`
	State   string         `json:"state,omitempty"`
	City    string         `json:"city,omitempty"`
	Contact ContactRequest `json:"contact,omitempty"`
}

type LocationRequest struct {
	Uid     string `json:"uid,omitempty"`
	Name    string `json:"name,omitempty"`
	State   string `json:"state,omitempty"`
	City    string `json:"city,omitempty"`
	ZipCode string `json:"zip_code,omitempty"`
}

type EventRequest struct {
	Uid            string                 `json:"uid,omitempty"`
	HostingChapter *HostingChapterRequest `json:"hosting_chapter,omitempty"`
	Name           string                 `json:"name,omitempty"`
	Time           string                 `json:"time,omitempty"`
	Description    string                 `json:"description,omitempty"`
	Date           string                 `json:"date,omitempty"`
	Location       *LocationRequest       `json:"location,omitempty"`
}
