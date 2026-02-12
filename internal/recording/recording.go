package recording

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"time"

	"hepic-cli/internal/api"
)

// SearchParams represents the parameters for a recording search request.
type SearchParams struct {
	Param     map[string]interface{} `json:"param"`
	Timestamp map[string]interface{} `json:"timestamp,omitempty"`
}

// NewSearchParams builds a SearchParams with from/to timestamps in unix milliseconds.
func NewSearchParams(from, to string) (SearchParams, error) {
	params := SearchParams{
		Param: make(map[string]interface{}),
	}

	ts := make(map[string]interface{})

	if from != "" {
		fromTime, err := parseTime(from)
		if err != nil {
			return params, fmt.Errorf("invalid --from time: %w", err)
		}
		ts["from"] = fromTime.UnixMilli()
	}

	if to != "" {
		toTime, err := parseTime(to)
		if err != nil {
			return params, fmt.Errorf("invalid --to time: %w", err)
		}
		ts["to"] = toTime.UnixMilli()
	}

	if len(ts) > 0 {
		params.Timestamp = ts
	}

	return params, nil
}

// parseTime tries to parse a time string in common formats.
func parseTime(s string) (time.Time, error) {
	// Try unix timestamp (milliseconds)
	if ms, err := strconv.ParseInt(s, 10, 64); err == nil {
		return time.UnixMilli(ms), nil
	}

	// Try common date/time formats
	formats := []string{
		time.RFC3339,
		"2006-01-02T15:04:05",
		"2006-01-02 15:04:05",
		"2006-01-02",
	}
	for _, f := range formats {
		if t, err := time.Parse(f, s); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unsupported time format: %s (use RFC3339, YYYY-MM-DD, or unix ms)", s)
}

// SearchData performs a POST to /call/recording/data to search for recordings.
func SearchData(ctx context.Context, client *api.Client, params SearchParams) (json.RawMessage, error) {
	var result json.RawMessage
	if err := client.Post(ctx, "/call/recording/data", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// Play performs a GET to /call/recording/play/{uuid} and returns the raw stream.
func Play(ctx context.Context, client *api.Client, uuid string) (io.ReadCloser, error) {
	return client.GetRaw(ctx, "/call/recording/play/"+api.PathEscape(uuid))
}

// Download performs a GET to /call/recording/download/{type}/{uuid} and returns the raw stream.
// dlType should be "audio" or "pcap".
func Download(ctx context.Context, client *api.Client, dlType, uuid string) (io.ReadCloser, error) {
	return client.GetRaw(ctx, "/call/recording/download/"+api.PathEscape(dlType)+"/"+api.PathEscape(uuid))
}

// Info performs a GET to /call/recording/info/{uuid} and returns recording metadata.
func Info(ctx context.Context, client *api.Client, uuid string) (json.RawMessage, error) {
	var result json.RawMessage
	if err := client.Get(ctx, "/call/recording/info/"+api.PathEscape(uuid), &result); err != nil {
		return nil, err
	}
	return result, nil
}
