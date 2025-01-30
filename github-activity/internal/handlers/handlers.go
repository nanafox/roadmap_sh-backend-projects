package handlers

import (
	"errors"
	"fmt"

	"github.com/nanafox/github-activity/internal/utils"
)

// PushEventHandler handles the PushEvent type for a specific event.
func PushEventHandler(event *Event) (err error) {
	numOfPushes, ok := event.Payload["distinct_size"].(float64)
	if !ok {
		return errors.New("push_even_handler: Failed to retrieve number of pushes")
	}

	repoName, err := utils.GetRepoName(event.Repo)
	if err != nil {
		return err
	}

	commits := "commits"
	if numOfPushes <= 1 {
		commits = "commit"
	}

	fmt.Printf("- Pushed %v %s to %s\n", numOfPushes, commits, repoName)
	return
}

// WatchEventHandler handles the WatchEvent for a specific event.
func WatchEventHandler(event *Event) (err error) {
	repoName, err := utils.GetRepoName(event.Repo)
	if err != nil {
		return err
	}

	action, err := utils.GetPayloadAction(event.Payload)
	if err != nil {
		return err
	}

	fmt.Printf("- %s watching %s\n", action, repoName)
	return
}

// IssuesEventHandler handles issue based events for the user.
func IssuesEventHandler(event *Event) (err error) {
	repoName, err := utils.GetRepoName(event.Repo)
	if err != nil {
		return err
	}

	action, err := utils.GetPayloadAction(event.Payload)

	fmt.Printf("- %s an issue in %s\n", action, repoName)
	return
}

// CreateEventHandler handles all aspects of user creation actions on the
// platform. By all aspects, we mean the publicly available actions.
func CreateEventHandler(event *Event) (err error) {
	repoName, err := utils.GetRepoName(event.Repo)
	if err != nil {
		return err
	}

	refType, ok := event.Payload["ref_type"].(string)
	if !ok {
		return errors.New("create_event_handler: failed to extract ref_type")
	}

	switch refType {
	case "repository":
		fmt.Printf("- Created the %v %v\n", refType, repoName)
	case "branch", "tag":
		ref, _ := event.Payload["ref"].(string)
		fmt.Printf("- Created the %v %v in %v\n", ref, refType, repoName)
	default:
		fmt.Printf("- A %v was created\n", refType)
	}

	return nil
}

// PullRequestEventHandler handles the PullRequestEvent type.
func PullRequestEventHandler(event *Event) (err error) {
	action, err := utils.GetPayloadAction(event.Payload)
	if err != nil {
		return err
	}

	repo, err := utils.GetRepoName(event.Repo)
	if err != nil {
		return err
	}

	fmt.Printf("- %s a pull request in %s\n", action, repo)
	return nil
}

// IssueCommentEventHandler handles the IssueCommentEvent type.
func IssueCommentEventHandler(event *Event) (err error) {
	action, err := utils.GetPayloadAction(event.Payload)
	if err != nil {
		return err
	}

	issueUrl, ok := event.
		Payload["issue"].(map[string]any)["html_url"].(string)
	if !ok {
		return errors.New("failed to retrieve the HTML URL")
	}

	fmt.Printf("- %s a comment in issue %s\n", action, issueUrl)
	return
}

// DeleteEventHandler handles the DeleteEvent type.
func DeleteEventHandler(event *Event) (err error) {
	repo, err := utils.GetRepoName(event.Repo)
	if err != nil {
		return err
	}

	refType, ok := event.Payload["ref_type"].(string)
	if !ok {
		err = errors.New("delete_event_handler: failed to extract ref_type")
	}

	ref, ok := event.Payload["ref"]
	if !ok {
		err = errors.New("delete_event_handler: failed to extract ref")
	}

	fmt.Printf("- Deleted %s '%s' from %s\n", refType, ref, repo)
	return
}

// PullRequestReviewEventHandler handles the PullRequestReviewEvent type.
func PullRequestReviewEventHandler(event *Event) (err error) {
	action, err := utils.GetPayloadAction(event.Payload)
	if err != nil {
		return err
	}

	prUrl, ok := event.
		Payload["review"].(map[string]any)["html_url"].(string)
	if !ok {
		err = errors.New("failed to retrieve the URL for the pull request")
	}

	fmt.Printf("- %s a review comment to pull request %s\n", action, prUrl)
	return
}
