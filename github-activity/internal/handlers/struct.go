package handlers

// Event struct serves as the place to store the GitHub Event data for easy
// processing.
type Event struct {
	Actor   map[string]any
	Repo    map[string]any
	Payload map[string]any
	ID      string
	Type    string
}
