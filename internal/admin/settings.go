package admin

import (
	"context"
	"encoding/json"

	"hepic-cli/internal/api"
)

// ListSettings retrieves all user settings. GET /user/settings
func ListSettings(ctx context.Context, client *api.Client) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/user/settings", &result)
	return result, err
}

// CreateSetting creates a new user setting. POST /user/settings
func CreateSetting(ctx context.Context, client *api.Client, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Post(ctx, "/user/settings", data, &result)
	return result, err
}

// GetSettingsByCategory retrieves settings by category. GET /user/settings/{category}
func GetSettingsByCategory(ctx context.Context, client *api.Client, category string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/user/settings/"+api.PathEscape(category), &result)
	return result, err
}

// UpdateSetting updates an existing user setting. PUT /user/settings/{uuid}
func UpdateSetting(ctx context.Context, client *api.Client, uuid string, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Put(ctx, "/user/settings/"+api.PathEscape(uuid), data, &result)
	return result, err
}

// DeleteSetting deletes a user setting. DELETE /user/settings/{uuid}
func DeleteSetting(ctx context.Context, client *api.Client, uuid string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Delete(ctx, "/user/settings/"+api.PathEscape(uuid), &result)
	return result, err
}

// ListAdvanced retrieves all advanced settings. GET /advanced
func ListAdvanced(ctx context.Context, client *api.Client) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/advanced", &result)
	return result, err
}

// CreateAdvanced creates a new advanced setting. POST /advanced
func CreateAdvanced(ctx context.Context, client *api.Client, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Post(ctx, "/advanced", data, &result)
	return result, err
}

// GetAdvanced retrieves an advanced setting by UUID. GET /advanced/{uuid}
func GetAdvanced(ctx context.Context, client *api.Client, uuid string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/advanced/"+api.PathEscape(uuid), &result)
	return result, err
}

// UpdateAdvanced updates an existing advanced setting. PUT /advanced/{uuid}
func UpdateAdvanced(ctx context.Context, client *api.Client, uuid string, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Put(ctx, "/advanced/"+api.PathEscape(uuid), data, &result)
	return result, err
}

// DeleteAdvanced deletes an advanced setting. DELETE /advanced/{uuid}
func DeleteAdvanced(ctx context.Context, client *api.Client, uuid string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Delete(ctx, "/advanced/"+api.PathEscape(uuid), &result)
	return result, err
}
