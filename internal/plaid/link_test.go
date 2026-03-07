package plaid

import "testing"

func TestExtractPublicToken(t *testing.T) {
	t.Parallel()

	payload := map[string]any{
		"link_sessions": []any{
			map[string]any{
				"results": map[string]any{
					"item_add_results": []any{
						map[string]any{
							"public_token": "public-sandbox-123",
						},
					},
				},
			},
		},
	}

	token, ok := extractPublicToken(payload)
	if !ok {
		t.Fatal("extractPublicToken() did not find a token")
	}
	if token != "public-sandbox-123" {
		t.Fatalf("token = %q, want %q", token, "public-sandbox-123")
	}
}

func TestSessionFinished(t *testing.T) {
	t.Parallel()

	payload := map[string]any{
		"link_sessions": []any{
			map[string]any{
				"finished_at": "2026-03-07T00:00:00Z",
			},
		},
	}

	if !sessionFinished(payload) {
		t.Fatal("sessionFinished() = false, want true")
	}
}
