// Package client provides a thin Go client for the HDGP Engine API.
// Use it to call /hdgp/v1/evaluate from Go applications.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/HumanDignityGuardian/HDGP-Core/internal/engine"
)

// Client calls the HDGP Engine API.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// DefaultTimeout is the default HTTP client timeout (10s).
const DefaultTimeout = 10 * time.Second

// New creates a Client with default HTTP client.
func New(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}
}

// Evaluate sends the request to the Engine and returns the response.
func (c *Client) Evaluate(ctx context.Context, req engine.EvaluateRequest, includeReport bool) (*engine.EvaluateResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	url := c.BaseURL + "/hdgp/v1/evaluate"
	if includeReport {
		url += "?include_report=true"
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("engine returned %d: %s", resp.StatusCode, string(b))
	}

	var evalResp engine.EvaluateResponse
	if err := json.NewDecoder(resp.Body).Decode(&evalResp); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}
	return &evalResp, nil
}

// FallbackAction defines what to return when Engine is unavailable (fail-closed).
type FallbackAction string

const (
	// FallbackBlock returns a synthetic response with verdict=block when Engine fails.
	FallbackBlock FallbackAction = "block"
	// FallbackAllow returns verdict=allow (use only when acceptable for low-risk paths).
	FallbackAllow FallbackAction = "allow"
)

// EvaluateWithFallback calls Evaluate; on timeout or error, returns a synthetic
// response per FallbackAction (default: block) and sets EngineInfo to signal engine_unavailable.
func (c *Client) EvaluateWithFallback(ctx context.Context, req engine.EvaluateRequest, opts *FallbackOptions) (*engine.EvaluateResponse, error) {
	if opts == nil {
		opts = &FallbackOptions{}
	}
	action := opts.OnFailure
	if action == "" {
		action = FallbackBlock
	}

	resp, err := c.Evaluate(ctx, req, opts.IncludeReport)
	if err == nil {
		return resp, nil
	}

	verdict := "block"
	if action == FallbackAllow {
		verdict = "allow"
	}
	return &engine.EvaluateResponse{
		RequestID: req.Meta.RequestID,
		Verdict:   verdict,
		EngineInfo: engine.EngineInfo{
			Version: "engine_unavailable",
			Policy:  req.Meta.Policy,
		},
	}, nil
}

// FallbackOptions configures EvaluateWithFallback.
type FallbackOptions struct {
	IncludeReport bool          // request report in response
	OnFailure     FallbackAction // block or allow when Engine fails
}
