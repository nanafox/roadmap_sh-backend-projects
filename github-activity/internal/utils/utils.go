package utils

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/nanafox/gofetch"
)

const cacheDir string = "/tmp/github-activity-cache/"

// init creates the caching directory that will be used to save responses for
// quicker responses when the same request is made.
func init() {
	if _, err := os.Stat(cacheDir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(cacheDir, os.ModePerm)
		if err != nil {
			log.Println("init: failed to create caching directory")
		}

	}
}

// GetPayloadAction returns the action performed by event which is present in
// the `Payload` of the of response.
func GetPayloadAction(payload map[string]any) (action string, err error) {
	name, ok := payload["action"].(string)
	if !ok {
		err = errors.New("get_payload_action: failed to retrieve payload action")
	}

	action = capitalize(name)
	return
}

// capitalize capitalizes the first letter in a given word.
func capitalize(word string) (capitalizedWord string) {
	return strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
}

// getRepoName returns the name of of repo in the map provided. This is usually
// from the response body.
func GetRepoName(repoMap map[string]any) (repoName string, err error) {
	name, ok := repoMap["name"].(string)
	if !ok {
		err = errors.New("get_repo_name: failed to retrieve repo name")
	}

	repoName = name
	return
}

// RequestHelper handles the request to the GitHub API to retrieve the user's
// most recent events.
//
// Parameters:
//
//	client (gofetch.Client) The API client to use for the request. This
//	uses my personal API gofetch.
//
// Behavior:
//
//	On the first request, a request is made to the GitHub API server for the
//	results of the user's events. On consequent requests using the same
//	information as the last, the cached response is returned and no API call
//	is made to the GitHub to save bandwidth and make the operation faster.
//
//	It does not return anything, the `client` is updated with the results.
//	This is because a reference to the API client is used and the `Body` and
//	`Error` fields of the `client` allow updates.
//
// Cache File Location:
//
//	The cached files have the format <username>_<page_size> and stored in the
//	/tmp directory.
func RequestHelper(client *gofetch.Client) {
	username := os.Args[1]

	page_size := "10" // Default to 10 events at a time.
	if len(os.Args) == 3 {
		page_size = os.Args[2]
	}

	cacheFile := cacheDir + username + "_" + page_size + ".json"

	url := "https://api.github.com/users/" + username + "/events/public"
	queryParams := []gofetch.Query{
		{Key: "per_page", Value: page_size},
	}
	headers := gofetch.Header{
		Key: "Accept", Value: "application/vnd.github+json",
	}

	// Try to use the cached response if it exists
	if cacheExists(cacheFile) {
		// The cache was hit, set up the response body with this data
		cacheContent, err := readFromCache(cacheFile)
		if err != nil {
			log.Fatal(err)
			// remove the file so that the next request will start afresh
			go os.Remove(cacheFile)

			// issue the same request again since the file was damages and didn't
			// return the right information needed
			RequestHelper(client)
		}

		// set the fields of the api client to use the results from the cache
		client.Body = string(cacheContent)
		client.Error = err
	} else { // make the API request to retrieve the user's events.
		client.Get(url, queryParams, headers)
		// cache the response on success
		if client.Error == nil {
			go writeToCache(cacheFile, []byte(client.Body))
		}
	}
}

// writeToCache writes the received body to a file for later access. This helps
// to speed up the request-response cycles by returning the same data from a
// previous request.
func writeToCache(cacheFile string, body []byte) (err error) {
	err = os.WriteFile(cacheFile, []byte(body), 0644)
	if err != nil {
		return err
	}

	return nil
}

// readFromCache fetches the data saved from the last API request.
func readFromCache(cacheFile string) (cacheContent []byte, err error) {
	cacheContent, err = os.ReadFile(cacheFile)
	if err != nil {
		return nil, err
	}

	return cacheContent, nil
}

func cacheExists(cacheFile string) bool {
	_, err := os.Stat(cacheFile)
	errors.Is(err, os.ErrNotExist)

	return err == nil
}
