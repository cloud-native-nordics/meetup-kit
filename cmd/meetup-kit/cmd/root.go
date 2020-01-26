package cmd

import (
	"io"
	"os"

	"github.com/cloud-native-nordics/meetup-kit/pkg/logs"
	logflag "github.com/cloud-native-nordics/meetup-kit/pkg/logs/flag"
	versioncmd "github.com/cloud-native-nordics/meetup-kit/pkg/version/cmd"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var logLevel = logrus.InfoLevel

// NewMeetupKitCommand returns the root command for ignite
func NewMeetupKitCommand(in io.Reader, out, err io.Writer) *cobra.Command {

	root := &cobra.Command{
		Use:   "meetup-kit",
		Short: "meetup-kit: Manage Meetups by Pull Request -- MeetOps!",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Set the desired logging level, now that the flags are parsed
			logs.Logger.SetLevel(logLevel)
		},
	}

	addGlobalFlags(root.PersistentFlags())

	root.AddCommand(NewGenerateCommand())
	root.AddCommand(NewServeCommand())
	root.AddCommand(versioncmd.NewCmdVersion(os.Stdout))
	return root
}

func addGlobalFlags(fs *pflag.FlagSet) {
	logflag.LogLevelFlagVar(fs, &logLevel)
}
