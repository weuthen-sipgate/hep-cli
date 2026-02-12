package statistic

import (
	"context"
	"encoding/json"
	"hepic-cli/internal/api"
)

// QueryData queries Prometheus data. POST /prometheus/data
func QueryData(ctx context.Context, client *api.Client, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Post(ctx, "/prometheus/data", data, &result)
	return result, err
}

// QueryValue queries a Prometheus metric value. POST /prometheus/value
func QueryValue(ctx context.Context, client *api.Client, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Post(ctx, "/prometheus/value", data, &result)
	return result, err
}

// Labels retrieves all available Prometheus labels. GET /prometheus/labels
func Labels(ctx context.Context, client *api.Client) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/prometheus/labels", &result)
	return result, err
}

// LabelDetail retrieves details for a specific Prometheus label. GET /prometheus/label/{userlabel}
func LabelDetail(ctx context.Context, client *api.Client, label string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/prometheus/label/"+label, &result)
	return result, err
}
