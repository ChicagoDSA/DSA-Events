package payloads

// Used for dgraph query parsing
type Contact struct {
	Uid            	string         	`json:"uid,omitempty"`
	Name        	string 			`json:"name,omitempty"`
	PhoneNumber 	string 			`json:"phone_number,omitempty"`
	Email       	string 			`json:"email,omitempty"`
	Facebook    	string 			`json:"facebook,omitempty"`
	Twitter     	string 			`json:"twitter,omitempty"`
}

type HostingChapter struct {
	Uid            	string         	`json:"uid,omitempty"`
	Title   		string  		`json:"title,omitempty"`
	State   		string  		`json:"state,omitempty"`
	City    		string  		`json:"city,omitempty"`
	Contact 		[]Contact 		`json:"contact,omitempty"`
}

type Location struct {
	Uid            	string         	`json:"uid,omitempty"`
	Name    		string 			`json:"name,omitempty"`
	State   		string 			`json:"state,omitempty"`
	City    		string 			`json:"city,omitempty"`
	ZipCode 		string 			`json:"zip_code,omitempty"`
}

type Event struct {
	Uid            string         		`json:"uid,omitempty"`
	HostingChapter []HostingChapter 	`json:"hosting_chapter,omitempty"`
	Name           string         		`json:"name,omitempty"`
	Time           string         		`json:"time,omitempty"`
	Description    string         		`json:"description,omitempty"`
	Date           string         		`json:"date,omitempty"`
	Location       []Location       	`json:"location,omitempty"`	
}

type EventQuery struct {
	Event []Event `json:"Event"`
}

// Gin Request Payloads
type GContact struct {
	Uid            	string         	`json:"uid,omitempty"`
	Name        	string 			`json:"name,omitempty"`
	PhoneNumber 	string 			`json:"phone_number,omitempty"`
	Email       	string 			`json:"email,omitempty"`
	Facebook    	string 			`json:"facebook,omitempty"`
	Twitter     	string 			`json:"twitter,omitempty"`
}

type GHostingChapter struct {
	Uid            	string         	`json:"uid,omitempty"`
	Title   		string  		`json:"title,omitempty"`
	State   		string  		`json:"state,omitempty"`
	City    		string  		`json:"city,omitempty"`
	Contact 		GContact 		`json:"contact,omitempty"`
}

type GLocation struct {
	Uid            	string         	`json:"uid,omitempty"`
	Name    		string 			`json:"name,omitempty"`
	State   		string 			`json:"state,omitempty"`
	City    		string 			`json:"city,omitempty"`
	ZipCode 		string 			`json:"zip_code,omitempty"`
}

type GEvent struct {
	Uid            string         	`json:"uid,omitempty"`
	HostingChapter *GHostingChapter 	`json:"hosting_chapter,omitempty"`
	Name           string         	`json:"name,omitempty"`
	Time           string         	`json:"time,omitempty"`
	Description    string         	`json:"description,omitempty"`
	Date           string         	`json:"date,omitempty"`
	Location       *GLocation       	`json:"location,omitempty"`	
}