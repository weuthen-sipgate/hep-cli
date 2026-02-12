package call

import (
	"context"
	"encoding/json"

	"hepic-cli/internal/api"
)

// ReportDTMF retrieves a DTMF report via POST /call/report/dtmf.
func ReportDTMF(ctx context.Context, client *api.Client, params SearchParams) (json.RawMessage, error) {
	var result json.RawMessage
	if err := client.Post(ctx, "/call/report/dtmf", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// ReportLog retrieves a call log report via POST /call/report/log.
func ReportLog(ctx context.Context, client *api.Client, params SearchParams) (json.RawMessage, error) {
	var result json.RawMessage
	if err := client.Post(ctx, "/call/report/log", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// ReportQOS retrieves a QoS report via POST /call/report/qos.
func ReportQOS(ctx context.Context, client *api.Client, params SearchParams) (json.RawMessage, error) {
	var result json.RawMessage
	if err := client.Post(ctx, "/call/report/qos", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
