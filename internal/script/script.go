package script

import (
	"context"
	"encoding/json"

	"hepic-cli/internal/api"
)

// List retrieves all scripts. GET /script
func List(ctx context.Context, client *api.Client) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/script", &result)
	return result, err
}

// Create creates a new script. POST /script
func Create(ctx context.Context, client *api.Client, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Post(ctx, "/script", data, &result)
	return result, err
}

// Update updates an existing script. PUT /script/{uuid}
func Update(ctx context.Context, client *api.Client, uuid string, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Put(ctx, "/script/"+api.PathEscape(uuid), data, &result)
	return result, err
}

// Delete deletes a script. DELETE /script/{uuid}
func Delete(ctx context.Context, client *api.Client, uuid string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Delete(ctx, "/script/"+api.PathEscape(uuid), &result)
	return result, err
}
