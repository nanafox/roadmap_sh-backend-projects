package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nanafox/simple-http-client/pkg/client"

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

	apiClient := client.ApiClient{Timeout: 5 * time.Second}

	// handles the requests and caching
	utils.RequestHelper(&apiClient)
	if apiClient.Error != nil {
		log.Fatal(apiClient.Error)
	}

	var events []handlers.Event

	err := apiClient.ResponseToStruct(&events)
	if err != nil {
		log.Fatal(err)
	}

	for _, event := range events {
		err := cmd.ParseEvent(&event)
		if err != nil {
			fmt.Println(err)
		}
	}
}
