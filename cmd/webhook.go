package cmd

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	"github.com/spf13/cobra"
)

const webhookVerificationDocPath = "docs/plaid/api/webhooks/webhook-verification/index.md"

func newWebhookCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "webhook",
		Short: "Webhook utility commands",
		Long:  "Commands for working with Plaid webhook verification and related operational flows.",
	}

	cmd.AddCommand(newWebhookVerificationKeyCmd())

	return cmd
}

func newWebhookVerificationKeyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verification-key",
		Short: "Webhook verification key commands",
		Long:  "Retrieve Plaid webhook verification keys by key ID or signed webhook header.",
	}

	cmd.AddCommand(newWebhookVerificationKeyGetCmd())

	return cmd
}

func extractWebhookKeyID(signedJWT string) (string, error) {
	signedJWT = strings.TrimSpace(signedJWT)
	if signedJWT == "" {
		return "", errors.New("signed JWT must not be empty")
	}

	parts := strings.Split(signedJWT, ".")
	if len(parts) < 2 {
		return "", errors.New("signed JWT must contain at least two dot-separated segments")
	}

	headerBytes, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return "", err
	}

	var header struct {
		KeyID string `json:"kid"`
	}
	if err := json.Unmarshal(headerBytes, &header); err != nil {
		return "", err
	}
	if strings.TrimSpace(header.KeyID) == "" {
		return "", errors.New("signed JWT header does not contain a kid")
	}
	return header.KeyID, nil
}

func newWebhookVerificationKeyGetCmd() *cobra.Command {
	var keyID, signedJWT string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /webhook_verification_key/get",
		Long:  "Capability: read. Retrieves the JWK used to verify Plaid webhooks.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"key_id": "<kid>"}
			if handled, err := maybeWriteInfo(cmd, info, webhookVerificationDocPath, template); handled || err != nil {
				return err
			}
			if cmd.Flags().Changed("key-id") && cmd.Flags().Changed("plaid-verification") {
				return errors.New("provide only one of --key-id or --plaid-verification")
			}

			_, _, client, err := loadClientFromState(cmd)
			if err != nil {
				return err
			}
			body, err := loadRequestBody(bodyFlags.body)
			if err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "key-id", keyID, "key_id"); err != nil {
				return err
			}
			if cmd.Flags().Changed("plaid-verification") {
				resolvedKeyID, err := extractWebhookKeyID(signedJWT)
				if err != nil {
					return err
				}
				if err := setBodyValue(body, resolvedKeyID, "key_id"); err != nil {
					return err
				}
			}
			if err := requireBodyFields(body, map[string][]string{
				"--key-id or --plaid-verification": {"key_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/webhook_verification_key/get", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&keyID, "key-id", "", "Key ID from the Plaid-Verification JWT header")
	cmd.Flags().StringVar(&signedJWT, "plaid-verification", "", "Full Plaid-Verification JWT; the key ID will be extracted automatically")
	return cmd
}
