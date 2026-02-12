package export

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"hepic-cli/internal/api"
)

func TestNewExportParams_WithCallID(t *testing.T) {
	params, err := NewExportParams("", "", "test-call-id")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	search, ok := params.Param["search"].(map[string]interface{})
	if !ok {
		t.Fatal("expected param.search to be a map")
	}
	callids, ok := search["callid"].([]string)
	if !ok {
		t.Fatal("expected param.search.callid to be a string slice")
	}
	if len(callids) != 1 || callids[0] != "test-call-id" {
		t.Errorf("expected callid [test-call-id], got %v", callids)
	}

	if params.Timestamp != nil {
		t.Errorf("expected nil timestamp when from/to are empty, got %v", params.Timestamp)
	}
}

func TestNewExportParams_WithTimestamps(t *testing.T) {
	params, err := NewExportParams("2025-01-01", "2025-01-02", "call-123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if params.Timestamp == nil {
		t.Fatal("expected timestamp to be set")
	}
	if _, ok := params.Timestamp["from"]; !ok {
		t.Error("expected timestamp.from to be set")
	}
	if _, ok := params.Timestamp["to"]; !ok {
		t.Error("expected timestamp.to to be set")
	}
}

func TestNewExportParams_WithRFC3339(t *testing.T) {
	params, err := NewExportParams("2025-06-15T10:30:00Z", "", "call-456")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if params.Timestamp == nil {
		t.Fatal("expected timestamp to be set")
	}
	fromMs, ok := params.Timestamp["from"].(int64)
	if !ok {
		t.Fatal("expected from to be int64")
	}
	if fromMs <= 0 {
		t.Errorf("expected positive unix ms, got %d", fromMs)
	}
}

func TestNewExportParams_WithUnixMs(t *testing.T) {
	params, err := NewExportParams("1735689600000", "", "call-789")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if params.Timestamp == nil {
		t.Fatal("expected timestamp to be set")
	}
	fromMs, ok := params.Timestamp["from"].(int64)
	if !ok {
		t.Fatal("expected from to be int64")
	}
	if fromMs != 1735689600000 {
		t.Errorf("expected 1735689600000, got %d", fromMs)
	}
}

func TestNewExportParams_InvalidTime(t *testing.T) {
	_, err := NewExportParams("not-a-date", "", "call-id")
	if err == nil {
		t.Fatal("expected error for invalid date")
	}
}

func TestExportPCAPData_ReturnsBinaryData(t *testing.T) {
	pcapData := []byte{0xd4, 0xc3, 0xb2, 0xa1, 0x02, 0x00, 0x04, 0x00} // fake PCAP magic bytes
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/export/call/data/pcap" {
			t.Errorf("expected path /export/call/data/pcap, got %s", r.URL.Path)
		}
		if r.Header.Get("Auth-Token") != "test-token" {
			t.Errorf("expected Auth-Token test-token, got %s", r.Header.Get("Auth-Token"))
		}

		// Verify request body contains expected params
		var body ExportParams
		json.NewDecoder(r.Body).Decode(&body)
		if body.Param == nil {
			t.Error("expected param in request body")
		}

		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(pcapData)
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	params, _ := NewExportParams("", "", "test-call")

	body, err := ExportPCAPData(context.Background(), client, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer body.Close()

	data, err := io.ReadAll(body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}
	if len(data) != len(pcapData) {
		t.Errorf("expected %d bytes, got %d", len(pcapData), len(data))
	}
	for i, b := range pcapData {
		if data[i] != b {
			t.Errorf("byte %d: expected 0x%02x, got 0x%02x", i, b, data[i])
			break
		}
	}
}

func TestExportText_ReturnsTextData(t *testing.T) {
	textContent := "SIP/2.0 200 OK\r\nFrom: <sip:alice@example.com>\r\nTo: <sip:bob@example.com>\r\n"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/export/call/messages/text" {
			t.Errorf("expected path /export/call/messages/text, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(textContent))
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	params, _ := NewExportParams("", "", "test-call")

	body, err := ExportText(context.Background(), client, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer body.Close()

	data, err := io.ReadAll(body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}
	if string(data) != textContent {
		t.Errorf("expected %q, got %q", textContent, string(data))
	}
}

func TestExportMessagesPCAP_CorrectPath(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/export/call/messages/pcap" {
			t.Errorf("expected path /export/call/messages/pcap, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pcap-data"))
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	params, _ := NewExportParams("", "", "call-id")
	body, err := ExportMessagesPCAP(context.Background(), client, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	body.Close()
}

func TestExportSIPP_CorrectPath(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/export/call/messages/sipp" {
			t.Errorf("expected path /export/call/messages/sipp, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<sipp-scenario/>"))
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	params, _ := NewExportParams("", "", "call-id")
	body, err := ExportSIPP(context.Background(), client, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	body.Close()
}

func TestExportStenographer_CorrectPath(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/export/call/stenographer" {
			t.Errorf("expected path /export/call/stenographer, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("steno-data"))
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	params, _ := NewExportParams("", "", "call-id")
	body, err := ExportStenographer(context.Background(), client, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	body.Close()
}

func TestExportTransactionReport_CorrectPath(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/export/call/transaction/report" {
			t.Errorf("expected path /export/call/transaction/report, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("report-data"))
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	params, _ := NewExportParams("", "", "call-id")
	body, err := ExportTransactionReport(context.Background(), client, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	body.Close()
}

func TestExportTransactionArchive_CorrectPath(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/export/call/transaction/archive" {
			t.Errorf("expected path /export/call/transaction/archive, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("archive-data"))
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	params, _ := NewExportParams("", "", "call-id")
	body, err := ExportTransactionArchive(context.Background(), client, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	body.Close()
}

func TestExportTransactionLink_ReturnsJSON(t *testing.T) {
	expectedResp := `{"url":"https://hepic.example.com/share/abc123"}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/export/call/transaction/link" {
			t.Errorf("expected path /export/call/transaction/link, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(expectedResp))
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	params, _ := NewExportParams("", "", "call-id")
	result, err := ExportTransactionLink(context.Background(), client, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(result) != expectedResp {
		t.Errorf("expected %s, got %s", expectedResp, string(result))
	}
}

func TestExportAction_ReturnsJSON(t *testing.T) {
	expectedResp := `{"data":{"active":true,"count":42}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/export/action/active" {
			t.Errorf("expected path /export/action/active, got %s", r.URL.Path)
		}
		if r.Header.Get("Auth-Token") != "test-token" {
			t.Errorf("expected Auth-Token test-token, got %s", r.Header.Get("Auth-Token"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(expectedResp))
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := ExportAction(context.Background(), client, "active")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(result) != expectedResp {
		t.Errorf("expected %s, got %s", expectedResp, string(result))
	}
}

func TestExportAction_Logs(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/export/action/logs" {
			t.Errorf("expected path /export/action/logs, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"logs":[]}`))
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := ExportAction(context.Background(), client, "logs")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Error("expected non-nil result")
	}
}

func TestExportAction_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal error"})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	_, err := ExportAction(context.Background(), client, "active")
	if err == nil {
		t.Fatal("expected error for 500 response")
	}
}

func TestExportPCAPData_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad request"})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	params, _ := NewExportParams("", "", "bad-call")
	_, err := ExportPCAPData(context.Background(), client, params)
	if err == nil {
		t.Fatal("expected error for 400 response")
	}
}
