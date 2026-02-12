package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.Header.Get("Auth-Token") != "test-token" {
			t.Errorf("missing Auth-Token header")
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("missing Content-Type header")
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}))
	defer server.Close()

	client := NewClientWith(server.URL, "test-token")
	var result map[string]string
	err := client.Get(context.Background(), "", &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result["status"] != "ok" {
		t.Errorf("expected status ok, got %s", result["status"])
	}
}

func TestPost_WithBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		var req map[string]string
		json.Unmarshal(body, &req)
		if req["query"] != "test" {
			t.Errorf("expected query=test, got %s", req["query"])
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]int{"count": 42})
	}))
	defer server.Close()

	client := NewClientWith(server.URL, "token")
	var result map[string]int
	err := client.Post(context.Background(), "", map[string]string{"query": "test"}, &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result["count"] != 42 {
		t.Errorf("expected count 42, got %d", result["count"])
	}
}

func TestPut_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"updated": "true"})
	}))
	defer server.Close()

	client := NewClientWith(server.URL, "token")
	var result map[string]string
	err := client.Put(context.Background(), "/resource/123", map[string]string{"name": "new"}, &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDelete_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"deleted": "true"})
	}))
	defer server.Close()

	client := NewClientWith(server.URL, "token")
	var result map[string]string
	err := client.Delete(context.Background(), "/resource/123", &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGet_Unauthorized(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized", "message": "invalid token"})
	}))
	defer server.Close()

	client := NewClientWith(server.URL, "bad-token")
	var result map[string]string
	err := client.Get(context.Background(), "", &result)
	if err == nil {
		t.Fatal("expected error for 401")
	}
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.StatusCode != 401 {
		t.Errorf("expected status 401, got %d", apiErr.StatusCode)
	}
}

func TestGet_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "not found"})
	}))
	defer server.Close()

	client := NewClientWith(server.URL, "token")
	err := client.Get(context.Background(), "/missing", nil)
	if err == nil {
		t.Fatal("expected error for 404")
	}
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.StatusCode != 404 {
		t.Errorf("expected status 404, got %d", apiErr.StatusCode)
	}
}

func TestGet_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("not json"))
	}))
	defer server.Close()

	client := NewClientWith(server.URL, "token")
	err := client.Get(context.Background(), "", nil)
	if err == nil {
		t.Fatal("expected error for 500")
	}
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.StatusCode != 500 {
		t.Errorf("expected status 500, got %d", apiErr.StatusCode)
	}
}

func TestGet_NetworkError(t *testing.T) {
	client := NewClientWith("http://localhost:1", "token")
	err := client.Get(context.Background(), "", nil)
	if err == nil {
		t.Fatal("expected error for network failure")
	}
}

func TestGetRaw_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("binary-data"))
	}))
	defer server.Close()

	client := NewClientWith(server.URL, "token")
	body, err := client.GetRaw(context.Background(), "/download")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer body.Close()
	data, _ := io.ReadAll(body)
	if string(data) != "binary-data" {
		t.Errorf("expected binary-data, got %s", string(data))
	}
}

func TestVerboseLogging(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{})
	}))
	defer server.Close()

	client := NewClientWith(server.URL, "token")
	client.Verbose = true
	// Just ensure it doesn't panic
	client.Get(context.Background(), "/test", nil)
}
