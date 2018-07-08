package main

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/PuloV/ics-golang"
)

func main() {
	logLevel, err := logrus.ParseLevel("debug")
	if err != nil {
		logrus.WithError(err).Fatal("Error parsing log level.")
	}

	logger := logrus.New()
	logger.Level = logLevel
	logger.WithField("level", logLevel.String()).Debug("Log Level Set")

	// Parse json link file
	linkFile, err := ioutil.ReadFile("./sites.json")
	if err != nil {
		logger.WithError(err).Fatal("Failed to read ")
	}

	type link struct {
		Name string `json:"name"`
		Link string `json:"link"`
	}
	var links struct {
		Sites []link `json:"sites"`
	}

	if err = json.Unmarshal(linkFile, &links); err != nil {
		logger.WithError(err).Fatal("Failed parsing links json file.")
	}

	logger.Info("First link parsed from JSON file: "+links.Sites[0].Link)

	client := &http.Client{}
	// Assuming the link to the ICS file
	request, err := http.NewRequest("GET", links.Sites[0].Link, nil)
	if err != nil {
		logger.WithError(err).Fatal("Failed requesting for ICS file from link!")
	}
	resp, err := client.Do(request)
	if err != nil {
		logger.WithError(err).Fatal("Error retrieving ICS file from link!")
	}
	body, err := ioutil.ReadAll(resp.Body)
	//logger.WithField("ICS File", string(body)).Info("ICS file from link")

	f, err := os.Create("cal_data.ics")
	if (err != nil) {
		logger.WithError(err).Fatal("Error creating temporary dir+file for data");
	}
	defer f.Close()

	_, err = f.Write(body)
	if err != nil {
		logger.WithError(err).Fatal("Error writing bytes from ICS file into temp file")
	}

	f.Sync()

	parser := ics.New()

	inputChat := parser.GetInputChan()

	inputChat <- "cal_data.ics"

	parser.Wait()

	cal, err := parser.GetCalendars()

	parser.Wait()
	if err != nil {
		logger.WithError(err).Fatal("Error getting calendars from ICS parser")
	}


	for _, calendar := range cal {
		for _, event := range calendar.GetEvents() {
			logger.WithField("Event Description", event.GetDescription()).Info("Event from calendar")
			logger.WithField("Event Location", event.GetLocation()).Info("Event from calendar")
		}
	}

	// TODO: Create Event struct from Event

	eventMutation := `
	{
		"name":"Name of event",
		"description":"Event's description.",
		"location": {
			"type": "Point", 
			"coordinates": [17.8803304,-245.6662756]
		},
		"time":"2018-05-01T15:30:00Z",
		"working_group": {
			"name": "Working group's name",
			"description": "Description of working group."
		},
		"chapter": {
			"name": "Hosting chapter's name",
			"location": {
				"type": "Point", 
				"coordinates": [18.8803304,-245.6662756]
			},
			"contact": {
				"name":"John Doe",
				"email":"doe@hotmail.com",
				"facebook":"John Doe",
				"twitter":"@johndoe"
			}
		}
	}
	`

}
