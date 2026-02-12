package agent

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"hepic-cli/internal/api"
)

func TestList(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/agent/subscribe" {
			t.Errorf("expected path /agent/subscribe, got %s", r.URL.Path)
		}
		if r.Header.Get("Auth-Token") != "test-token" {
			t.Errorf("missing auth token")
		}
		w.Write([]byte(`{"data":[{"uuid":"abc","host":"10.0.0.1","port":9060}]}`))
	}))
	defer srv.Close()

	client := api.NewClientWith(srv.URL, "test-token")
	result, err := List(context.Background(), client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected result")
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}
	data, ok := parsed["data"].([]interface{})
	if !ok {
		t.Fatal("expected data array")
	}
	if len(data) != 1 {
		t.Fatalf("expected 1 agent, got %d", len(data))
	}
}

func TestGet(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/agent/subscribe/abc-123" {
			t.Errorf("expected path /agent/subscribe/abc-123, got %s", r.URL.Path)
		}
		if r.Header.Get("Auth-Token") != "test-token" {
			t.Errorf("missing auth token")
		}
		w.Write([]byte(`{"data":{"uuid":"abc-123","host":"10.0.0.1","port":9060,"type":"home"}}`))
	}))
	defer srv.Close()

	client := api.NewClientWith(srv.URL, "test-token")
	result, err := Get(context.Background(), client, "abc-123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected result")
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}
	data, ok := parsed["data"].(map[string]interface{})
	if !ok {
		t.Fatal("expected data object")
	}
	if data["uuid"] != "abc-123" {
		t.Errorf("expected uuid abc-123, got %v", data["uuid"])
	}
}

func TestSearch(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/agent/search/guid-1/home" {
			t.Errorf("expected path /agent/search/guid-1/home, got %s", r.URL.Path)
		}
		if r.Header.Get("Auth-Token") != "test-token" {
			t.Errorf("missing auth token")
		}
		w.Write([]byte(`{"data":[{"uuid":"abc","host":"10.0.0.1","type":"home"}]}`))
	}))
	defer srv.Close()

	client := api.NewClientWith(srv.URL, "test-token")
	result, err := Search(context.Background(), client, "guid-1", "home")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected result")
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}
	data, ok := parsed["data"].([]interface{})
	if !ok {
		t.Fatal("expected data array")
	}
	if len(data) != 1 {
		t.Fatalf("expected 1 agent, got %d", len(data))
	}
}

func TestUpdate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/agent/subscribe/abc-123" {
			t.Errorf("expected path /agent/subscribe/abc-123, got %s", r.URL.Path)
		}
		w.Write([]byte(`{"data":"successfully updated","message":"ok"}`))
	}))
	defer srv.Close()

	client := api.NewClientWith(srv.URL, "test-token")
	data := map[string]interface{}{"host": "10.0.0.2", "port": 9061}
	result, err := Update(context.Background(), client, "abc-123", data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected result")
	}
}

func TestDelete(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/agent/subscribe/abc-123" {
			t.Errorf("expected path /agent/subscribe/abc-123, got %s", r.URL.Path)
		}
		w.Write([]byte(`{"data":"successfully deleted","message":"ok"}`))
	}))
	defer srv.Close()

	client := api.NewClientWith(srv.URL, "test-token")
	result, err := Delete(context.Background(), client, "abc-123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected result")
	}
}

func TestListByType(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/agent/type/home" {
			t.Errorf("expected path /agent/type/home, got %s", r.URL.Path)
		}
		w.Write([]byte(`{"data":[{"uuid":"abc","type":"home"}]}`))
	}))
	defer srv.Close()

	client := api.NewClientWith(srv.URL, "test-token")
	result, err := ListByType(context.Background(), client, "home")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected result")
	}
}

func TestAddSubscription(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/agentsub/protocol" {
			t.Errorf("expected path /agentsub/protocol, got %s", r.URL.Path)
		}
		w.Write([]byte(`{"data":"successfully created","message":"ok"}`))
	}))
	defer srv.Close()

	client := api.NewClientWith(srv.URL, "test-token")
	data := map[string]interface{}{"protocol": "sip", "port": 9060}
	result, err := AddSubscription(context.Background(), client, data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected result")
	}
}

func TestList_Error(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"internal server error","message":"something went wrong"}`))
	}))
	defer srv.Close()

	client := api.NewClientWith(srv.URL, "test-token")
	_, err := List(context.Background(), client)
	if err == nil {
		t.Fatal("expected error for 500 response")
	}
}
