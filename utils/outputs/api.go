package outputs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"monitor/models"
	"net/http"
	"strings"
	"time"
)

type APIOutputHandler struct {
	url    string
	apiKey string
	method string
	client *http.Client
}

func NewAPIOutputHandler(url, apiKey, method string) (*APIOutputHandler, error) {
	if url == "" {
		return nil, fmt.Errorf("API URL cannot be empty")
	}

	if method == "" {
		method = http.MethodPost
	}

	method = strings.ToUpper(method)

	switch method {
	case http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch:
	default:
		return nil, fmt.Errorf("unsupported HTTP method: %s", method)
	}

	return &APIOutputHandler{
		url:    url,
		apiKey: apiKey,
		method: method,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}, nil
}

func (a *APIOutputHandler) Write(info *models.SystemInfo) error {
	data, err := json.Marshal(info)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	req, err := http.NewRequest(a.method, a.url, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if a.apiKey != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.apiKey))
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

func (a *APIOutputHandler) Name() string {
	return fmt.Sprintf("API Output: %s [%s]", a.url, a.method)
} 