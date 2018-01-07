package main

import (
	"github.com/ChicagoDSA/DSA-Events/api"
	"github.com/ChicagoDSA/DSA-Events/payloads"

	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dgraph-io/dgraph/client"
	protosAPI "github.com/dgraph-io/dgraph/protos/api"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	logger       *logrus.Logger
	dGraphClient *client.Dgraph
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	logger = logrus.New()
	logger.Level = logrus.DebugLevel

	// Establish DGraph connection via gRPC
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		logrus.WithError(err).Fatal("Error establishing gRPC connection with DGraph instance.")
	}
	defer conn.Close()

	dgc := protosAPI.NewDgraphClient(conn)
	dGraphClient = client.NewDgraphClient(dgc)

	os.Exit(m.Run())
}

func getRouter() *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("log", logger)
	})
	router.Use(func(c *gin.Context) {
		c.Set("dGraphClient", dGraphClient)
	})

	return router
}

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

func getQueryEventStr(eventUid string) string {
	return `
	{
		Event(func: uid(` + eventUid + `)) {
			uid
			name
			time
			description
			data
			hosting_chapter {
				title
				state
				city
				contact {
					name
					phone_number
					email
					facebook
					twitter
				}
			}
			location {
				name
				state
				city
				zip_code
			}
		}
	}
	`
}

func getDeleteEventStr(eventUid string) string {
	return `
	{
		"uid":"` + eventUid + `"
	}
	`
}

func getUpdateEventStr(eventUid string) string {
	return `
	{
		"uid": "` + eventUid + `",
		"name": "UPDATED NAME"
	}
	`
}

var CreateEventStr string = `
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

func TestMutations(t *testing.T) {
	router := getRouter()

	router.POST("/mutate", api.MutationHandler)
	router.POST("/query", api.QueryHandler)

	var jsonStr = []byte(CreateEventStr)

	req, _ := http.NewRequest("POST", "/mutate", bytes.NewBuffer(jsonStr))
	newEventUid := ""

	// Create new event
	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		resultBody, _ := ioutil.ReadAll(w.Result().Body)

		var f interface{}
		err := json.Unmarshal([]byte(resultBody), &f)
		if err != nil {
			logger.WithError(err).Fatal("Error unmarshalling response from test event creation")
		}
		newEvent := f.(map[string]interface{})

		logger.WithFields(logrus.Fields{
			"newEventId": newEvent["blank-0"].(string),
		}).Info("Created test event.")
		newEventUid = newEvent["blank-0"].(string)

		return statusOK
	})

	// Query for new event
	req_2, _ := http.NewRequest("POST", "/query", bytes.NewBuffer([]byte(getQueryEventStr(newEventUid))))
	testHTTPResponse(t, router, req_2, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		resp, err := ioutil.ReadAll(w.Result().Body)
		if err != nil {
			logger.Error("Error reading response")
		}

		var f []payloads.EventResponse
		err = json.Unmarshal([]byte(resp), &f)

		logger.WithFields(logrus.Fields{
			"uid":        f[0].Uid,
			"eventName":  f[0].Name,
			"statusCode": w.Code,
		}).Info("Queried newly created test event.")

		return statusOK
	})

	// Update event's name
	req_3, _ := http.NewRequest("POST", "/mutate", bytes.NewBuffer([]byte(getUpdateEventStr(newEventUid))))
	testHTTPResponse(t, router, req_3, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		logger.WithFields(logrus.Fields{
			"statusCode": w.Code,
		}).Info("Updated name of newly created test event.")

		return statusOK
	})

	// Query for updated event
	req_4, _ := http.NewRequest("POST", "/query", bytes.NewBuffer([]byte(getQueryEventStr(newEventUid))))
	testHTTPResponse(t, router, req_4, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		resp, err := ioutil.ReadAll(w.Result().Body)
		if err != nil {
			logger.Error("Error reading response")
		}

		var f []payloads.EventResponse
		err = json.Unmarshal([]byte(resp), &f)

		logger.WithFields(logrus.Fields{
			"eventName": f[0].Name,
		}).Info("Queried newly updated test event.")

		return statusOK
	})

	// Delete new event
	req_5, _ := http.NewRequest("POST", "/mutate", bytes.NewBuffer([]byte(getDeleteEventStr(newEventUid))))
	testHTTPResponse(t, router, req_5, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		logger.WithFields(logrus.Fields{
			"statusCode": w.Code,
		}).Info("Deleted newly created test event.")

		return statusOK
	})
}
