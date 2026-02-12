package config_resources

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"hepic-cli/internal/api"
)

func TestListAliases(t *testing.T) {
	expected := map[string]interface{}{
		"count": float64(2),
		"data": []interface{}{
			map[string]interface{}{"uuid": "aaa", "alias": "proxy1", "ip": "10.0.0.1"},
			map[string]interface{}{"uuid": "bbb", "alias": "proxy2", "ip": "10.0.0.2"},
		},
	}
	respBody, _ := json.Marshal(expected)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/ipalias" {
			t.Errorf("expected /ipalias, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respBody)
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := ListAliases(context.Background(), client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var got map[string]interface{}
	if err := json.Unmarshal(result, &got); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}
	if got["count"] != float64(2) {
		t.Errorf("expected count 2, got %v", got["count"])
	}
	data, ok := got["data"].([]interface{})
	if !ok || len(data) != 2 {
		t.Errorf("expected 2 aliases in data, got %v", got["data"])
	}
}

func TestCreateAlias(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/ipalias" {
			t.Errorf("expected /ipalias, got %s", r.URL.Path)
		}

		body, _ := io.ReadAll(r.Body)
		var req map[string]interface{}
		json.Unmarshal(body, &req)

		if req["ip"] != "10.0.0.1" {
			t.Errorf("expected ip 10.0.0.1, got %v", req["ip"])
		}
		if req["alias"] != "proxy1" {
			t.Errorf("expected alias proxy1, got %v", req["alias"])
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"data": "af1234", "message": "successfully created"})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	data := map[string]interface{}{
		"ip":    "10.0.0.1",
		"alias": "proxy1",
		"port":  5060,
		"mask":  32,
	}
	result, err := CreateAlias(context.Background(), client, data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var got map[string]string
	if err := json.Unmarshal(result, &got); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}
	if got["message"] != "successfully created" {
		t.Errorf("expected 'successfully created', got %s", got["message"])
	}
}

func TestDeleteAlias(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/ipalias/test-uuid-123" {
			t.Errorf("expected /ipalias/test-uuid-123, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"data": "test-uuid-123", "message": "successfully deleted"})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := DeleteAlias(context.Background(), client, "test-uuid-123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var got map[string]string
	if err := json.Unmarshal(result, &got); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}
	if got["message"] != "successfully deleted" {
		t.Errorf("expected 'successfully deleted', got %s", got["message"])
	}
}

func TestDeleteAllAliases(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/ipalias/all" {
			t.Errorf("expected /ipalias/all, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"data": "all", "message": "successfully deleted all"})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := DeleteAllAliases(context.Background(), client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var got map[string]string
	if err := json.Unmarshal(result, &got); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}
	if got["message"] != "successfully deleted all" {
		t.Errorf("expected 'successfully deleted all', got %s", got["message"])
	}
}
