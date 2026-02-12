package call

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"hepic-cli/internal/api"
)

func TestSearchData(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/search/call/data" {
			t.Errorf("expected path /search/call/data, got %s", r.URL.Path)
		}
		if r.Header.Get("Auth-Token") != "test-token" {
			t.Errorf("expected Auth-Token test-token, got %s", r.Header.Get("Auth-Token"))
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected Content-Type application/json, got %s", r.Header.Get("Content-Type"))
		}

		// Verify request body
		var body SearchParams
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Errorf("failed to decode request body: %v", err)
		}
		if body.Timestamp == nil {
			t.Error("expected timestamp in request body")
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data": []map[string]interface{}{
				{
					"id":          1,
					"method":      "INVITE",
					"ruri_user":   "+49456",
					"srcIp":       "10.0.0.1",
					"dstIp":       "10.0.0.2",
					"create_date": 1704067200,
				},
			},
			"keys":  []string{"id", "method"},
			"total": 1,
		})
	}))
	defer srv.Close()

	client := api.NewClientWith(srv.URL, "test-token")
	params, err := NewSearchParams("2025-01-01", "2025-01-31", "", "+49456", "")
	if err != nil {
		t.Fatalf("NewSearchParams failed: %v", err)
	}

	result, err := SearchData(context.Background(), client, params)
	if err != nil {
		t.Fatalf("SearchData failed: %v", err)
	}

	if result.Total != 1 {
		t.Errorf("expected total=1, got %d", result.Total)
	}
	if len(result.Data) != 1 {
		t.Errorf("expected 1 data item, got %d", len(result.Data))
	}
}

func TestSearchMessage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/search/call/message" {
			t.Errorf("expected path /search/call/message, got %s", r.URL.Path)
		}
		if r.Header.Get("Auth-Token") != "test-token" {
			t.Errorf("expected Auth-Token test-token, got %s", r.Header.Get("Auth-Token"))
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data": []interface{}{"SIP/2.0 200 OK"},
		})
	}))
	defer srv.Close()

	client := api.NewClientWith(srv.URL, "test-token")
	params, err := NewSearchParams("2025-01-01", "", "", "", "test-call-id")
	if err != nil {
		t.Fatalf("NewSearchParams failed: %v", err)
	}

	result, err := SearchMessage(context.Background(), client, params)
	if err != nil {
		t.Fatalf("SearchMessage failed: %v", err)
	}

	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestDecodeMessage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/search/call/decode/message" {
			t.Errorf("expected path /search/call/decode/message, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data": []map[string]interface{}{
				{"method": "INVITE", "from": "sip:alice@example.com"},
			},
		})
	}))
	defer srv.Close()

	client := api.NewClientWith(srv.URL, "test-token")
	params, err := NewSearchParams("2025-01-01", "2025-01-31", "", "", "")
	if err != nil {
		t.Fatalf("NewSearchParams failed: %v", err)
	}

	result, err := DecodeMessage(context.Background(), client, params)
	if err != nil {
		t.Fatalf("DecodeMessage failed: %v", err)
	}

	if len(result.Data) != 1 {
		t.Errorf("expected 1 decoded message, got %d", len(result.Data))
	}
}

func TestGetTransaction(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/call/transaction" {
			t.Errorf("expected path /call/transaction, got %s", r.URL.Path)
		}
		if r.Header.Get("Auth-Token") != "test-token" {
			t.Errorf("expected Auth-Token test-token, got %s", r.Header.Get("Auth-Token"))
		}

		// Verify the request body contains the call-id
		var body SearchParams
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Errorf("failed to decode request body: %v", err)
		}
		search, ok := body.Param["search"].(map[string]interface{})
		if !ok {
			t.Error("expected search param in request body")
		} else if search["callid"] != "test-call-id-123" {
			t.Errorf("expected callid=test-call-id-123, got %v", search["callid"])
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"Data": []map[string]interface{}{
				{
					"id":     1,
					"method": "INVITE",
					"srcIp":  "10.0.0.1",
					"dstIp":  "10.0.0.2",
				},
				{
					"id":     2,
					"method": "200 OK",
					"srcIp":  "10.0.0.2",
					"dstIp":  "10.0.0.1",
				},
			},
			"keys":  []string{"id", "method"},
			"total": 2,
		})
	}))
	defer srv.Close()

	client := api.NewClientWith(srv.URL, "test-token")
	params, err := NewSearchParams("2025-01-01", "2025-01-31", "", "", "test-call-id-123")
	if err != nil {
		t.Fatalf("NewSearchParams failed: %v", err)
	}

	result, err := GetTransaction(context.Background(), client, params)
	if err != nil {
		t.Fatalf("GetTransaction failed: %v", err)
	}

	if result.Total != 2 {
		t.Errorf("expected total=2, got %d", result.Total)
	}
	if len(result.Data) != 2 {
		t.Errorf("expected 2 transaction items, got %d", len(result.Data))
	}
}

