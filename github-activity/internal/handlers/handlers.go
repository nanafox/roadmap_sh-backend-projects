package handlers

import (
	"errors"
	"fmt"

	"github.com/nanafox/github-activity/internal/utils"
)

// PushEventHandler handles the PushEvent type for a specific event.
func PushEventHandler(event *Event) (err error) {
	numOfPushes, found := event.Payload["distinct_size"]
	if !found {
		return errors.New("push_even_handler: Failed to retrieve number of pushes")
	}

	repoName, err := utils.GetRepoName(event.Repo)
	if err != nil {
		return err
	}

	commits := "commits"
	if numOfPushes.(float64) <= 1 {
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

	fmt.Printf("- %v watching %v\n", action, repoName)
	return
}

// IssuesEventHandler handles issue based events for the user.
func IssuesEventHandler(event *Event) (err error) {
	repoName, err := utils.GetRepoName(event.Repo)
	if err != nil {
		return err
	}

	action, err := utils.GetPayloadAction(event.Payload)

	fmt.Printf("- %v an issue in %v\n", action, repoName)
	return
}

// CreateEventHandler handles all aspects of user creation actions on the
// platform. By all aspects, we mean the publicly available actions.
func CreateEventHandler(event *Event) (err error) {
	repoName, err := utils.GetRepoName(event.Repo)
	if err != nil {
		return err
	}

	refType := event.Payload["ref_type"]
	switch refType {
	case "repository":
		fmt.Printf("- Created the %v %v\n", refType, repoName)
	case "branch", "tag":
		ref := event.Payload["ref"]
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

	fmt.Printf("- %v a pull request in %v\n", action, repo)
	return nil
}

// IssueCommentEventHandler handles the IssueCommentEvent type.
func IssueCommentEventHandler(event *Event) (err error) {
	action, err := utils.GetPayloadAction(event.Payload)
	if err != nil {
		return err
	}

	issueUrl := event.Payload["issue"].(map[string]any)["html_url"]

	fmt.Printf("- %v a comment in issue %v\n", action, issueUrl)
	return
}

// DeleteEventHandler handles the DeleteEvent type.
func DeleteEventHandler(event *Event) (err error) {
	repo, err := utils.GetRepoName(event.Repo)
	if err != nil {
		return err
	}

	refType := event.Payload["ref_type"]
	ref := event.Payload["ref"]

	fmt.Printf("- Deleted %v '%v' from %v\n", refType, ref, repo)
	return nil
}

// PullRequestReviewEventHandler handles the PullRequestReviewEvent type.
func PullRequestReviewEventHandler(event *Event) (err error) {
	action, err := utils.GetPayloadAction(event.Payload)
	if err != nil {
		return err
	}

	prUrl := event.Payload["review"].(map[string]any)["html_url"]

	fmt.Printf("- %v a review comment to pull request %v\n", action, prUrl)
	return nil
}
