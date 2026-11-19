package eve

import (
	"errors"
	"testing"
)

func TestParseLineValid(t *testing.T) {
	got, err := ParseLine([]byte(`{"timestamp":"2026-11-19T10:00:00.123Z","event_type":"alert"}`))
	if err != nil {
		t.Fatalf("ParseLine returned error: %v", err)
	}
	if got.EventType != "alert" {
		t.Fatalf("expected alert event, got %q", got.EventType)
	}
}

func TestParseLineRejectsMalformedJSON(t *testing.T) {
	if _, err := ParseLine([]byte(`{"event_type":`)); err == nil {
		t.Fatal("expected malformed JSON to be rejected")
	}
}

func TestParseLineRequiresEventType(t *testing.T) {
	_, err := ParseLine([]byte(`{"timestamp":"2026-11-19T10:00:00Z"}`))
	if !errors.Is(err, ErrMissingEventType) {
		t.Fatalf("expected ErrMissingEventType, got %v", err)
	}
}

func TestParseLineRejectsInvalidTimestamp(t *testing.T) {
	_, err := ParseLine([]byte(`{"timestamp":"not-a-time","event_type":"flow"}`))
	if err == nil {
		t.Fatal("expected invalid timestamp to be rejected")
	}
}
