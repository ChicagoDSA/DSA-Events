package main

import (
	"github.com/ChicagoDSA/DSA-Events/api"
	"github.com/ChicagoDSA/DSA-Events/auth"

	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/dgraph-io/dgo"
	protosAPI "github.com/dgraph-io/dgo/protos/api"
	"github.com/gin-gonic/gin"
	"github.com/justinas/nosurf"
	"github.com/sirupsen/logrus"
	"github.com/gocolly/colly"
	"google.golang.org/grpc"
)

var (
	host string = "localhost"
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

func setUpRouter(logger *logrus.Logger, dGraphClient *dgo.Dgraph) *gin.Engine {
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
		c.Set("host", host)
		c.Set("port", port)
		c.Set("dGraphClient", dGraphClient)
		c.Set("log", logger)
	})

	router.POST("/query", api.QueryHandler)
	router.POST("/mutate", api.MutationHandler)
	router.POST("/alter", api.AlterationHandler)

	// GitHub OAuth2
	router.GET("/account/github/callback", auth.GithubCallback)
	router.GET("/link/github", auth.GithubInit)

	return router
}

func csrfFailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", nosurf.Reason(r))
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

	// Testing web scraper
	c := colly.NewCollector()

	testLink := "https://duckduckgo.com/"
	logger.Info("Testing Colly WebScraper on: "+testLink)
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		logger.WithFields(logrus.Fields{
			"text": e.Text,
			"link": link,
		}).Info("Link found!")
	})

	c.Visit(testLink)

	// Establish DGraph connection via gRPC
	conn, err := grpc.Dial(grpcHost+":"+grpcPort, grpc.WithInsecure())
	if err != nil {
		logrus.WithError(err).Fatal("Error establishing gRPC connection with DGraph instance.")
	}
	logrus.RegisterExitHandler(func() {
		conn.Close()
	})

	dgc := protosAPI.NewDgraphClient(conn)
	dGraphClient := dgo.NewDgraphClient(dgc)

	router := setUpRouter(logger, dGraphClient)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)
		_ = (<-c)
		logger.Info("Shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := server.Shutdown(ctx)
		if err != nil {
			logger.WithError(server.Shutdown(ctx)).Fatal("Server shutdown")
		}
		os.Exit(0)
	}()

	logger.WithError(router.Run(host + ":" + port)).Fatal("Error in setting up HTTP server.")
}
