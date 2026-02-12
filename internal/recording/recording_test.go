package recording

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"hepic-cli/internal/api"
	"hepic-cli/internal/models"
)

func TestSearchData(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/call/recording/data" {
			t.Errorf("expected path /call/recording/data, got %s", r.URL.Path)
		}

		body, _ := io.ReadAll(r.Body)
		var params SearchParams
		if err := json.Unmarshal(body, &params); err != nil {
			t.Fatalf("failed to unmarshal request body: %v", err)
		}
		if params.Timestamp == nil {
			t.Error("expected timestamp in request body")
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data": []map[string]string{
				{"uuid": "rec-001", "date": "2025-01-15"},
				{"uuid": "rec-002", "date": "2025-01-16"},
			},
		})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	params, err := NewSearchParams("2025-01-01", "2025-01-31")
	if err != nil {
		t.Fatalf("failed to create search params: %v", err)
	}

	result, err := SearchData(context.Background(), client, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}

	data, ok := parsed["data"].([]interface{})
	if !ok {
		t.Fatal("expected data to be an array")
	}
	if len(data) != 2 {
		t.Errorf("expected 2 recordings, got %d", len(data))
	}
}

func TestInfo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/call/recording/info/test-uuid-123" {
			t.Errorf("expected path /call/recording/info/test-uuid-123, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"uuid":     "test-uuid-123",
			"duration": 120,
			"codec":    "g711",
		})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := Info(context.Background(), client, "test-uuid-123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}

	if parsed["uuid"] != "test-uuid-123" {
		t.Errorf("expected uuid test-uuid-123, got %v", parsed["uuid"])
	}
}

func TestListInterceptions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/interceptions" {
			t.Errorf("expected path /interceptions, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data": []map[string]string{
				{"uuid": "int-001", "search_caller": "+4912345"},
				{"uuid": "int-002", "search_caller": "+4967890"},
			},
		})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := ListInterceptions(context.Background(), client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}

	data, ok := parsed["data"].([]interface{})
	if !ok {
		t.Fatal("expected data to be an array")
	}
	if len(data) != 2 {
		t.Errorf("expected 2 interceptions, got %d", len(data))
	}
}

func TestCreateInterception(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/interceptions" {
			t.Errorf("expected path /interceptions, got %s", r.URL.Path)
		}

		body, _ := io.ReadAll(r.Body)
		var data models.InterceptionsStruct
		if err := json.Unmarshal(body, &data); err != nil {
			t.Fatalf("failed to unmarshal request body: %v", err)
		}
		if data.SearchCaller != "+4912345" {
			t.Errorf("expected caller +4912345, got %s", data.SearchCaller)
		}
		if data.SearchCallee != "+4967890" {
			t.Errorf("expected callee +4967890, got %s", data.SearchCallee)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"uuid":          "new-int-001",
			"search_caller": data.SearchCaller,
			"search_callee": data.SearchCallee,
			"status":        "created",
		})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	input := models.InterceptionsStruct{
		SearchCaller: "+4912345",
		SearchCallee: "+4967890",
		Description:  "Test interception",
	}

	result, err := CreateInterception(context.Background(), client, input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}

	if parsed["uuid"] != "new-int-001" {
		t.Errorf("expected uuid new-int-001, got %v", parsed["uuid"])
	}
}

func TestDeleteInterception(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/interceptions/del-uuid-001" {
			t.Errorf("expected path /interceptions/del-uuid-001, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "ok",
			"message": "interception deleted",
		})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := DeleteInterception(context.Background(), client, "del-uuid-001")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}

	if parsed["status"] != "ok" {
		t.Errorf("expected status ok, got %v", parsed["status"])
	}
}

func TestNewSearchParams(t *testing.T) {
	// Test with valid date strings
	params, err := NewSearchParams("2025-01-01", "2025-01-31")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if params.Timestamp == nil {
		t.Fatal("expected timestamp to be set")
	}
	if _, ok := params.Timestamp["from"]; !ok {
		t.Error("expected 'from' in timestamp")
	}
	if _, ok := params.Timestamp["to"]; !ok {
		t.Error("expected 'to' in timestamp")
	}

	// Test with empty strings
	params, err = NewSearchParams("", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if params.Timestamp != nil {
		t.Error("expected timestamp to be nil when no dates provided")
	}

	// Test with invalid date
	_, err = NewSearchParams("not-a-date", "")
	if err == nil {
		t.Error("expected error for invalid date")
	}
}

func TestSearchData_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "server error"})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	params := SearchParams{Param: map[string]interface{}{}}
	_, err := SearchData(context.Background(), client, params)
	if err == nil {
		t.Fatal("expected error for server error response")
	}
}

func TestInfo_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "not found"})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	_, err := Info(context.Background(), client, "nonexistent")
	if err == nil {
		t.Fatal("expected error for not found response")
	}
}

func TestUpdateInterception(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/interceptions/upd-uuid-001" {
			t.Errorf("expected path /interceptions/upd-uuid-001, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"uuid":   "upd-uuid-001",
			"status": "updated",
		})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	input := models.InterceptionsStruct{
		SearchCaller: "+4999999",
	}

	result, err := UpdateInterception(context.Background(), client, "upd-uuid-001", input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}

	if parsed["uuid"] != "upd-uuid-001" {
		t.Errorf("expected uuid upd-uuid-001, got %v", parsed["uuid"])
	}
}
