package graphql

import "fmt"

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
	fmt.Println("Serving graphql... no-op at the moment!")
	return nil
}
