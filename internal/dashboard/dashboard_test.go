package dashboard

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"hepic-cli/internal/api"
)

func TestList_Success(t *testing.T) {
	expected := map[string]interface{}{
		"status": "ok",
		"data": []interface{}{
			map[string]interface{}{"id": "dash1", "name": "Dashboard 1"},
			map[string]interface{}{"id": "dash2", "name": "Dashboard 2"},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/dashboard/info" {
			t.Errorf("expected path /dashboard/info, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expected)
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := List(context.Background(), client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}
	if parsed["status"] != "ok" {
		t.Errorf("expected status ok, got %v", parsed["status"])
	}
}

func TestStore_SendsPUT(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/dashboard/store/dash1" {
			t.Errorf("expected path /dashboard/store/dash1, got %s", r.URL.Path)
		}
		body, _ := io.ReadAll(r.Body)
		var req map[string]interface{}
		json.Unmarshal(body, &req)
		if req["name"] != "Updated Dashboard" {
			t.Errorf("expected name 'Updated Dashboard', got %v", req["name"])
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	data := map[string]interface{}{"name": "Updated Dashboard", "type": 1}
	result, err := Store(context.Background(), client, "dash1", data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed map[string]string
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}
	if parsed["status"] != "ok" {
		t.Errorf("expected status ok, got %s", parsed["status"])
	}
}

func TestDelete_SendsDELETE(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/dashboard/store/dash1" {
			t.Errorf("expected path /dashboard/store/dash1, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := Delete(context.Background(), client, "dash1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed map[string]string
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}
	if parsed["status"] != "ok" {
		t.Errorf("expected status ok, got %s", parsed["status"])
	}
}
