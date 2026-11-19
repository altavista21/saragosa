package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHealthEndpoint(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_ = json.NewEncoder(w).Encode(healthResponse{
			Status:    "ok",
			Service:   "saragosa-api",
			Mode:      "ids-passive",
			Timestamp: time.Now().UTC(),
			Scope:     "api-process-only",
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/api/v1/health", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var got healthResponse
	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if got.Status != "ok" || got.Service != "saragosa-api" || got.Mode != "ids-passive" {
		t.Fatalf("unexpected health response: %+v", got)
	}
	if got.Scope != "api-process-only" {
		t.Fatalf("health scope must not imply Suricata health, got %q", got.Scope)
	}
}
