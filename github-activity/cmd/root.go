package cmd

import (
	"fmt"

	"github.com/nanafox/github-activity/internal/handlers"
)

// ParseEvent parses the event type and hands control to a specialized function
// for that event data processing.
func ParseEvent(event *handlers.Event) (err error) {
	actionHandlers := map[string]func(event *handlers.Event) (err error){
		"PushEvent":              handlers.PushEventHandler,
		"WatchEvent":             handlers.WatchEventHandler,
		"IssuesEvent":            handlers.IssuesEventHandler,
		"CreateEvent":            handlers.CreateEventHandler,
		"PullRequestEvent":       handlers.PullRequestEventHandler,
		"IssueCommentEvent":      handlers.IssueCommentEventHandler,
		"DeleteEvent":            handlers.DeleteEventHandler,
		"PullRequestReviewEvent": handlers.PullRequestReviewEventHandler,
	}

	if handler, exists := actionHandlers[event.Type]; exists {
		return handler(event)
	}

	return fmt.Errorf("parse_event: %s is not handled at the moment", event.Type)
}
