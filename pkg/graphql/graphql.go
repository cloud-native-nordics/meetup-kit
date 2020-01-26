package graphql

import (
	goflag "flag"
	"fmt"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/handler"
	"github.com/cloud-native-nordics/meetup-kit/pkg/graphql/generated"
	"github.com/cloud-native-nordics/meetup-kit/pkg/graphql/handlers"
	"github.com/cloud-native-nordics/meetup-kit/pkg/graphql/repositories"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/golang/glog"
	"github.com/rs/cors"
	flag "github.com/spf13/pflag"
)

type Options struct {
	// What port to serve GraphQL on
	Port uint64
	// ConfigPath describes the location of the config.json file, can also be an URL
	ConfigPath string
	// SlackToken is the Slack token to produce invites
	SlackToken string
	// SlackURL is the URL to the Slack community
	SlackURL string
	// SlackName is the name of the Slack community
	SlackName string
}

func Serve(opts *Options) error {
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

	db := handlers.NewStatsManager(opts.ConfigPath)

	statsRepo := repositories.NewStatsRepository(db)
	slackRepo := repositories.NewSlackRepository(opts.SlackToken, opts.SlackURL, opts.SlackName)

	resolver := handlers.NewResolver(statsRepo, slackRepo)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(5 * time.Second))

	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: resolver})))

	glog.V(5).Infof("Connect to http://localhost:%s/ for GraphQL playground", opts.Port)
	glog.Fatalf("Fatal: %s", http.ListenAndServe(fmt.Sprintf(":%d", opts.Port), router))
	return nil
}
