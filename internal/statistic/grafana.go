package statistic

import (
	"context"
	"encoding/json"
	"hepic-cli/internal/api"
)

// GetDashboard retrieves a Grafana dashboard by UID. GET /proxy/grafana/dashboards/uid/{uid}
func GetDashboard(ctx context.Context, client *api.Client, uid string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/proxy/grafana/dashboards/uid/"+uid, &result)
	return result, err
}

// Folders retrieves all Grafana folders. GET /proxy/grafana/folders
func Folders(ctx context.Context, client *api.Client) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/proxy/grafana/folders", &result)
	return result, err
}

// Org retrieves the Grafana organization info. GET /proxy/grafana/org
func Org(ctx context.Context, client *api.Client) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/proxy/grafana/org", &result)
	return result, err
}

// SearchDashboard searches for a Grafana dashboard by UID. GET /proxy/grafana/search/{uid}
func SearchDashboard(ctx context.Context, client *api.Client, uid string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/proxy/grafana/search/"+uid, &result)
	return result, err
}

// Status retrieves the Grafana connection status. GET /proxy/grafana/status
func Status(ctx context.Context, client *api.Client) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/proxy/grafana/status", &result)
	return result, err
}