func TestReportQOS(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/call/report/qos" {
			t.Errorf("expected path /call/report/qos, got %s", r.URL.Path)
		}
		if r.Header.Get("Auth-Token") != "test-token" {
			t.Errorf("expected Auth-Token test-token, got %s", r.Header.Get("Auth-Token"))
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"rtcp": map[string]interface{}{
				"data": []interface{}{},
			},
			"rtp": map[string]interface{}{
				"data": []interface{}{},
			},
		})
	}))
	defer srv.Close()

	client := api.NewClientWith(srv.URL, "test-token")
	params, err := NewSearchParams("2025-01-01", "2025-01-31", "", "", "test-call-id")
	if err != nil {
		t.Fatalf("NewSearchParams failed: %v", err)
	}

	result, err := ReportQOS(context.Background(), client, params)
	if err != nil {
		t.Fatalf("ReportQOS failed: %v", err)
	}

	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestNewSearchParams(t *testing.T) {
	t.Run("basic date parsing", func(t *testing.T) {
		params, err := NewSearchParams("2025-01-01", "2025-01-31", "", "", "")
		if err != nil {
			t.Fatalf("NewSearchParams failed: %v", err)
		}
		if params.Timestamp == nil {
			t.Fatal("expected timestamp to be set")
		}
		if params.Timestamp["from"] == nil {
			t.Error("expected from timestamp")
		}
		if params.Timestamp["to"] == nil {
			t.Error("expected to timestamp")
		}
	})

	t.Run("RFC3339 date parsing", func(t *testing.T) {
		params, err := NewSearchParams("2025-01-01T10:00:00Z", "2025-01-31T23:59:59Z", "", "", "")
		if err != nil {
			t.Fatalf("NewSearchParams failed: %v", err)
		}
		if params.Timestamp["from"] == nil {
			t.Error("expected from timestamp")
		}
	})

	t.Run("empty to defaults to now", func(t *testing.T) {
		params, err := NewSearchParams("2025-01-01", "", "", "", "")
		if err != nil {
			t.Fatalf("NewSearchParams failed: %v", err)
		}
		if params.Timestamp["to"] == nil {
			t.Error("expected to timestamp to be set (default: now)")
		}
	})

	t.Run("caller and callee in orlogic", func(t *testing.T) {
		params, err := NewSearchParams("2025-01-01", "", "+49123", "+49456", "")
		if err != nil {
			t.Fatalf("NewSearchParams failed: %v", err)
		}
		orlogic, ok := params.Param["orlogic"].(map[string]interface{})
		if !ok {
			t.Fatal("expected orlogic in param")
		}
		if orlogic["from_user"] != "+49123" {
			t.Errorf("expected from_user=+49123, got %v", orlogic["from_user"])
		}
		if orlogic["ruri_user"] != "+49456" {
			t.Errorf("expected ruri_user=+49456, got %v", orlogic["ruri_user"])
		}
	})

	t.Run("call-id in search", func(t *testing.T) {
		params, err := NewSearchParams("2025-01-01", "", "", "", "abc123")
		if err != nil {
			t.Fatalf("NewSearchParams failed: %v", err)
		}
		search, ok := params.Param["search"].(map[string]interface{})
		if !ok {
			t.Fatal("expected search in param")
		}
		if search["callid"] != "abc123" {
			t.Errorf("expected callid=abc123, got %v", search["callid"])
		}
	})

	t.Run("invalid from date", func(t *testing.T) {
		_, err := NewSearchParams("not-a-date", "", "", "", "")
		if err == nil {
			t.Error("expected error for invalid from date")
		}
	})

	t.Run("invalid to date", func(t *testing.T) {
		_, err := NewSearchParams("2025-01-01", "not-a-date", "", "", "")
		if err == nil {
			t.Error("expected error for invalid to date")
		}
	})
}

func TestSearchDataAPIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"statuscode": 401,
			"error":      "Unauthorized",
			"message":    "invalid token",
		})
	}))
	defer srv.Close()

	client := api.NewClientWith(srv.URL, "bad-token")
	params, _ := NewSearchParams("2025-01-01", "2025-01-31", "", "", "")

	_, err := SearchData(context.Background(), client, params)
	if err == nil {
		t.Fatal("expected error for 401 response")
	}
}

func TestSearchRemote(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/search/remote/data" {
			t.Errorf("expected path /search/remote/data, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"Data": []map[string]interface{}{},
		})
	}))
	defer srv.Close()

	client := api.NewClientWith(srv.URL, "test-token")
	params, err := NewSearchParams("2025-01-01", "2025-01-31", "", "", "")
	if err != nil {
		t.Fatalf("NewSearchParams failed: %v", err)
	}

	result, err := SearchRemote(context.Background(), client, params)
	if err != nil {
		t.Fatalf("SearchRemote failed: %v", err)
	}

	if result == nil {
		t.Fatal("expected non-nil result")
	}
}
