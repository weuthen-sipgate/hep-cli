package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/spf13/viper"
)

// Client is the central HTTP client for the HEPIC API.
type Client struct {
	BaseURL    string
	Token      string
	HTTPClient *http.Client
	Verbose    bool
}

// NewClient creates a Client from the current viper configuration.
func NewClient() (*Client, error) {
	host := viper.GetString("host")
	if host == "" {
		return nil, fmt.Errorf("host is not configured. Run 'hepic init' or set HEPIC_HOST")
	}
	token := viper.GetString("token")
	if token == "" {
		return nil, fmt.Errorf("token is not configured. Run 'hepic init' or set HEPIC_TOKEN")
	}

	return &Client{
		BaseURL:    host + "/api/v3",
		Token:      token,
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
		Verbose:    viper.GetBool("verbose"),
	}, nil
}

// NewClientWith creates a Client with explicit parameters (useful for testing).
func NewClientWith(baseURL, token string) *Client {
	return &Client{
		BaseURL:    baseURL,
		Token:      token,
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
	}
}

// Get performs a GET request and decodes the JSON response into result.
func (c *Client) Get(ctx context.Context, path string, result interface{}) error {
	return c.do(ctx, http.MethodGet, path, nil, result)
}

// Post performs a POST request with a JSON body and decodes the response.
func (c *Client) Post(ctx context.Context, path string, body interface{}, result interface{}) error {
	return c.do(ctx, http.MethodPost, path, body, result)
}

// Put performs a PUT request with a JSON body and decodes the response.
func (c *Client) Put(ctx context.Context, path string, body interface{}, result interface{}) error {
	return c.do(ctx, http.MethodPut, path, body, result)
}

// Delete performs a DELETE request and decodes the response.
func (c *Client) Delete(ctx context.Context, path string, result interface{}) error {
	return c.do(ctx, http.MethodDelete, path, nil, result)
}

// GetRaw performs a GET request and returns the raw response body (for binary data like PCAP).
func (c *Client) GetRaw(ctx context.Context, path string) (io.ReadCloser, error) {
	req, err := c.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		return nil, c.parseError(resp)
	}
	return resp.Body, nil
}

// PostRaw performs a POST request with a JSON body and returns the raw response body.
func (c *Client) PostRaw(ctx context.Context, path string, body interface{}) (io.ReadCloser, error) {
	req, err := c.newRequest(ctx, http.MethodPost, path, body)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		return nil, c.parseError(resp)
	}
	return resp.Body, nil
}

func (c *Client) do(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	req, err := c.newRequest(ctx, method, path, body)
	if err != nil {
		return err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if c.Verbose {
		fmt.Fprintf(os.Stderr, "[verbose] %s %s → %d\n", method, req.URL.String(), resp.StatusCode)
	}

	if resp.StatusCode >= 400 {
		return c.parseError(resp)
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}
	return nil
}

func (c *Client) newRequest(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {
	url := c.BaseURL + path

	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Auth-Token", c.Token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if c.Verbose {
		fmt.Fprintf(os.Stderr, "[verbose] %s %s\n", method, url)
	}

	return req, nil
}

func (c *Client) parseError(resp *http.Response) error {
	var apiErr APIError
	if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
		return &APIError{
			StatusCode: resp.StatusCode,
			ErrorText:  http.StatusText(resp.StatusCode),
			Message:    "failed to parse error response",
		}
	}
	if apiErr.StatusCode == 0 {
		apiErr.StatusCode = resp.StatusCode
	}
	if apiErr.ErrorText == "" {
		apiErr.ErrorText = http.StatusText(resp.StatusCode)
	}
	return &apiErr
}
