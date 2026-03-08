package plaid

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"plaid-cli/internal/state"
)

type APIError struct {
	StatusCode     int    `json:"-"`
	RequestID      string `json:"request_id,omitempty"`
	ErrorType      string `json:"error_type,omitempty"`
	ErrorCode      string `json:"error_code,omitempty"`
	ErrorMessage   string `json:"error_message,omitempty"`
	DisplayMessage string `json:"display_message,omitempty"`
}

func (e *APIError) Error() string {
	parts := []string{"Plaid API error"}
	if e.StatusCode != 0 {
		parts = append(parts, fmt.Sprintf("status=%d", e.StatusCode))
	}
	if e.ErrorType != "" {
		parts = append(parts, fmt.Sprintf("type=%s", e.ErrorType))
	}
	if e.ErrorCode != "" {
		parts = append(parts, fmt.Sprintf("code=%s", e.ErrorCode))
	}
	if e.ErrorMessage != "" {
		parts = append(parts, fmt.Sprintf("message=%s", e.ErrorMessage))
	}
	if e.RequestID != "" {
		parts = append(parts, fmt.Sprintf("request_id=%s", e.RequestID))
	}
	return strings.Join(parts, " ")
}

type Client struct {
	baseURL    string
	clientID   string
	secret     string
	httpClient *http.Client
}

type BinaryResponse struct {
	Body    []byte
	Headers http.Header
}

type CreateHostedLinkTokenInput struct {
	ClientName   string
	Language     string
	CountryCodes []string
	ClientUserID string
	Products     []string
	Webhook      string
	RedirectURI  string
}

type LinkTokenCreateResponse struct {
	Expiration    string `json:"expiration"`
	HostedLinkURL string `json:"hosted_link_url"`
	LinkToken     string `json:"link_token"`
	RequestID     string `json:"request_id"`
}

type PublicTokenExchangeResponse struct {
	AccessToken string `json:"access_token"`
	ItemID      string `json:"item_id"`
	RequestID   string `json:"request_id"`
}

type ItemGetResponse struct {
	Item struct {
		ItemID        string `json:"item_id"`
		InstitutionID string `json:"institution_id"`
	} `json:"item"`
	RequestID string `json:"request_id"`
}

type Account struct {
	AccountID string `json:"account_id"`
	Name      string `json:"name"`
	Mask      string `json:"mask"`
	Type      string `json:"type"`
	Subtype   string `json:"subtype"`
}

func (a Account) GetAccountID() string { return a.AccountID }
func (a Account) GetName() string      { return a.Name }
func (a Account) GetMask() string      { return a.Mask }
func (a Account) GetType() string      { return a.Type }
func (a Account) GetSubtype() string   { return a.Subtype }

type AccountsGetResponse struct {
	Accounts  []Account `json:"accounts"`
	RequestID string    `json:"request_id"`
}

type institutionGetByIDResponse struct {
	Institution struct {
		Name string `json:"name"`
	} `json:"institution"`
	RequestID string `json:"request_id"`
}

func NewClient(profile state.AppProfile) (*Client, error) {
	baseURL, err := baseURLForEnv(profile.Env)
	if err != nil {
		return nil, err
	}

	return &Client{
		baseURL:  baseURL,
		clientID: profile.ClientID,
		secret:   profile.Secret,
		httpClient: &http.Client{
			Timeout: 45 * time.Second,
		},
	}, nil
}

func (c *Client) CreateHostedLinkToken(ctx context.Context, input CreateHostedLinkTokenInput) (*LinkTokenCreateResponse, error) {
	body := map[string]any{
		"client_name":   input.ClientName,
		"language":      input.Language,
		"country_codes": input.CountryCodes,
		"user": map[string]any{
			"client_user_id": input.ClientUserID,
		},
		"products":    input.Products,
		"hosted_link": map[string]any{},
	}
	if input.Webhook != "" {
		body["webhook"] = input.Webhook
	}
	if input.RedirectURI != "" {
		body["redirect_uri"] = input.RedirectURI
	}

	var out LinkTokenCreateResponse
	if err := c.postJSON(ctx, "/link/token/create", body, &out); err != nil {
		return nil, err
	}
	if out.HostedLinkURL == "" {
		return nil, errors.New("Plaid did not return hosted_link_url")
	}
	return &out, nil
}

