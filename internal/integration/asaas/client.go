package asaas

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Client struct {
	APIKey string
	APIUrl string
}

func NewClient(apiKey string) *Client {
	apiUrl := os.Getenv("ASAAS_API_URL")
	if apiUrl == "" {
		apiUrl = "https://www.asaas.com/api/v3"
	}

	return &Client{
		APIKey: apiKey,
		APIUrl: apiUrl,
	}
}

func (c *Client) DoRequest(ctx context.Context, method, endpoint string, body any) ([]byte, error) {
	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("Asaas.DoRequest error: %w", err)
		}

		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.APIUrl+endpoint, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("Asaas.DoRequest error: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("access_token", c.APIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Asaas.DoRequest error: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Asaas.DoRequest error: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("Asaas.DoRequest error: status %d, response: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}
