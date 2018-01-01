package main

import (
	"github.com/ChicagoDSA/DSA-Events/api"

	"flag"
	"os"
	"os/signal"

	"github.com/dgraph-io/dgraph/client"
	protosAPI "github.com/dgraph-io/dgraph/protos/api"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	host string = "0.0.0.0"
	port string = "5000"

	log string = "debug"

	grpcHost string = "localhost"
	grpcPort string = "9080"
)

func init() {
	flag.StringVar(&host, "host", host, "")
	flag.StringVar(&port, "port", port, "")

	flag.StringVar(&log, "log", log, "")

	flag.StringVar(&grpcHost, "grpcHost", grpcHost, "")
	flag.StringVar(&grpcPort, "grpcPort", grpcPort, "")
}

func setUpRouter(logger *logrus.Logger, dGraphClient *client.Dgraph) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
	})

	router.Use(func(c *gin.Context) {
		c.Set("log", logger)
	})
	router.Use(func(c *gin.Context) {
		c.Set("dGraphClient", dGraphClient)
	})
	router.POST("/query", api.QueryHandler)
	router.POST("/mutate", api.MutationHandler)
	router.POST("/alter", api.AlterationHandler)
	router.GET("/test", func(c *gin.Context) {
		c.String(200, "test")
	})

	return router
}

func main() {
	flag.Parse()

	logLevel, err := logrus.ParseLevel(log)
	if err != nil {
		logrus.WithError(err).Fatal("Error parsing log level.")
	}

	logger := logrus.New()
	logger.Level = logLevel
	if logLevel != logrus.DebugLevel {
		gin.SetMode(gin.ReleaseMode)
	}
	logger.WithField("level", logLevel.String()).Debug("Log Level Set")

	// Establish DGraph connection via gRPC
	conn, err := grpc.Dial(grpcHost+":"+grpcPort, grpc.WithInsecure())
	if err != nil {
		logrus.WithError(err).Fatal("Error establishing gRPC connection with DGraph instance.")
	}
	defer conn.Close()

	dgc := protosAPI.NewDgraphClient(conn)
	dGraphClient := client.NewDgraphClient(dgc)

	router := setUpRouter(logger, dGraphClient)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		signal.Notify(c, os.Kill)
		_ = (<-c)
		logger.Info("Shutting down...")
		os.Exit(0)
	}()

	logger.WithError(router.Run(host + ":" + port)).Fatal("Error in setting up HTTP server.")

	logger.Info("API is running on port " + port)
}
