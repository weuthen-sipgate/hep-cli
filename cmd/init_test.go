package cmd

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidateConnection_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/version/api/info" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Header.Get("Auth-Token") != "test-token" {
			t.Errorf("expected Auth-Token header, got %s", r.Header.Get("Auth-Token"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data":{"version":"1.0"}}`))
	}))
	defer server.Close()

	if err := validateConnection(server.URL, "test-token"); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestValidateConnection_Unauthorized(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	err := validateConnection(server.URL, "bad-token")
	if err == nil {
		t.Error("expected error for unauthorized request")
	}
}

func TestValidateConnection_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	err := validateConnection(server.URL, "token")
	if err == nil {
		t.Error("expected error for server error")
	}
}

func TestValidateConnection_Unreachable(t *testing.T) {
	err := validateConnection("http://localhost:1", "token")
	if err == nil {
		t.Error("expected error for unreachable host")
	}
}
