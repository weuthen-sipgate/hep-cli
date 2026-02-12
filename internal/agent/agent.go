package agent

import (
	"context"
	"encoding/json"

	"hepic-cli/internal/api"
)

// List retrieves all registered capture agents.
// GET /agent/subscribe
func List(ctx context.Context, client *api.Client) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/agent/subscribe", &result)
	return result, err
}

// Get retrieves a single agent by UUID.
// GET /agent/subscribe/{uuid}
func Get(ctx context.Context, client *api.Client, uuid string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/agent/subscribe/"+uuid, &result)
	return result, err
}

// Update modifies an existing agent.
// PUT /agent/subscribe/{uuid}
func Update(ctx context.Context, client *api.Client, uuid string, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Put(ctx, "/agent/subscribe/"+uuid, data, &result)
	return result, err
}

// Delete removes an agent by UUID.
// DELETE /agent/subscribe/{uuid}
func Delete(ctx context.Context, client *api.Client, uuid string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Delete(ctx, "/agent/subscribe/"+uuid, &result)
	return result, err
}

// Search searches for agents by GUID and type.
// POST /agent/search/{guid}/{type}
func Search(ctx context.Context, client *api.Client, guid, agentType string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Post(ctx, "/agent/search/"+guid+"/"+agentType, nil, &result)
	return result, err
}

// ListByType retrieves agents filtered by type.
// GET /agent/type/{type}
func ListByType(ctx context.Context, client *api.Client, agentType string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/agent/type/"+agentType, &result)
	return result, err
}

// AddSubscription creates a new agent subscription protocol.
// POST /agentsub/protocol
func AddSubscription(ctx context.Context, client *api.Client, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Post(ctx, "/agentsub/protocol", data, &result)
	return result, err
}
