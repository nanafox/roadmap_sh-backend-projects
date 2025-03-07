package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nanafox/gofetch"

	"github.com/nanafox/github-activity/cmd"
	"github.com/nanafox/github-activity/internal/handlers"
	"github.com/nanafox/github-activity/internal/utils"
)

const usage = `Usage: github-activity <username> [number of events: Default is 10],

  github-activity is a simple CLI tool to fetch the latest activities a user
  is performing with their GitHub account. It returns the results in a simple
  listing format.

  For efficiency, responses are cached for later in the /tmp/github-activity-cache/
  directory. If you need to get fresh updates for the same user, feel free to
  delete the cached files for that user or all users if you prefer.

  The cache files are saved in this format, <username>_<page_size>. So as an
  example running 'github-activity nanafox 10' will return the first 10 events
  of the 'nanafox' user. The resultant cache file will be 'nanafox_10.json'
`

// main is the entry point for the GitHub User Activity CLI program.
func main() {
	if len(os.Args) < 2 {
		os.Stderr.WriteString(usage)
		os.Exit(1)
	}

	client := gofetch.New(gofetch.Config{Timeout: 5 * time.Second})

	// handles the requests and caching
	utils.RequestHelper(client)
	if client.Error != nil {
		log.Fatal(client.Error)
	}

	var events []handlers.Event

	if client.StatusCode == 404 {
		fmt.Println("github-activity: This user does not exist!")
		return
	}

	err := client.ResponseToStruct(&events)
	if err != nil {
		log.Fatal(err)
	}

	results := make(chan error, len(events))

	for _, event := range events {
		go func(event handlers.Event) {
			results <- cmd.ParseEvent(&event)
		}(event)
	}

	for range events {
		if err := <-results; err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
	close(results)
}
