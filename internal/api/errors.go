package api

import "fmt"

// APIError represents a structured error from the HEPIC API.
type APIError struct {
	StatusCode int    `json:"statuscode"`
	ErrorText  string `json:"error"`
	Message    string `json:"message"`
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("HTTP %d: %s — %s", e.StatusCode, e.ErrorText, e.Message)
	}
	return fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.ErrorText)
}
