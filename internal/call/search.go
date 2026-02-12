package call

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"hepic-cli/internal/api"
	"hepic-cli/internal/models"
)

// SearchParams represents the request body for HEPIC search endpoints.
type SearchParams struct {
	Param     map[string]interface{} `json:"param"`
	Timestamp map[string]interface{} `json:"timestamp,omitempty"`
}

// NewSearchParams builds a SearchParams from CLI flags.
// from and to are parsed as RFC3339 or "2006-01-02" date strings.
// If to is empty, the current time is used.
func NewSearchParams(from, to, caller, callee, callID string) (SearchParams, error) {
	params := SearchParams{
		Param: make(map[string]interface{}),
	}

	// Parse timestamps
	fromTime, err := parseTime(from)
	if err != nil {
		return params, fmt.Errorf("invalid --from value: %w", err)
	}

	var toTime time.Time
	if to == "" {
		toTime = time.Now()
	} else {
		toTime, err = parseTime(to)
		if err != nil {
			return params, fmt.Errorf("invalid --to value: %w", err)
		}
	}

	params.Timestamp = map[string]interface{}{
		"from": fromTime.UnixMilli(),
		"to":   toTime.UnixMilli(),
	}

	// Build search filters
	search := make(map[string]interface{})
	orlogic := make(map[string]interface{})

	if callID != "" {
		search["callid"] = callID
	}
	if caller != "" {
		orlogic["from_user"] = caller
	}
	if callee != "" {
		orlogic["ruri_user"] = callee
	}

	if len(search) > 0 {
		params.Param["search"] = search
	}
	if len(orlogic) > 0 {
		params.Param["orlogic"] = orlogic
	}

	return params, nil
}

// SearchData searches for call data via POST /search/call/data.
func SearchData(ctx context.Context, client *api.Client, params SearchParams) (*models.SearchCallData, error) {
	var result models.SearchCallData
	if err := client.Post(ctx, "/search/call/data", params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SearchMessage searches for call messages via POST /search/call/message.
func SearchMessage(ctx context.Context, client *api.Client, params SearchParams) (json.RawMessage, error) {
	var result json.RawMessage
	if err := client.Post(ctx, "/search/call/message", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// DecodeMessage decodes SIP messages via POST /search/call/decode/message.
func DecodeMessage(ctx context.Context, client *api.Client, params SearchParams) (*models.MessageDecoded, error) {
	var result models.MessageDecoded
	if err := client.Post(ctx, "/search/call/decode/message", params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SearchRemote performs a remote search via POST /search/remote/data.
func SearchRemote(ctx context.Context, client *api.Client, params SearchParams) (*models.RemoteResponseData, error) {
	var result models.RemoteResponseData
	if err := client.Post(ctx, "/search/remote/data", params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// parseTime parses a time string in RFC3339 or "2006-01-02" format.
func parseTime(s string) (time.Time, error) {
	// Try RFC3339 first
	t, err := time.Parse(time.RFC3339, s)
	if err == nil {
		return t, nil
	}

	// Try date-only format
	t, err = time.Parse("2006-01-02", s)
	if err == nil {
		return t, nil
	}

	return time.Time{}, fmt.Errorf("cannot parse %q as RFC3339 or YYYY-MM-DD", s)
}
