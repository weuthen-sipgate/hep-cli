package user

import (
	"context"
	"encoding/json"

	"hepic-cli/internal/api"
)

// CreateToken creates a new auth token. POST /token/auth
func CreateToken(ctx context.Context, client *api.Client, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Post(ctx, "/token/auth", data, &result)
	return result, err
}

// GetToken retrieves a specific auth token. GET /token/auth/{uuid}
func GetToken(ctx context.Context, client *api.Client, uuid string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/token/auth/"+uuid, &result)
	return result, err
}

// DeleteToken deletes an auth token. DELETE /token/auth/{uuid}
func DeleteToken(ctx context.Context, client *api.Client, uuid string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Delete(ctx, "/token/auth/"+uuid, &result)
	return result, err
}
