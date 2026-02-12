package statistic

import (
	"context"
	"encoding/json"
	"hepic-cli/internal/api"
)

// DBStats retrieves database statistics. GET /statistic/_db
func DBStats(ctx context.Context, client *api.Client) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/statistic/_db", &result)
	return result, err
}

// Measurements queries measurements for a given database ID. POST /statistic/_measurements/{dbid}
func Measurements(ctx context.Context, client *api.Client, dbid string, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Post(ctx, "/statistic/_measurements/"+api.PathEscape(dbid), data, &result)
	return result, err
}

// Metrics queries available metrics. POST /statistic/_metrics
func Metrics(ctx context.Context, client *api.Client, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Post(ctx, "/statistic/_metrics", data, &result)
	return result, err
}

// Retentions queries retention policies. POST /statistic/_retentions
func Retentions(ctx context.Context, client *api.Client, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Post(ctx, "/statistic/_retentions", data, &result)
	return result, err
}

// Data queries statistical data with timestamp parameters. POST /statistic/data
func Data(ctx context.Context, client *api.Client, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Post(ctx, "/statistic/data", data, &result)
	return result, err
}
