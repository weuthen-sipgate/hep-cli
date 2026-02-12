package config_resources

import (
	"context"
	"encoding/json"

	"hepic-cli/internal/api"
)

// SearchProtocol searches for a protocol by ID.
// GET /protocol/search/{id}
func SearchProtocol(ctx context.Context, client *api.Client, id string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/protocol/search/"+api.PathEscape(id), &result)
	return result, err
}

// CreateProtocol creates a new protocol definition.
// POST /protocol/{id}
func CreateProtocol(ctx context.Context, client *api.Client, id string, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Post(ctx, "/protocol/"+api.PathEscape(id), data, &result)
	return result, err
}

// UpdateProtocol updates an existing protocol by UUID.
// PUT /protocol/{uuid}
func UpdateProtocol(ctx context.Context, client *api.Client, uuid string, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Put(ctx, "/protocol/"+api.PathEscape(uuid), data, &result)
	return result, err
}

// DeleteProtocol deletes a protocol by UUID.
// DELETE /protocol/{uuid}
func DeleteProtocol(ctx context.Context, client *api.Client, uuid string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Delete(ctx, "/protocol/"+api.PathEscape(uuid), &result)
	return result, err
}
