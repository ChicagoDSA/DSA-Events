package main

import (
	"github.com/ChicagoDSA/DSA-Events/api"

	"bytes"
	"encoding/json"
	"fmt"
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

var NewEventStr string = `
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

func TestCreateEvent(t *testing.T) {
	router := getRouter()

	router.POST("/mutate", api.MutationHandler)

	var jsonStr = []byte(NewEventStr)

	req, _ := http.NewRequest("POST", "/mutate", bytes.NewBuffer(jsonStr))

	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		resultBody, _ := ioutil.ReadAll(w.Result().Body)

		var f interface{}
		err := json.Unmarshal([]byte(resultBody), &f)
		if err != nil {
			logger.WithError(err).Fatal("Error unmarshalling response from test event creation")
		}
		newEvent := f.(map[string]interface{})

		fmt.Printf("New event id: " + newEvent["blank-0"].(string) + "\n")

		return statusOK
	})
}