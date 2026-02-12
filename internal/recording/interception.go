package recording

import (
	"context"
	"encoding/json"

	"hepic-cli/internal/api"
	"hepic-cli/internal/models"
)

// ListInterceptions performs a GET to /interceptions to list active interceptions.
func ListInterceptions(ctx context.Context, client *api.Client) (json.RawMessage, error) {
	var result json.RawMessage
	if err := client.Get(ctx, "/interceptions", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateInterception performs a POST to /interceptions to create a new interception.
func CreateInterception(ctx context.Context, client *api.Client, data models.InterceptionsStruct) (json.RawMessage, error) {
	var result json.RawMessage
	if err := client.Post(ctx, "/interceptions", data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateInterception performs a PUT to /interceptions/{uuid} to update an interception.
func UpdateInterception(ctx context.Context, client *api.Client, uuid string, data models.InterceptionsStruct) (json.RawMessage, error) {
	var result json.RawMessage
	if err := client.Put(ctx, "/interceptions/"+uuid, data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteInterception performs a DELETE to /interceptions/{uuid} to remove an interception.
func DeleteInterception(ctx context.Context, client *api.Client, uuid string) (json.RawMessage, error) {
	var result json.RawMessage
	if err := client.Delete(ctx, "/interceptions/"+uuid, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// AddRTPRecord performs a POST to /interception/add/rtprecord to add an RTP recording.
func AddRTPRecord(ctx context.Context, client *api.Client, data interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	if err := client.Post(ctx, "/interception/add/rtprecord", data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
