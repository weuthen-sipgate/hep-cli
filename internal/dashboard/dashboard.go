package dashboard

import (
	"context"
	"encoding/json"
	"hepic-cli/internal/api"
)

// List retrieves all dashboards. GET /dashboard/info
func List(ctx context.Context, client *api.Client) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/dashboard/info", &result)
	return result, err
}

// Store creates or updates a dashboard. PUT /dashboard/store/{dashboardId}
func Store(ctx context.Context, client *api.Client, dashboardID string, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Put(ctx, "/dashboard/store/"+api.PathEscape(dashboardID), data, &result)
	return result, err
}

// Delete removes a dashboard. DELETE /dashboard/store/{dashboardId}
func Delete(ctx context.Context, client *api.Client, dashboardID string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Delete(ctx, "/dashboard/store/"+api.PathEscape(dashboardID), &result)
	return result, err
}
