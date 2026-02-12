package config_resources

import (
	"context"
	"encoding/json"
	"io"

	"hepic-cli/internal/api"
)

// ListAliases retrieves all IP aliases.
// GET /ipalias
func ListAliases(ctx context.Context, client *api.Client) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/ipalias", &result)
	return result, err
}

// CreateAlias creates a new IP alias.
// POST /ipalias
func CreateAlias(ctx context.Context, client *api.Client, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Post(ctx, "/ipalias", data, &result)
	return result, err
}

// UpdateAlias updates an existing IP alias by UUID.
// PUT /ipalias/{uuid}
func UpdateAlias(ctx context.Context, client *api.Client, uuid string, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Put(ctx, "/ipalias/"+api.PathEscape(uuid), data, &result)
	return result, err
}

// DeleteAlias deletes an IP alias by UUID.
// DELETE /ipalias/{uuid}
func DeleteAlias(ctx context.Context, client *api.Client, uuid string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Delete(ctx, "/ipalias/"+api.PathEscape(uuid), &result)
	return result, err
}

// DeleteAllAliases deletes all IP aliases.
// DELETE /ipalias/all
func DeleteAllAliases(ctx context.Context, client *api.Client) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Delete(ctx, "/ipalias/all", &result)
	return result, err
}

// ExportAliases exports all IP aliases as CSV.
// GET /ipalias/export
func ExportAliases(ctx context.Context, client *api.Client) (io.ReadCloser, error) {
	return client.GetRaw(ctx, "/ipalias/export")
}

// ImportAliases imports IP aliases from a CSV file.
// POST /ipalias/import
func ImportAliases(ctx context.Context, client *api.Client, filePath string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.PostFormFile(ctx, "/ipalias/import", "file", filePath, &result)
	return result, err
}
