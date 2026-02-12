package export

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"time"

	"hepic-cli/internal/api"
)

// ExportParams holds the request body for export API calls.
type ExportParams struct {
	Param     map[string]interface{} `json:"param"`
	Timestamp map[string]interface{} `json:"timestamp,omitempty"`
}

// NewExportParams builds an ExportParams from from/to timestamps and a call ID.
// from and to are parsed as RFC3339 or date-only (2006-01-02) and converted to
// Unix milliseconds. If from/to are empty, reasonable defaults are used.
func NewExportParams(from, to, callID string) (ExportParams, error) {
	params := ExportParams{
		Param: map[string]interface{}{
			"search": map[string]interface{}{
				"callid": []string{callID},
			},
		},
	}

	ts := make(map[string]interface{})
	if from != "" {
		fromMs, err := parseTimeToMillis(from)
		if err != nil {
			return params, fmt.Errorf("invalid --from value: %w", err)
		}
		ts["from"] = fromMs
	}
	if to != "" {
		toMs, err := parseTimeToMillis(to)
		if err != nil {
			return params, fmt.Errorf("invalid --to value: %w", err)
		}
		ts["to"] = toMs
	}
	if len(ts) > 0 {
		params.Timestamp = ts
	}

	return params, nil
}

// parseTimeToMillis parses a time string as RFC3339, date-only (2006-01-02),
// or a raw Unix millisecond value, and returns Unix milliseconds.
func parseTimeToMillis(s string) (int64, error) {
	// Try as raw number first (already milliseconds)
	if ms, err := strconv.ParseInt(s, 10, 64); err == nil {
		return ms, nil
	}
	// Try RFC3339
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		return t.UnixMilli(), nil
	}
	// Try date-only
	if t, err := time.Parse("2006-01-02", s); err == nil {
		return t.UnixMilli(), nil
	}
	return 0, fmt.Errorf("cannot parse %q as RFC3339, date (2006-01-02), or unix ms", s)
}

// ExportPCAPData exports call data as PCAP.
// POST /export/call/data/pcap
func ExportPCAPData(ctx context.Context, client *api.Client, params ExportParams) (io.ReadCloser, error) {
	return client.PostRaw(ctx, "/export/call/data/pcap", params)
}

// ExportMessagesPCAP exports messages as PCAP.
// POST /export/call/messages/pcap
func ExportMessagesPCAP(ctx context.Context, client *api.Client, params ExportParams) (io.ReadCloser, error) {
	return client.PostRaw(ctx, "/export/call/messages/pcap", params)
}

// ExportSIPP exports messages as SIPp format.
// POST /export/call/messages/sipp
func ExportSIPP(ctx context.Context, client *api.Client, params ExportParams) (io.ReadCloser, error) {
	return client.PostRaw(ctx, "/export/call/messages/sipp", params)
}

// ExportText exports messages as plain text.
// POST /export/call/messages/text
func ExportText(ctx context.Context, client *api.Client, params ExportParams) (io.ReadCloser, error) {
	return client.PostRaw(ctx, "/export/call/messages/text", params)
}

// ExportStenographer exports via stenographer.
// POST /export/call/stenographer
func ExportStenographer(ctx context.Context, client *api.Client, params ExportParams) (io.ReadCloser, error) {
	return client.PostRaw(ctx, "/export/call/stenographer", params)
}

// ExportTransactionReport exports a transaction report.
// POST /export/call/transaction/report
func ExportTransactionReport(ctx context.Context, client *api.Client, params ExportParams) (io.ReadCloser, error) {
	return client.PostRaw(ctx, "/export/call/transaction/report", params)
}

// ExportTransactionLink returns a shareable link for a transaction.
// POST /export/call/transaction/link
func ExportTransactionLink(ctx context.Context, client *api.Client, params ExportParams) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Post(ctx, "/export/call/transaction/link", params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExportTransactionArchive exports a transaction archive.
// POST /export/call/transaction/archive
func ExportTransactionArchive(ctx context.Context, client *api.Client, params ExportParams) (io.ReadCloser, error) {
	return client.PostRaw(ctx, "/export/call/transaction/archive", params)
}

// ExportAction retrieves action data by type.
// GET /export/action/{type}
// Valid types: active, hepicapp, logs, picserver, rtpagent
func ExportAction(ctx context.Context, client *api.Client, actionType string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/export/action/"+api.PathEscape(actionType), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
