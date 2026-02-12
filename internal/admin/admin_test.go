package admin

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"hepic-cli/internal/api"
)

func TestProfiles(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/admin/profiles" {
			t.Errorf("expected path /admin/profiles, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"count": 2,
			"data":  []string{"admin", "user"},
		})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := Profiles(context.Background(), client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}
	if parsed["count"] != float64(2) {
		t.Errorf("expected count 2, got %v", parsed["count"])
	}
}

func TestConfigDBResync(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/configdb/tables/resync" {
			t.Errorf("expected path /configdb/tables/resync, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok", "message": "resync completed"})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := ConfigDBResync(context.Background(), client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}

	var parsed map[string]string
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}
	if parsed["status"] != "ok" {
		t.Errorf("expected status ok, got %s", parsed["status"])
	}
}

func TestAPIVersion(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/version/api/info" {
			t.Errorf("expected path /version/api/info, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"version": "1.2.1", "build": "abc123"})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := APIVersion(context.Background(), client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}

	var parsed map[string]string
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}
	if parsed["version"] != "1.2.1" {
		t.Errorf("expected version 1.2.1, got %s", parsed["version"])
	}
}

func TestClickhouseQuery(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/clickhouse/query/raw" {
			t.Errorf("expected path /clickhouse/query/raw, got %s", r.URL.Path)
		}

		body, _ := io.ReadAll(r.Body)
		var req map[string]string
		json.Unmarshal(body, &req)
		if req["query"] != "SELECT count() FROM hep" {
			t.Errorf("expected query 'SELECT count() FROM hep', got %s", req["query"])
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data": []map[string]interface{}{
				{"count()": 42},
			},
		})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := ClickhouseQuery(context.Background(), client, "SELECT count() FROM hep")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}
	data, ok := parsed["data"].([]interface{})
	if !ok || len(data) == 0 {
		t.Fatal("expected data array with at least one element")
	}
}

func TestConfigDBTables(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/configdb/tables/list" {
			t.Errorf("expected path /configdb/tables/list, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"count": 3,
			"data":  []string{"users", "settings", "profiles"},
		})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := ConfigDBTables(context.Background(), client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestTroubleshootingLog(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/troubleshooting/log/system/status" {
			t.Errorf("expected path /troubleshooting/log/system/status, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"log": "all systems operational"})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := TroubleshootingLog(context.Background(), client, "system", "status")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestShareTransaction(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/share/call/transaction/test-uuid-123" {
			t.Errorf("expected path /share/call/transaction/test-uuid-123, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"url": "https://example.com/share/abc"})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := ShareTransaction(context.Background(), client, "test-uuid-123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}

	var parsed map[string]string
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}
	if parsed["url"] != "https://example.com/share/abc" {
		t.Errorf("expected share URL, got %s", parsed["url"])
	}
}

func TestShareReport(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/share/call/report/dtmf/test-uuid" {
			t.Errorf("expected path /share/call/report/dtmf/test-uuid, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "shared"})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := ShareReport(context.Background(), client, "dtmf", "test-uuid")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestUIVersion(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/version/ui/info" {
			t.Errorf("expected path /version/ui/info, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"version": "3.0.0"})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := UIVersion(context.Background(), client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}

	var parsed map[string]string
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}
	if parsed["version"] != "3.0.0" {
		t.Errorf("expected version 3.0.0, got %s", parsed["version"])
	}
}
