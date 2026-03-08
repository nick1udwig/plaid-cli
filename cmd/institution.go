package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

func newInstitutionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "institution",
		Short: "Search and inspect supported institutions",
		Long:  "Read-only institution discovery commands.",
	}
	cmd.AddCommand(newInstitutionGetCmd())
	cmd.AddCommand(newInstitutionGetByIDCmd())
	cmd.AddCommand(newInstitutionSearchCmd())
	return cmd
}

func newInstitutionGetCmd() *cobra.Command {
	var count, offset int
	var countryCodes []string
	info := bindInfoFlags(&cobra.Command{})

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Call /institutions/get",
		Long:  "Capability: read. Retrieves a paginated institution list.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			store, profile, client, err := loadClientFromState(cmd)
			_ = store
			if err != nil {
				return err
			}
			countryCodes = defaultCountryCodes(profile, countryCodes)

			template := map[string]any{
				"count":         count,
				"offset":        offset,
				"country_codes": countryCodes,
			}
			if handled, err := maybeWriteInfo(cmd, info, "docs/plaid/api/institutions/index.md", template); handled || err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/institutions/get", template)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	cmd.Flags().IntVar(&count, "count", 20, "Maximum institutions to return")
	cmd.Flags().IntVar(&offset, "offset", 0, "Offset into institution list")
	cmd.Flags().StringSliceVar(&countryCodes, "country-code", nil, "Country code filter (repeatable)")
	return cmd
}

func newInstitutionGetByIDCmd() *cobra.Command {
	var institutionID string
	var countryCodes []string
	info := bindInfoFlags(&cobra.Command{})

	cmd := &cobra.Command{
		Use:   "get-by-id",
		Short: "Call /institutions/get_by_id",
		Long:  "Capability: read. Retrieves a single institution by institution_id.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if institutionID == "" {
				return errors.New("--institution-id is required")
			}

			store, profile, client, err := loadClientFromState(cmd)
			_ = store
			if err != nil {
				return err
			}
			countryCodes = defaultCountryCodes(profile, countryCodes)

			body := map[string]any{
				"institution_id": institutionID,
				"country_codes":  countryCodes,
			}
			if handled, err := maybeWriteInfo(cmd, info, "docs/plaid/api/institutions/index.md", body); handled || err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/institutions/get_by_id", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	cmd.Flags().StringVar(&institutionID, "institution-id", "", "Plaid institution_id")
	cmd.Flags().StringSliceVar(&countryCodes, "country-code", nil, "Country code filter (repeatable)")
	return cmd
}

func newInstitutionSearchCmd() *cobra.Command {
	var query string
	var products []string
	var countryCodes []string
	info := bindInfoFlags(&cobra.Command{})

	cmd := &cobra.Command{
		Use:   "search",
		Short: "Call /institutions/search",
		Long:  "Capability: read. Searches institutions by query text and optional product filters.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if query == "" {
				return errors.New("--query is required")
			}

			store, profile, client, err := loadClientFromState(cmd)
			_ = store
			if err != nil {
				return err
			}
			countryCodes = defaultCountryCodes(profile, countryCodes)

			body := map[string]any{
				"query":         query,
				"country_codes": countryCodes,
			}
			if len(products) > 0 {
				body["products"] = products
			}
			if handled, err := maybeWriteInfo(cmd, info, "docs/plaid/api/institutions/index.md", body); handled || err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/institutions/search", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	cmd.Flags().StringVar(&query, "query", "", "Institution search string")
	cmd.Flags().StringSliceVar(&products, "product", nil, "Filter institutions by supported product (repeatable)")
	cmd.Flags().StringSliceVar(&countryCodes, "country-code", nil, "Country code filter (repeatable)")
	return cmd
}