func (c *Client) GetLinkToken(ctx context.Context, linkToken string) (map[string]any, error) {
	body := map[string]any{"link_token": linkToken}
	var out map[string]any
	if err := c.postJSON(ctx, "/link/token/get", body, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) ExchangePublicToken(ctx context.Context, publicToken string) (*PublicTokenExchangeResponse, error) {
	body := map[string]any{"public_token": publicToken}
	var out PublicTokenExchangeResponse
	if err := c.postJSON(ctx, "/item/public_token/exchange", body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetItem(ctx context.Context, accessToken string) (*ItemGetResponse, error) {
	body := map[string]any{"access_token": accessToken}
	var out ItemGetResponse
	if err := c.postJSON(ctx, "/item/get", body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetAccounts(ctx context.Context, accessToken string) (*AccountsGetResponse, error) {
	body := map[string]any{"access_token": accessToken}
	var out AccountsGetResponse
	if err := c.postJSON(ctx, "/accounts/get", body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetInstitutionName(ctx context.Context, institutionID string, countryCodes []string) (string, error) {
	body := map[string]any{
		"institution_id": institutionID,
		"country_codes":  countryCodes,
	}
	var out institutionGetByIDResponse
	if err := c.postJSON(ctx, "/institutions/get_by_id", body, &out); err != nil {
		return "", err
	}
	return out.Institution.Name, nil
}

func (c *Client) Call(ctx context.Context, path string, requestBody any) (map[string]any, error) {
	var out map[string]any
	if err := c.postJSON(ctx, path, requestBody, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) CallBytes(ctx context.Context, path string, requestBody any) (*BinaryResponse, error) {
	bodyBytes, headers, err := c.post(ctx, path, requestBody)
	if err != nil {
		return nil, err
	}
	return &BinaryResponse{
		Body:    bodyBytes,
		Headers: headers,
	}, nil
}

func (c *Client) CallMultipart(ctx context.Context, path string, fields map[string]string, fileField, filePath string) (map[string]any, error) {
	var payload bytes.Buffer
	writer := multipart.NewWriter(&payload)

	for key, value := range fields {
		if strings.TrimSpace(value) == "" {
			continue
		}
		if err := writer.WriteField(key, value); err != nil {
			return nil, fmt.Errorf("write multipart field %s: %w", key, err)
		}
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open multipart file: %w", err)
	}
	defer file.Close()

	part, err := writer.CreateFormFile(fileField, filepath.Base(filePath))
	if err != nil {
		return nil, fmt.Errorf("create multipart file field: %w", err)
	}
	if _, err := io.Copy(part, file); err != nil {
		return nil, fmt.Errorf("copy multipart file: %w", err)
	}
	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("close multipart writer: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+path, &payload)
	if err != nil {
		return nil, fmt.Errorf("create multipart request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("PLAID-CLIENT-ID", c.clientID)
	req.Header.Set("PLAID-SECRET", c.secret)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send multipart request: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read multipart response body: %w", err)
	}
	if resp.StatusCode >= http.StatusBadRequest {
		apiErr := &APIError{StatusCode: resp.StatusCode}
		if err := json.Unmarshal(bodyBytes, apiErr); err != nil {
			apiErr.ErrorMessage = string(bodyBytes)
		}
		return nil, apiErr
	}

	var out map[string]any
	if len(bodyBytes) == 0 {
		return map[string]any{}, nil
	}
	if err := json.Unmarshal(bodyBytes, &out); err != nil {
		return nil, fmt.Errorf("decode multipart response body: %w", err)
	}
	return out, nil
}

func (c *Client) postJSON(ctx context.Context, path string, requestBody, responseBody any) error {
	bodyBytes, _, err := c.post(ctx, path, requestBody)
	if err != nil {
		return err
	}

	if responseBody == nil || len(bodyBytes) == 0 {
		return nil
	}
	if err := json.Unmarshal(bodyBytes, responseBody); err != nil {
		return fmt.Errorf("decode response body: %w", err)
	}
	return nil
}

func (c *Client) post(ctx context.Context, path string, requestBody any) ([]byte, http.Header, error) {
	payload, err := json.Marshal(requestBody)
	if err != nil {
		return nil, nil, fmt.Errorf("marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+path, bytes.NewReader(payload))
	if err != nil {
		return nil, nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("PLAID-CLIENT-ID", c.clientID)
	req.Header.Set("PLAID-SECRET", c.secret)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("read response body: %w", err)
	}

	if resp.StatusCode >= http.StatusBadRequest {
		apiErr := &APIError{StatusCode: resp.StatusCode}
		if err := json.Unmarshal(bodyBytes, apiErr); err != nil {
			apiErr.ErrorMessage = string(bodyBytes)
		}
		return nil, nil, apiErr
	}

	return bodyBytes, resp.Header.Clone(), nil
}

func baseURLForEnv(env string) (string, error) {
	switch strings.ToLower(strings.TrimSpace(env)) {
	case "sandbox":
		return "https://sandbox.plaid.com", nil
	case "development":
		return "https://development.plaid.com", nil
	case "production":
		return "https://production.plaid.com", nil
	default:
		return "", fmt.Errorf("unsupported Plaid environment %q", env)
	}
}
