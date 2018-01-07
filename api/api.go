package api

import (
	"github.com/ChicagoDSA/DSA-Events/payloads"

	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dgraph-io/dgraph/client"
	protosAPI "github.com/dgraph-io/dgraph/protos/api"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Handles all GraphQL+_ queries
func QueryHandler(c *gin.Context) {
	log := c.MustGet("log").(*logrus.Logger).WithField("api", "queryHandler")
	dGraphClient := c.MustGet("dGraphClient").(*client.Dgraph)

	txn := dGraphClient.NewTxn()
	defer txn.Discard(context.Background())

	query, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.WithError(err).Fatal("Error reading query from request body.")
	}

	resp, err := txn.Query(context.Background(), string(query))
	if err != nil {
		log.WithError(err).Fatal("Error with GraphQL+- query.")
	}

	var root payloads.EventQuery

	err = json.Unmarshal(resp.Json, &root)
	if err != nil {
		log.WithError(err).Fatal("Error umarshalling Dgraph query response.")
	}

	err = txn.Commit(context.Background())
	if err != nil {
		log.WithError(err).Fatal("Error commiting query transaction.")
	}

	c.JSON(http.StatusOK, root.Event)
}

// Handles all GraphQL+_ mutations
func MutationHandler(c *gin.Context) {
	log := c.MustGet("log").(*logrus.Logger).WithField("api", "mutationHandler")
	dGraphClient := c.MustGet("dGraphClient").(*client.Dgraph)

	txn := dGraphClient.NewTxn()
	defer txn.Discard(context.Background())

	eventRequest := &payloads.EventRequest{}
	err := c.BindJSON(eventRequest)
	if err != nil {
		log.WithError(err).Fatal("Error unmarshalling mutation request body into Event object.")
	}

	eventMutation := &protosAPI.Mutation{}

	eventJson, err := json.Marshal(eventRequest)
	if err != nil {
		log.WithError(err).Fatal("Error marshalling mutation request into JSON.")
	}

	eventComparatorData := payloads.EventRequest{Uid: eventRequest.Uid}
	evenComparator, _ := json.Marshal(eventComparatorData)
	if bytes.Equal(eventJson, evenComparator) {
		log.Warn("Deleting node.")
		eventMutation.DeleteJson = eventJson
	} else {
		log.Warn("Creating/Updating node.")
		eventMutation.SetJson = eventJson
	}

	// Send mutation
	resp, err := txn.Mutate(context.Background(), eventMutation)
	if err != nil {
		log.WithError(err).Fatal("Error with GraphQL+- mutation.")
	}

	// Commit mutation
	err = txn.Commit(context.Background())
	if err != nil {
		log.WithError(err).Fatal("Error commiting mutation transaction.")
	}

	c.JSON(http.StatusOK, resp.Uids)
}

// Handles all GraphQL+_ alterations (usually handled by an admin)
func AlterationHandler(c *gin.Context) {
	log := c.MustGet("log").(*logrus.Logger).WithField("api", "alterationHandler")
	dGraphClient := c.MustGet("dGraphClient").(*client.Dgraph)

	query, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.WithError(err).Fatal("Error reading from alteration request body.")
	}

	err = dGraphClient.Alter(context.Background(), &protosAPI.Operation{Schema: string(query)})
	if err != nil {
		log.WithError(err).Fatal("Error with GraphQL+- alteration.")
	}
}
