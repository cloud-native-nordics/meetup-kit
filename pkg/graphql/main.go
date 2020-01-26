package main

import (
	goflag "flag"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/handler"
	"github.com/cloud-native-nordics/stats-api/generated"
	"github.com/cloud-native-nordics/stats-api/handlers"
	"github.com/cloud-native-nordics/stats-api/repositories"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/golang/glog"
	"github.com/rs/cors"
	flag "github.com/spf13/pflag"
)

var port = flag.String("port", "8080", "Application port to use")
var statsURL = flag.String("stats-url", "https://raw.githubusercontent.com/cloud-native-nordics/meetups/master/config.json", "Location of the stats file")
var slackToken = flag.String("slack-token", "", "Slack token to produce invites")
var slackURL = flag.String("slack-url", "https://cloud-native-nordics.slack.com", "URL to the slack community")
var slackCommunity = flag.String("slack-community", "Cloud Native Nordics", "Name of the slack community")

func main() {
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	db := handlers.NewStatsManager(*statsURL)

	statsRepo := repositories.NewStatsRepository(db)
	slackRepo := repositories.NewSlackRepository(*slackToken, *slackURL, *slackCommunity)

	resolver := handlers.NewResolver(statsRepo, slackRepo)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(5 * time.Second))

	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: resolver})))

	glog.V(5).Infof("Connect to http://localhost:%s/ for GraphQL playground", *port)
	glog.Fatalf("Fatal: %s", http.ListenAndServe(":"+*port, router))
}
