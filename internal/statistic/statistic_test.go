package statistic

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"hepic-cli/internal/api"
)

func TestDBStats_Success(t *testing.T) {
	expected := map[string]interface{}{
		"database_name":    "hepic",
		"database_version": "10.5",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/statistic/_db" {
			t.Errorf("expected path /statistic/_db, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expected)
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := DBStats(context.Background(), client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}
	if parsed["database_name"] != "hepic" {
		t.Errorf("expected database_name 'hepic', got %v", parsed["database_name"])
	}
}

func TestData_PostWithBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/statistic/data" {
			t.Errorf("expected path /statistic/data, got %s", r.URL.Path)
		}
		body, _ := io.ReadAll(r.Body)
		var req map[string]interface{}
		json.Unmarshal(body, &req)
		if req["param"] == nil {
			t.Errorf("expected param in request body")
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": []interface{}{}, "count": 0})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	data := map[string]interface{}{
		"param":     map[string]interface{}{},
		"timestamp": map[string]interface{}{"from": 1700000000, "to": 1700003600},
	}
	result, err := Data(context.Background(), client, data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}
}

func TestMeasurements_PostWithDBID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/statistic/_measurements/mydb" {
			t.Errorf("expected path /statistic/_measurements/mydb, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": []interface{}{}})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := Measurements(context.Background(), client, "mydb", map[string]interface{}{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestMetrics_Post(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/statistic/_metrics" {
			t.Errorf("expected path /statistic/_metrics, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": []interface{}{}})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := Metrics(context.Background(), client, map[string]interface{}{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestRetentions_Post(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/statistic/_retentions" {
			t.Errorf("expected path /statistic/_retentions, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": []interface{}{}})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := Retentions(context.Background(), client, map[string]interface{}{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestLabels_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/prometheus/labels" {
			t.Errorf("expected path /prometheus/labels, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data": []interface{}{"instance", "job", "method"},
		})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := Labels(context.Background(), client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}
	data, ok := parsed["data"].([]interface{})
	if !ok {
		t.Fatal("expected data to be an array")
	}
	if len(data) != 3 {
		t.Errorf("expected 3 labels, got %d", len(data))
	}
}

func TestLabelDetail_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/prometheus/label/instance" {
			t.Errorf("expected path /prometheus/label/instance, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data": []interface{}{"localhost:9090"},
		})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := LabelDetail(context.Background(), client, "instance")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestQueryData_Post(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/prometheus/data" {
			t.Errorf("expected path /prometheus/data, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": []interface{}{}})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := QueryData(context.Background(), client, map[string]interface{}{
		"param": map[string]interface{}{"query": "up"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestQueryValue_Post(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/prometheus/value" {
			t.Errorf("expected path /prometheus/value, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": 42})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := QueryValue(context.Background(), client, map[string]interface{}{
		"param": map[string]interface{}{"query": "up"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestNodeList_Success(t *testing.T) {
	expected := map[string]interface{}{
		"count": float64(2),
		"data": []interface{}{
			map[string]interface{}{"name": "node1"},
			map[string]interface{}{"name": "node2"},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/database/node/list" {
			t.Errorf("expected path /database/node/list, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expected)
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := NodeList(context.Background(), client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}
	if parsed["count"] != float64(2) {
		t.Errorf("expected count 2, got %v", parsed["count"])
	}
}

func TestGetDashboard_Grafana(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/proxy/grafana/dashboards/uid/abc123" {
			t.Errorf("expected path /proxy/grafana/dashboards/uid/abc123, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"dashboard": map[string]interface{}{"uid": "abc123"}})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := GetDashboard(context.Background(), client, "abc123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestFolders_Grafana(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/proxy/grafana/folders" {
			t.Errorf("expected path /proxy/grafana/folders, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode([]interface{}{
			map[string]interface{}{"id": 1, "title": "General"},
		})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := Folders(context.Background(), client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestStatus_Grafana(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/proxy/grafana/status" {
			t.Errorf("expected path /proxy/grafana/status, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"status": "ok"})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := Status(context.Background(), client)
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
