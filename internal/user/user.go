package user

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"hepic-cli/internal/api"
)

// List retrieves all users. GET /users
func List(ctx context.Context, client *api.Client) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/users", &result)
	return result, err
}

// Create creates a new user. POST /users
func Create(ctx context.Context, client *api.Client, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Post(ctx, "/users", data, &result)
	return result, err
}

// Update updates an existing user. PUT /users/{uuid}
func Update(ctx context.Context, client *api.Client, uuid string, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Put(ctx, "/users/"+uuid, data, &result)
	return result, err
}

// Delete deletes a user. DELETE /users/{uuid}
func Delete(ctx context.Context, client *api.Client, uuid string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Delete(ctx, "/users/"+uuid, &result)
	return result, err
}

// UpdatePassword changes a user's password. PUT /users/update/password/{uuid}
func UpdatePassword(ctx context.Context, client *api.Client, uuid string, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Put(ctx, "/users/update/password/"+uuid, data, &result)
	return result, err
}

// Import imports users from a CSV file. POST /users/import (multipart file upload)
func Import(ctx context.Context, client *api.Client, filePath string) (json.RawMessage, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}
	if _, err := io.Copy(part, file); err != nil {
		return nil, fmt.Errorf("failed to copy file data: %w", err)
	}
	writer.Close()

	url := client.BaseURL + "/users/import"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Auth-Token", client.Token)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Accept", "application/json")

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(respBody))
	}

	return json.RawMessage(respBody), nil
}

// Export exports users as CSV. GET /users/export
func Export(ctx context.Context, client *api.Client) (io.ReadCloser, error) {
	return client.GetRaw(ctx, "/users/export")
}

// Groups retrieves the list of user groups. GET /users/groups
func Groups(ctx context.Context, client *api.Client) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/users/groups", &result)
	return result, err
}

// AuthTypes retrieves the list of auth types. GET /auth/type/list
func AuthTypes(ctx context.Context, client *api.Client) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/auth/type/list", &result)
	return result, err
}
