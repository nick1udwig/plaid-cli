package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const oauthDocPath = "docs/plaid/api/oauth/index.md"

func newOAuthCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oauth",
		Short: "OAuth token lifecycle commands",
		Long:  "Admin-style commands for creating, introspecting, and revoking Plaid OAuth tokens.",
	}

	cmd.AddCommand(newOAuthTokenCmd())
	cmd.AddCommand(newOAuthIntrospectCmd())
	cmd.AddCommand(newOAuthRevokeCmd())

	return cmd
}

func newOAuthTokenCmd() *cobra.Command {
	var grantType, scope, refreshToken, resource, audience, subjectToken, subjectTokenType string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "token",
		Short: "Call /oauth/token",
		Long:  "Capability: admin. Creates or refreshes an OAuth access token for Plaid services.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := oauthTokenTemplate(grantType)
			if handled, err := maybeWriteInfo(cmd, info, oauthDocPath, template); handled || err != nil {
				return err
			}

			_, _, client, err := loadClientFromState(cmd)
			if err != nil {
				return err
			}
			body, err := loadRequestBody(bodyFlags.body)
			if err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "grant-type", grantType, "grant_type"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "scope", scope, "scope"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "refresh-token", refreshToken, "refresh_token"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "resource", resource, "resource"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "audience", audience, "audience"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "subject-token", subjectToken, "subject_token"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "subject-token-type", subjectTokenType, "subject_token_type"); err != nil {
				return err
			}

			rawGrantType, ok := bodyValue(body, "grant_type")
			if !ok {
				return fmt.Errorf("--grant-type is required")
			}
			resolvedGrantType, ok := rawGrantType.(string)
			if !ok || resolvedGrantType == "" {
				return fmt.Errorf("request body field grant_type must be a non-empty string")
			}

			switch resolvedGrantType {
			case "client_credentials":
			case "refresh_token":
				if err := requireBodyFields(body, map[string][]string{
					"--refresh-token": {"refresh_token"},
				}); err != nil {
					return err
				}
			case "urn:ietf:params:oauth:grant-type:token-exchange":
				if err := requireBodyFields(body, map[string][]string{
					"--audience":           {"audience"},
					"--subject-token":      {"subject_token"},
					"--subject-token-type": {"subject_token_type"},
				}); err != nil {
					return err
				}
			default:
				return fmt.Errorf("unsupported grant_type %q", resolvedGrantType)
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/oauth/token", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&grantType, "grant-type", "", "OAuth grant type: client_credentials, refresh_token, or urn:ietf:params:oauth:grant-type:token-exchange")
	cmd.Flags().StringVar(&scope, "scope", "", "Space-separated scope string, e.g. 'user:read user:write'")
	cmd.Flags().StringVar(&refreshToken, "refresh-token", "", "OAuth refresh token for refresh_token grant type")
	cmd.Flags().StringVar(&resource, "resource", "", "Optional target resource URI")
	cmd.Flags().StringVar(&audience, "audience", "", "Audience for token exchange flows")
	cmd.Flags().StringVar(&subjectToken, "subject-token", "", "Subject token for token exchange flows")
	cmd.Flags().StringVar(&subjectTokenType, "subject-token-type", "", "Subject token type for token exchange flows")
	return cmd
}

func newOAuthIntrospectCmd() *cobra.Command {
	var token string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "introspect",
		Short: "Call /oauth/introspect",
		Long:  "Capability: admin. Retrieves metadata about an OAuth token.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"token": "<oauth-token>",
			}
			if handled, err := maybeWriteInfo(cmd, info, oauthDocPath, template); handled || err != nil {
				return err
			}

			_, _, client, err := loadClientFromState(cmd)
			if err != nil {
				return err
			}
			body, err := loadRequestBody(bodyFlags.body)
			if err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "token", token, "token"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--token": {"token"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/oauth/introspect", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&token, "token", "", "OAuth token to introspect")
	return cmd
}

func newOAuthRevokeCmd() *cobra.Command {
	var token string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "revoke",
		Short: "Call /oauth/revoke",
		Long:  "Capability: admin. Revokes an OAuth access token or refresh token.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"token": "<oauth-token>",
			}
			if handled, err := maybeWriteInfo(cmd, info, oauthDocPath, template); handled || err != nil {
				return err
			}

			_, _, client, err := loadClientFromState(cmd)
			if err != nil {
				return err
			}
			body, err := loadRequestBody(bodyFlags.body)
			if err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "token", token, "token"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--token": {"token"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/oauth/revoke", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&token, "token", "", "OAuth token to revoke")
	return cmd
}

func oauthTokenTemplate(grantType string) map[string]any {
	switch grantType {
	case "refresh_token":
		return map[string]any{
			"grant_type":    "refresh_token",
			"refresh_token": "<refresh-token>",
		}
	case "urn:ietf:params:oauth:grant-type:token-exchange":
		return map[string]any{
			"grant_type":         "urn:ietf:params:oauth:grant-type:token-exchange",
			"audience":           "<audience>",
			"subject_token":      "<subject-token>",
			"subject_token_type": "urn:plaid:params:oauth:user-token",
		}
	default:
		return map[string]any{
			"grant_type": "client_credentials",
		}
	}
}
