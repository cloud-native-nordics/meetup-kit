package generator

import "fmt"

// Options for the generator
// TODO: Make this better :)
type Options struct {
	// SpeakersFile points to the speakers.yaml file
	SpeakersFile string
	// CompaniesFile points to the companies.yaml file
	CompaniesFile string
	// RootDir points to the directory that has all meetup groups as subfolders, each with a meetup.yaml file
	RootDir string
	// DryRun controls whether to actually apply the changes or not
	DryRun bool
	// Validate controls whether to validate the current state of the repo content with the spec
	Validate bool
}

func Generate(opts *Options) error {
	fmt.Println("Generating... no-op at the moment!")
	return nil
}
