package main

import (
	"strings"
	"testing"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

func testGetEvent(t *testing.T) {

}

func testGetEvents(t *testing.T) {

}

func testCreateEvent(t *testing.T) {

}

func testUpdateEvent(t *testing.T) {

}

func testDeleteEvent(t *testing.T) {

}

func TestQueryHandler(t *testing.T) {
	router := setUpRouter()

	// TODO: Create mock data
	reqBody := `
	{
		"name":"Created Event 1",
		"date":"08-11-17",
		"time":"14:30",
		"hosting_chapter":{
			"title":"Milwaukee DSA",
			"city":"Milwaukee",
			"state":"WI",
			"contact":{
				"name":"Jeb Bush",
				"phone_number":"123-456-7890",
				"email":"jeb@hotmail.com",
				"facebook":"jebisthebest",
				"twitter":"@jebisgood"
			}
		},
		"description":"Test Description",
		"location":{
			"name":"Location Name!",
			"city":"Milwaukee",
			"state":"WI",
			"zip_code":"12345"
		}
	}
	`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/mutate", strings.NewReader(reqBody))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code, "Successfully created mock data.")
	fmt.Printf("DGraph response: \n%s\n", w.Body.String)


	//t.Run("GetEvent", testGetEvent(t))
	//t.Run("GetEvents", testGetEvents(t))
	
	// TODO: Delete mock data
}

func TestMutationHandler(t *testing.T) {
	//router := setUpRouter()

	// TODO; Create mock data

	t.Run("CreateEvent", testCreateEvent)
	t.Run("UpdateEvent", testUpdateEvent)
	t.Run("DeleteEvent", testDeleteEvent)

	// TODO: Delete mock data
}