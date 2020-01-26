package cmd

import (
	"github.com/cloud-native-nordics/meetup-kit/pkg/generator"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// NewGenerateCommand returns the "generate" command
func NewGenerateCommand() *cobra.Command {
	opts := &generator.Options{}
	cmd := &cobra.Command{
		Use:     "generate",
		Aliases: []string{"gen"},
		Short:   "Generate a set of README files, etc. based on the YAML",
		Run:     RunGen(opts),
	}

	addGenFlags(cmd.PersistentFlags(), opts)
	return cmd
}

func addGenFlags(fs *pflag.FlagSet, opts *generator.Options) {
	fs.StringVar(&opts.SpeakersFile, "speakers-file", "speakers.yaml", "Point to the speakers.yaml file")
	fs.StringVar(&opts.CompaniesFile, "companies-file", "companies.yaml", "Point to the companies.yaml file")
	fs.StringVar(&opts.RootDir, "meetups-dir", ".", "Point to the directory that has all meetup groups as subfolders, each with a meetup.yaml file")
	fs.BoolVar(&opts.DryRun, "dry-run", false, "Whether to actually apply the changes or not")
	fs.BoolVar(&opts.Validate, "validate", false, "Whether to validate the current state of the repo content with the spec")
}

func RunGen(opts *generator.Options) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if err := generator.Generate(opts); err != nil {
			log.Fatal(err)
		}
	}
}
