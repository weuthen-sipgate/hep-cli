package call

import (
	"context"
	"encoding/json"

	"hepic-cli/internal/api"
	"hepic-cli/internal/models"
)

// GetTransaction retrieves transaction details via POST /call/transaction.
func GetTransaction(ctx context.Context, client *api.Client, params SearchParams) (*models.SearchTransactionResponse, error) {
	var result models.SearchTransactionResponse
	if err := client.Post(ctx, "/call/transaction", params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SearchByType searches transactions by type via POST /search/transaction/type.
func SearchByType(ctx context.Context, client *api.Client, params SearchParams) (json.RawMessage, error) {
	var result json.RawMessage
	if err := client.Post(ctx, "/search/transaction/type", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
