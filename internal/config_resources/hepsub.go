package config_resources

import (
	"context"
	"encoding/json"

	"hepic-cli/internal/api"
)

// ListHepsub retrieves all HEP subscriptions.
// GET /hepsub/protocol
func ListHepsub(ctx context.Context, client *api.Client) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/hepsub/protocol", &result)
	return result, err
}

// CreateHepsub creates a new HEP subscription.
// POST /hepsub/protocol
func CreateHepsub(ctx context.Context, client *api.Client, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Post(ctx, "/hepsub/protocol", data, &result)
	return result, err
}

// UpdateHepsub updates an existing HEP subscription by UUID.
// PUT /hepsub/protocol/{uuid}
func UpdateHepsub(ctx context.Context, client *api.Client, uuid string, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Put(ctx, "/hepsub/protocol/"+api.PathEscape(uuid), data, &result)
	return result, err
}

// DeleteHepsub deletes a HEP subscription by UUID.
// DELETE /hepsub/protocol/{uuid}
func DeleteHepsub(ctx context.Context, client *api.Client, uuid string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Delete(ctx, "/hepsub/protocol/"+api.PathEscape(uuid), &result)
	return result, err
}

// SearchHepsub searches HEP subscription data.
// POST /hepsub/search
func SearchHepsub(ctx context.Context, client *api.Client, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Post(ctx, "/hepsub/search", data, &result)
	return result, err
}
