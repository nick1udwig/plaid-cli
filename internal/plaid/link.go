package plaid

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var errSessionFinishedWithoutPublicToken = errors.New("link session finished without a public token")

func WaitForPublicToken(ctx context.Context, client *Client, linkToken string, pollInterval time.Duration) (string, error) {
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for {
		payload, err := client.GetLinkToken(ctx, linkToken)
		if err != nil {
			return "", err
		}

		if token, ok := extractPublicToken(payload); ok {
			return token, nil
		}
		if sessionFinished(payload) {
			return "", errSessionFinishedWithoutPublicToken
		}

		select {
		case <-ctx.Done():
			return "", fmt.Errorf("timed out waiting for Link completion: %w", ctx.Err())
		case <-ticker.C:
		}
	}
}

func extractPublicToken(value any) (string, bool) {
	switch typed := value.(type) {
	case map[string]any:
		if raw, ok := typed["public_token"].(string); ok && raw != "" {
			return raw, true
		}
		if rawList, ok := typed["public_tokens"].([]any); ok {
			for _, item := range rawList {
				if token, ok := item.(string); ok && token != "" {
					return token, true
				}
			}
		}
		for _, child := range typed {
			if token, ok := extractPublicToken(child); ok {
				return token, true
			}
		}
	case []any:
		for _, child := range typed {
			if token, ok := extractPublicToken(child); ok {
				return token, true
			}
		}
	}
	return "", false
}

func sessionFinished(value any) bool {
	switch typed := value.(type) {
	case map[string]any:
		if raw, ok := typed["finished_at"].(string); ok && raw != "" {
			return true
		}
		for _, child := range typed {
			if sessionFinished(child) {
				return true
			}
		}
	case []any:
		for _, child := range typed {
			if sessionFinished(child) {
				return true
			}
		}
	}
	return false
}
