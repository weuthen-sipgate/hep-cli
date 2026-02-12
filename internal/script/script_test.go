package script

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"hepic-cli/internal/api"
)

func TestList(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/script" {
			t.Errorf("expected path /script, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"count": 2,
			"data": []map[string]interface{}{
				{"uuid": "script-1", "data": "print('hello')"},
				{"uuid": "script-2", "data": "print('world')"},
			},
		})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := List(context.Background(), client)
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

func TestCreate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/script" {
			t.Errorf("expected path /script, got %s", r.URL.Path)
		}

		body, _ := io.ReadAll(r.Body)
		var req map[string]interface{}
		json.Unmarshal(body, &req)
		if req["data"] != "print('test')" {
			t.Errorf("expected data 'print(\"test\")', got %v", req["data"])
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"uuid":    "new-script-uuid",
			"message": "successfully created",
		})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	data := map[string]interface{}{
		"data":      "print('test')",
		"hep_alias": "default",
		"hepid":     1,
		"profile":   "default",
		"type":      "lua",
		"status":    true,
	}

	result, err := Create(context.Background(), client, data)
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
	if parsed["uuid"] != "new-script-uuid" {
		t.Errorf("expected uuid 'new-script-uuid', got %v", parsed["uuid"])
	}
}

func TestDelete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/script/test-uuid-123" {
			t.Errorf("expected path /script/test-uuid-123, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "successfully deleted"})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := Delete(context.Background(), client, "test-uuid-123")
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
	if parsed["message"] != "successfully deleted" {
		t.Errorf("expected message 'successfully deleted', got %s", parsed["message"])
	}
}

func TestUpdate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/script/test-uuid-456" {
			t.Errorf("expected path /script/test-uuid-456, got %s", r.URL.Path)
		}

		body, _ := io.ReadAll(r.Body)
		var req map[string]interface{}
		json.Unmarshal(body, &req)
		if req["data"] != "print('updated')" {
			t.Errorf("expected data 'print(\"updated\")', got %v", req["data"])
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "successfully updated"})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	data := map[string]interface{}{
		"data": "print('updated')",
	}

	result, err := Update(context.Background(), client, "test-uuid-456", data)
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
	if parsed["message"] != "successfully updated" {
		t.Errorf("expected message 'successfully updated', got %s", parsed["message"])
	}
}

func TestList_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal error"})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	_, err := List(context.Background(), client)
	if err == nil {
		t.Fatal("expected error for 500 response")
	}
}
