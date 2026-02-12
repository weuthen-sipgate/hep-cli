package admin

import (
	"context"
	"encoding/json"
	"fmt"

	"hepic-cli/internal/api"
)

// Profiles retrieves admin profiles. GET /admin/profiles
func Profiles(ctx context.Context, client *api.Client) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/admin/profiles", &result)
	return result, err
}

// ConfigDBTables retrieves the list of config DB tables. GET /configdb/tables/list
func ConfigDBTables(ctx context.Context, client *api.Client) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/configdb/tables/list", &result)
	return result, err
}

// ConfigDBResync triggers a resync of the config DB. POST /configdb/tables/resync
func ConfigDBResync(ctx context.Context, client *api.Client) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Post(ctx, "/configdb/tables/resync", nil, &result)
	return result, err
}

// APIVersion retrieves the API version info. GET /version/api/info
func APIVersion(ctx context.Context, client *api.Client) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/version/api/info", &result)
	return result, err
}

// UIVersion retrieves the UI version info. GET /version/ui/info
func UIVersion(ctx context.Context, client *api.Client) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/version/ui/info", &result)
	return result, err
}

// TroubleshootingLog retrieves troubleshooting logs. GET /troubleshooting/log/{type}/{action}
func TroubleshootingLog(ctx context.Context, client *api.Client, logType, action string) (json.RawMessage, error) {
	var result json.RawMessage
	path := fmt.Sprintf("/troubleshooting/log/%s/%s", logType, action)
	err := client.Get(ctx, path, &result)
	return result, err
}

// ImportPCAP imports a PCAP file. POST /import/data/pcap (file upload)
func ImportPCAP(ctx context.Context, client *api.Client, filePath string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.PostFormFile(ctx, "/import/data/pcap", "file", filePath, &result)
	return result, err
}

// ImportPCAPNow imports a PCAP file immediately. POST /import/data/pcap/now (file upload)
func ImportPCAPNow(ctx context.Context, client *api.Client, filePath string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.PostFormFile(ctx, "/import/data/pcap/now", "file", filePath, &result)
	return result, err
}

// ClickhouseQuery executes a raw ClickHouse query. POST /clickhouse/query/raw
func ClickhouseQuery(ctx context.Context, client *api.Client, query string) (json.RawMessage, error) {
	var result json.RawMessage
	body := map[string]string{"query": query}
	err := client.Post(ctx, "/clickhouse/query/raw", body, &result)
	return result, err
}

// ShareReport shares a call report. POST /share/call/report/{type}/{uuid}
func ShareReport(ctx context.Context, client *api.Client, reportType, uuid string) (json.RawMessage, error) {
	var result json.RawMessage
	path := fmt.Sprintf("/share/call/report/%s/%s", reportType, uuid)
	err := client.Post(ctx, path, nil, &result)
	return result, err
}

// ShareTransaction shares a call transaction. POST /share/call/transaction/{uuid}
func ShareTransaction(ctx context.Context, client *api.Client, uuid string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Post(ctx, "/share/call/transaction/"+uuid, nil, &result)
	return result, err
}

// SharePCAP shares a PCAP export. POST /share/export/call/messages/pcap/{uuid}
func SharePCAP(ctx context.Context, client *api.Client, uuid string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Post(ctx, "/share/export/call/messages/pcap/"+uuid, nil, &result)
	return result, err
}

// ShareText shares a text export. POST /share/export/call/messages/text/{uuid}
func ShareText(ctx context.Context, client *api.Client, uuid string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Post(ctx, "/share/export/call/messages/text/"+uuid, nil, &result)
	return result, err
}

// ShareIPAlias retrieves a shared IP alias. GET /share/ipalias/{uuid}
func ShareIPAlias(ctx context.Context, client *api.Client, uuid string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/share/ipalias/"+uuid, &result)
	return result, err
}

// ShareMapping retrieves a shared mapping. GET /share/mapping/protocol/{uuid}
func ShareMapping(ctx context.Context, client *api.Client, uuid string) (json.RawMessage, error) {
	var result json.RawMessage
	err := client.Get(ctx, "/share/mapping/protocol/"+uuid, &result)
	return result, err
}
