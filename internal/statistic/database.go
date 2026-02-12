package statistic

import (
	"context"
	"encoding/json"
	"hepic-cli/internal/api"
)

// NodeList retrieves the list of database nodes. GET /database/node/list
func NodeList(ctx context.Context, client *api.Client) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/database/node/list", &result)
	return result, err
}
