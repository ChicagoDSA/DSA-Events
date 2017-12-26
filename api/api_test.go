package api

import (
	"testing"
	//"github.com/gin-gonic/gin"
	//"github.com/stretchr/testify/assert"
)

func init() {
	// Mock some sample data
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		signal.Notify(c, os.Kill)
		_ = (<-c)

		logger.Info("Shutting down...")

		// TODO: Delete mocked up data

		os.Exit(0)
	}()

	// TODO: Create dummy data for test cases
}

func TestQueryHandler(t *testing.T) {

}

func TestMutationHandler(t *testing.T) {

}

func TestAlterationHandler(t *testing.T) {

}
