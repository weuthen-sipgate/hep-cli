package config_resources

import (
	"context"
	"encoding/json"

	"hepic-cli/internal/api"
)

// ListMappings retrieves all protocol mappings.
// GET /mapping/protocol
func ListMappings(ctx context.Context, client *api.Client) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/mapping/protocol", &result)
	return result, err
}

// CreateMapping creates a new protocol mapping.
// POST /mapping/protocol
func CreateMapping(ctx context.Context, client *api.Client, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Post(ctx, "/mapping/protocol", data, &result)
	return result, err
}

// UpdateMapping updates an existing protocol mapping by UUID.
// PUT /mapping/protocol/{uuid}
func UpdateMapping(ctx context.Context, client *api.Client, uuid string, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Put(ctx, "/mapping/protocol/"+api.PathEscape(uuid), data, &result)
	return result, err
}

// DeleteMapping deletes a protocol mapping by UUID.
// DELETE /mapping/protocol/{uuid}
func DeleteMapping(ctx context.Context, client *api.Client, uuid string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Delete(ctx, "/mapping/protocol/"+api.PathEscape(uuid), &result)
	return result, err
}

// ListAllProtocols retrieves all protocol definitions.
// GET /mapping/protocols
func ListAllProtocols(ctx context.Context, client *api.Client) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/mapping/protocols", &result)
	return result, err
}

// ResetAll resets all protocol mappings to defaults.
// GET /mapping/protocol/reset
func ResetAll(ctx context.Context, client *api.Client) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/mapping/protocol/reset", &result)
	return result, err
}

// ResetOne resets a single protocol mapping by UUID.
// GET /mapping/protocol/reset/{uuid}
func ResetOne(ctx context.Context, client *api.Client, uuid string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/mapping/protocol/reset/"+api.PathEscape(uuid), &result)
	return result, err
}
