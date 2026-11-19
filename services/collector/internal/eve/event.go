package eve

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

var ErrMissingEventType = errors.New("missing event_type")

// Event contains only the stable fields currently needed by the collector
// foundation. The original event is not persisted or printed by this package.
type Event struct {
	Timestamp string `json:"timestamp"`
	EventType string `json:"event_type"`
}

func ParseLine(line []byte) (Event, error) {
	var event Event
	if err := json.Unmarshal(line, &event); err != nil {
		return Event{}, fmt.Errorf("decode EVE JSON: %w", err)
	}
	event.EventType = strings.TrimSpace(event.EventType)
	if event.EventType == "" {
		return Event{}, ErrMissingEventType
	}
	if event.Timestamp != "" {
		if _, err := time.Parse(time.RFC3339Nano, event.Timestamp); err != nil {
			return Event{}, fmt.Errorf("invalid timestamp: %w", err)
		}
	}
	return event, nil
}
