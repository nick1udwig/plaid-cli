package cmd

import "testing"

func TestExtractWebhookKeyID(t *testing.T) {
	t.Parallel()

	t.Run("extracts kid from jwt header", func(t *testing.T) {
		t.Parallel()

		jwt := "eyJhbGciOiJFUzI1NiIsImtpZCI6ImJmYmQ1MTExLThlMzMtNDY0My04Y2VkLWIyZTY0MmE3MmYzYyIsInR5cCI6IkpXVCJ9.payload.signature"
		got, err := extractWebhookKeyID(jwt)
		if err != nil {
			t.Fatalf("extractWebhookKeyID() error = %v", err)
		}
		want := "bfbd5111-8e33-4643-8ced-b2e642a72f3c"
		if got != want {
			t.Fatalf("extractWebhookKeyID() = %q, want %q", got, want)
		}
	})

	t.Run("rejects invalid jwt", func(t *testing.T) {
		t.Parallel()

		if _, err := extractWebhookKeyID("not-a-jwt"); err == nil {
			t.Fatal("extractWebhookKeyID() error = nil, want error")
		}
	})
}
