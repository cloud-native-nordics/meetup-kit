package main

import (
	"log"
	"os"

	"github.com/cloud-native-nordics/meetup-kit/cmd/meetup-kit/cmd"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

// Run runs the main cobra command of this application
func Run() error {
	c := cmd.NewMeetupKitCommand(os.Stdin, os.Stdout, os.Stderr)
	return c.Execute()
}
