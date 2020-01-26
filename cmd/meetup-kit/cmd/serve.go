package cmd

import (
	"github.com/cloud-native-nordics/meetup-kit/pkg/graphql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// NewServeCommand returns the "serve" command
func NewServeCommand() *cobra.Command {
	opts := &graphql.Options{}
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve GraphQL requests and UI",
		Run:   RunServe(opts),
	}

	addServeFlags(cmd.PersistentFlags(), opts)
	return cmd
}

func addServeFlags(fs *pflag.FlagSet, opts *graphql.Options) {
	fs.Uint64Var(&opts.Port, "port", 8080, "Application port to use")
	fs.StringVar(&opts.ConfigPath, "stats-url", "https://raw.githubusercontent.com/cloud-native-nordics/meetups/master/config.json", "Location of the stats file")
	fs.StringVar(&opts.SlackToken, "slack-token", "", "Slack token to produce invites")
	fs.StringVar(&opts.SlackURL, "slack-url", "https://cloud-native-nordics.slack.com", "URL to the slack community")
	fs.StringVar(&opts.SlackName, "slack-community", "Cloud Native Nordics", "Name of the slack community")
}

func RunServe(opts *graphql.Options) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if err := graphql.Serve(opts); err != nil {
			log.Fatal(err)
		}
	}
}
