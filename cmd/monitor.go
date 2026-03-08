package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

const monitorDocPath = "docs/plaid/api/products/monitor/index.md"

type monitorIndividualSearchFlags struct {
	programID      string
	legalName      string
	dateOfBirth    string
	documentNumber string
	country        string
}

type monitorEntitySearchFlags struct {
	programID      string
	legalName      string
	documentNumber string
	emailAddress   string
	country        string
	phoneNumber    string
	url            string
}

type monitorUserRefFlags struct {
	clientUserID string
	userID       string
}

func newMonitorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "monitor",
		Short: "Monitor product commands",
		Long:  "Monitor watchlist screening commands for individuals and entities.",
	}

	cmd.AddCommand(newMonitorIndividualCmd())
	cmd.AddCommand(newMonitorEntityCmd())

	return cmd
}

func newMonitorIndividualCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "individual",
		Short: "Individual watchlist screening commands",
		Long:  "Create, update, inspect, and review individual watchlist screenings.",
	}

	cmd.AddCommand(newMonitorIndividualCreateCmd())
	cmd.AddCommand(newMonitorIndividualGetCmd())
	cmd.AddCommand(newMonitorIndividualListCmd())
	cmd.AddCommand(newMonitorIndividualUpdateCmd())
	cmd.AddCommand(newMonitorIndividualHistoryListCmd())
	cmd.AddCommand(newMonitorIndividualReviewCmd())
	cmd.AddCommand(newMonitorIndividualHitCmd())
	cmd.AddCommand(newMonitorIndividualProgramCmd())

	return cmd
}

func newMonitorEntityCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "entity",
		Short: "Entity watchlist screening commands",
		Long:  "Create, update, inspect, and review entity watchlist screenings.",
	}

	cmd.AddCommand(newMonitorEntityCreateCmd())
	cmd.AddCommand(newMonitorEntityGetCmd())
	cmd.AddCommand(newMonitorEntityListCmd())
	cmd.AddCommand(newMonitorEntityUpdateCmd())
	cmd.AddCommand(newMonitorEntityHistoryListCmd())
	cmd.AddCommand(newMonitorEntityReviewCmd())
	cmd.AddCommand(newMonitorEntityHitCmd())
	cmd.AddCommand(newMonitorEntityProgramCmd())

	return cmd
}

func bindMonitorUserRefFlags(cmd *cobra.Command) *monitorUserRefFlags {
	flags := &monitorUserRefFlags{}
	cmd.Flags().StringVar(&flags.clientUserID, "client-user-id", "", "Stable client-side identifier for the monitored subject")
	cmd.Flags().StringVar(&flags.userID, "user-id", "", "Plaid user_id for the monitored subject")
	return flags
}

func applyMonitorUserRefFlags(cmd *cobra.Command, body map[string]any, flags *monitorUserRefFlags) error {
	if flags == nil {
		return nil
	}
	if err := applyStringFlag(cmd, body, "client-user-id", flags.clientUserID, "client_user_id"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "user-id", flags.userID, "user_id"); err != nil {
		return err
	}
	return nil
}

func requireMonitorUserRef(body map[string]any) error {
	return requireAtLeastOneBodyField(body, map[string][]string{
		"--client-user-id": {"client_user_id"},
		"--user-id":        {"user_id"},
	})
}

func bindMonitorIndividualSearchFlags(cmd *cobra.Command, prefix string) *monitorIndividualSearchFlags {
	flags := &monitorIndividualSearchFlags{}
	cmd.Flags().StringVar(&flags.programID, prefix+"-program-id", "", "Watchlist program ID")
	cmd.Flags().StringVar(&flags.legalName, prefix+"-legal-name", "", "Legal name")
	cmd.Flags().StringVar(&flags.dateOfBirth, prefix+"-date-of-birth", "", "Date of birth in YYYY-MM-DD")
	cmd.Flags().StringVar(&flags.documentNumber, prefix+"-document-number", "", "Document number")
	cmd.Flags().StringVar(&flags.country, prefix+"-country", "", "ISO country code")
	return flags
}

func bindMonitorEntitySearchFlags(cmd *cobra.Command, prefix string) *monitorEntitySearchFlags {
	flags := &monitorEntitySearchFlags{}
	cmd.Flags().StringVar(&flags.programID, prefix+"-program-id", "", "Entity watchlist program ID")
	cmd.Flags().StringVar(&flags.legalName, prefix+"-legal-name", "", "Entity legal name")
	cmd.Flags().StringVar(&flags.documentNumber, prefix+"-document-number", "", "Entity document number")
	cmd.Flags().StringVar(&flags.emailAddress, prefix+"-email-address", "", "Entity email address")
	cmd.Flags().StringVar(&flags.country, prefix+"-country", "", "Entity ISO country code")
	cmd.Flags().StringVar(&flags.phoneNumber, prefix+"-phone-number", "", "Entity phone number")
	cmd.Flags().StringVar(&flags.url, prefix+"-url", "", "Entity URL")
	return flags
}

func applyMonitorIndividualSearchFlags(cmd *cobra.Command, body map[string]any, flags *monitorIndividualSearchFlags) error {
	if flags == nil {
		return nil
	}
	if err := applyStringFlag(cmd, body, "search-program-id", flags.programID, "search_terms", "watchlist_program_id"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "search-legal-name", flags.legalName, "search_terms", "legal_name"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "search-date-of-birth", flags.dateOfBirth, "search_terms", "date_of_birth"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "search-document-number", flags.documentNumber, "search_terms", "document_number"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "search-country", flags.country, "search_terms", "country"); err != nil {
		return err
	}
	return nil
}

func applyMonitorEntitySearchFlags(cmd *cobra.Command, body map[string]any, flags *monitorEntitySearchFlags) error {
	if flags == nil {
		return nil
	}
	if err := applyStringFlag(cmd, body, "search-program-id", flags.programID, "search_terms", "entity_watchlist_program_id"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "search-legal-name", flags.legalName, "search_terms", "legal_name"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "search-document-number", flags.documentNumber, "search_terms", "document_number"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "search-email-address", flags.emailAddress, "search_terms", "email_address"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "search-country", flags.country, "search_terms", "country"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "search-phone-number", flags.phoneNumber, "search_terms", "phone_number"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "search-url", flags.url, "search_terms", "url"); err != nil {
		return err
	}
	return nil
}

func monitorSearchTermsPresent(body map[string]any) bool {
	return bodyHasValue(body, "search_terms")
}

func validateMonitorIndividualCreate(body map[string]any) error {
	if !bodyHasValue(body, "search_terms", "watchlist_program_id") {
		return errors.New("search_terms.watchlist_program_id is required")
	}
	if !bodyHasValue(body, "search_terms", "legal_name") {
		return errors.New("search_terms.legal_name is required")
	}
	return requireMonitorUserRef(body)
}

func validateMonitorEntityCreate(body map[string]any) error {
	if !bodyHasValue(body, "search_terms", "entity_watchlist_program_id") {
		return errors.New("search_terms.entity_watchlist_program_id is required")
	}
	if !bodyHasValue(body, "search_terms", "legal_name") {
		return errors.New("search_terms.legal_name is required")
	}
	return requireMonitorUserRef(body)
}

func validateMonitorIndividualUpdate(body map[string]any) error {
	if !monitorSearchTermsPresent(body) && !bodyHasValue(body, "status") && !bodyHasValue(body, "assignee") &&
		!bodyHasValue(body, "client_user_id") && !bodyHasValue(body, "user_id") {
		return errors.New("provide at least one update field")
	}
	if monitorSearchTermsPresent(body) && bodyHasValue(body, "status") {
		return errors.New("search_terms and status cannot be updated in the same request")
	}
	return nil
}

func validateMonitorEntityUpdate(body map[string]any) error {
	if !monitorSearchTermsPresent(body) && !bodyHasValue(body, "status") && !bodyHasValue(body, "assignee") &&
		!bodyHasValue(body, "client_user_id") && !bodyHasValue(body, "user_id") {
		return errors.New("provide at least one update field")
	}
	if monitorSearchTermsPresent(body) && bodyHasValue(body, "status") {
		return errors.New("search_terms and status cannot be updated in the same request")
	}
	if monitorSearchTermsPresent(body) && !bodyHasValue(body, "search_terms", "entity_watchlist_program_id") {
		return errors.New("search_terms.entity_watchlist_program_id is required when updating entity search terms")
	}
	return nil
}

func applyMonitorListFilters(cmd *cobra.Command, body map[string]any, programFlag, programValue, status, assignee, cursor string) error {
	if err := applyStringFlag(cmd, body, programFlag, programValue, mapProgramPath(programFlag)...); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "status", status, "status"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "assignee", assignee, "assignee"); err != nil {
		return err
	}
	if err := applyStringFlag(cmd, body, "cursor", cursor, "cursor"); err != nil {
		return err
	}
	return nil
}

func mapProgramPath(flag string) []string {
	if flag == "entity-program-id" {
		return []string{"entity_watchlist_program_id"}
	}
	return []string{"watchlist_program_id"}
}

func applyMonitorReviewFields(cmd *cobra.Command, body map[string]any, confirmedHits, dismissedHits []string, comment string) error {
	if !bodyHasValue(body, "confirmed_hits") {
		if err := setBodyValue(body, []string{}, "confirmed_hits"); err != nil {
			return err
		}
	}
	if !bodyHasValue(body, "dismissed_hits") {
		if err := setBodyValue(body, []string{}, "dismissed_hits"); err != nil {
			return err
		}
	}
	if err := applyStringSliceFlag(cmd, body, "confirmed-hit", confirmedHits, "confirmed_hits"); err != nil {
		return err
	}
	if err := applyStringSliceFlag(cmd, body, "dismissed-hit", dismissedHits, "dismissed_hits"); err != nil {
		return err
	}
	return applyStringFlag(cmd, body, "comment", comment, "comment")
}

func applyMonitorPagingID(cmd *cobra.Command, body map[string]any, idFlag, idValue, idField, cursor string) error {
	if err := applyStringFlag(cmd, body, idFlag, idValue, idField); err != nil {
		return err
	}
	return applyStringFlag(cmd, body, "cursor", cursor, "cursor")
}

func newMonitorSimpleGetCmd(use, short, long, endpoint, idFlag, idField string) *cobra.Command {
	var id string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{idField: "<id>"}
			if handled, err := maybeWriteInfo(cmd, info, monitorDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, idFlag, id, idField); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--" + idFlag: {idField},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, endpoint, body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&id, idFlag, "", "Plaid identifier for this resource")
	return cmd
}

func newMonitorCursorListCmd(use, short, long, endpoint, idFlag, idField string) *cobra.Command {
	var id, cursor string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{idField: "<id>"}
			if cursor != "" {
				template["cursor"] = cursor
			}
			if handled, err := maybeWriteInfo(cmd, info, monitorDocPath, template); handled || err != nil {
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
			if err := applyMonitorPagingID(cmd, body, idFlag, id, idField, cursor); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--" + idFlag: {idField},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, endpoint, body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&id, idFlag, "", "Plaid identifier for this resource")
	cmd.Flags().StringVar(&cursor, "cursor", "", "Pagination cursor")
	return cmd
}

func newMonitorProgramGetCmd(use, short, long, endpoint, flagName, field string) *cobra.Command {
	return newMonitorSimpleGetCmd(use, short, long, endpoint, flagName, field)
}

func newMonitorProgramListCmd(use, short, long, endpoint string) *cobra.Command {
	var cursor string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{}
			if cursor != "" {
				template["cursor"] = cursor
			}
			if handled, err := maybeWriteInfo(cmd, info, monitorDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "cursor", cursor, "cursor"); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, endpoint, body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&cursor, "cursor", "", "Pagination cursor")
	return cmd
}

func newMonitorIndividualCreateCmd() *cobra.Command {
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags
	var searchFlags *monitorIndividualSearchFlags
	var userRefFlags *monitorUserRefFlags

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Call /watchlist_screening/individual/create",
		Long:  "Capability: write. Creates a watchlist screening for an individual.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"search_terms": map[string]any{
					"watchlist_program_id": "<program-id>",
					"legal_name":           "Jane Doe",
				},
				"client_user_id": "<client-user-id>",
			}
			if handled, err := maybeWriteInfo(cmd, info, monitorDocPath, template); handled || err != nil {
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
			if err := applyMonitorIndividualSearchFlags(cmd, body, searchFlags); err != nil {
				return err
			}
			if err := applyMonitorUserRefFlags(cmd, body, userRefFlags); err != nil {
				return err
			}
			if err := validateMonitorIndividualCreate(body); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/watchlist_screening/individual/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	searchFlags = bindMonitorIndividualSearchFlags(cmd, "search")
	userRefFlags = bindMonitorUserRefFlags(cmd)
	return cmd
}

func newMonitorIndividualGetCmd() *cobra.Command {
	return newMonitorSimpleGetCmd(
		"get",
		"Call /watchlist_screening/individual/get",
		"Capability: read. Retrieves an individual watchlist screening.",
		"/watchlist_screening/individual/get",
		"watchlist-screening-id",
		"watchlist_screening_id",
	)
}

func newMonitorIndividualListCmd() *cobra.Command {
	var programID, status, assignee, cursor string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags
	var userRefFlags *monitorUserRefFlags

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Call /watchlist_screening/individual/list",
		Long:  "Capability: read. Lists individual watchlist screenings for a program.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"watchlist_program_id": "<program-id>"}
			if handled, err := maybeWriteInfo(cmd, info, monitorDocPath, template); handled || err != nil {
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
			if err := applyMonitorListFilters(cmd, body, "individual-program-id", programID, status, assignee, cursor); err != nil {
				return err
			}
			if err := applyMonitorUserRefFlags(cmd, body, userRefFlags); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--individual-program-id": {"watchlist_program_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/watchlist_screening/individual/list", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	userRefFlags = bindMonitorUserRefFlags(cmd)
	cmd.Flags().StringVar(&programID, "individual-program-id", "", "Individual watchlist program ID")
	cmd.Flags().StringVar(&status, "status", "", "Filter by screening status")
	cmd.Flags().StringVar(&assignee, "assignee", "", "Filter by assignee dashboard user ID")
	cmd.Flags().StringVar(&cursor, "cursor", "", "Pagination cursor")
	return cmd
}

func newMonitorIndividualUpdateCmd() *cobra.Command {
	var screeningID, assignee, status string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags
	var searchFlags *monitorIndividualSearchFlags
	var userRefFlags *monitorUserRefFlags

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Call /watchlist_screening/individual/update",
		Long:  "Capability: write. Updates an individual watchlist screening.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"watchlist_screening_id": "<screening-id>"}
			if handled, err := maybeWriteInfo(cmd, info, monitorDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "watchlist-screening-id", screeningID, "watchlist_screening_id"); err != nil {
				return err
			}
			if err := applyMonitorIndividualSearchFlags(cmd, body, searchFlags); err != nil {
				return err
			}
			if err := applyMonitorUserRefFlags(cmd, body, userRefFlags); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "assignee", assignee, "assignee"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "status", status, "status"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--watchlist-screening-id": {"watchlist_screening_id"},
			}); err != nil {
				return err
			}
			if err := validateMonitorIndividualUpdate(body); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/watchlist_screening/individual/update", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	searchFlags = bindMonitorIndividualSearchFlags(cmd, "search")
	userRefFlags = bindMonitorUserRefFlags(cmd)
	cmd.Flags().StringVar(&screeningID, "watchlist-screening-id", "", "Individual watchlist screening ID")
	cmd.Flags().StringVar(&assignee, "assignee", "", "Dashboard user ID to assign the screening to")
	cmd.Flags().StringVar(&status, "status", "", "Set screening status, e.g. pending_review, rejected, or cleared")
	return cmd
}

func newMonitorIndividualHistoryListCmd() *cobra.Command {
	return newMonitorCursorListCmd(
		"history-list",
		"Call /watchlist_screening/individual/history/list",
		"Capability: read. Lists change history for an individual watchlist screening.",
		"/watchlist_screening/individual/history/list",
		"watchlist-screening-id",
		"watchlist_screening_id",
	)
}

func newMonitorIndividualReviewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "review",
		Short: "Individual screening review commands",
		Long:  "Create and list reviews for an individual watchlist screening.",
	}
	cmd.AddCommand(newMonitorIndividualReviewCreateCmd())
	cmd.AddCommand(newMonitorIndividualReviewListCmd())
	return cmd
}

func newMonitorIndividualReviewCreateCmd() *cobra.Command {
	var screeningID, comment string
	var confirmedHits, dismissedHits []string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Call /watchlist_screening/individual/review/create",
		Long:  "Capability: write. Creates a review for an individual watchlist screening.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"watchlist_screening_id": "<screening-id>"}
			if handled, err := maybeWriteInfo(cmd, info, monitorDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "watchlist-screening-id", screeningID, "watchlist_screening_id"); err != nil {
				return err
			}
			if err := applyMonitorReviewFields(cmd, body, confirmedHits, dismissedHits, comment); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--watchlist-screening-id": {"watchlist_screening_id"},
			}); err != nil {
				return err
			}
			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/watchlist_screening/individual/review/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}
	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&screeningID, "watchlist-screening-id", "", "Individual watchlist screening ID")
	cmd.Flags().StringSliceVar(&confirmedHits, "confirmed-hit", nil, "Confirmed hit ID (repeatable)")
	cmd.Flags().StringSliceVar(&dismissedHits, "dismissed-hit", nil, "Dismissed hit ID (repeatable)")
	cmd.Flags().StringVar(&comment, "comment", "", "Review comment")
	return cmd
}

func newMonitorIndividualReviewListCmd() *cobra.Command {
	return newMonitorCursorListCmd(
		"list",
		"Call /watchlist_screening/individual/review/list",
		"Capability: read. Lists reviews for an individual watchlist screening.",
		"/watchlist_screening/individual/review/list",
		"watchlist-screening-id",
		"watchlist_screening_id",
	)
}

func newMonitorIndividualHitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hit",
		Short: "Individual screening hit commands",
		Long:  "List hits for an individual watchlist screening.",
	}
	cmd.AddCommand(newMonitorCursorListCmd(
		"list",
		"Call /watchlist_screening/individual/hit/list",
		"Capability: read. Lists hits for an individual watchlist screening.",
		"/watchlist_screening/individual/hit/list",
		"watchlist-screening-id",
		"watchlist_screening_id",
	))
	return cmd
}

func newMonitorIndividualProgramCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "program",
		Short: "Individual watchlist program commands",
		Long:  "Inspect individual watchlist programs.",
	}
	cmd.AddCommand(newMonitorProgramGetCmd(
		"get",
		"Call /watchlist_screening/individual/program/get",
		"Capability: read. Retrieves an individual watchlist program.",
		"/watchlist_screening/individual/program/get",
		"watchlist-program-id",
		"watchlist_program_id",
	))
	cmd.AddCommand(newMonitorProgramListCmd(
		"list",
		"Call /watchlist_screening/individual/program/list",
		"Capability: read. Lists individual watchlist programs.",
		"/watchlist_screening/individual/program/list",
	))
	return cmd
}

func newMonitorEntityCreateCmd() *cobra.Command {
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags
	var searchFlags *monitorEntitySearchFlags
	var userRefFlags *monitorUserRefFlags

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Call /watchlist_screening/entity/create",
		Long:  "Capability: write. Creates a watchlist screening for an entity.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{
				"search_terms": map[string]any{
					"entity_watchlist_program_id": "<entity-program-id>",
					"legal_name":                  "Example LLC",
				},
				"client_user_id": "<client-user-id>",
			}
			if handled, err := maybeWriteInfo(cmd, info, monitorDocPath, template); handled || err != nil {
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
			if err := applyMonitorEntitySearchFlags(cmd, body, searchFlags); err != nil {
				return err
			}
			if err := applyMonitorUserRefFlags(cmd, body, userRefFlags); err != nil {
				return err
			}
			if err := validateMonitorEntityCreate(body); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/watchlist_screening/entity/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	searchFlags = bindMonitorEntitySearchFlags(cmd, "search")
	userRefFlags = bindMonitorUserRefFlags(cmd)
	return cmd
}

func newMonitorEntityGetCmd() *cobra.Command {
	return newMonitorSimpleGetCmd(
		"get",
		"Call /watchlist_screening/entity/get",
		"Capability: read. Retrieves an entity watchlist screening.",
		"/watchlist_screening/entity/get",
		"entity-watchlist-screening-id",
		"entity_watchlist_screening_id",
	)
}

func newMonitorEntityListCmd() *cobra.Command {
	var programID, status, assignee, cursor string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags
	var userRefFlags *monitorUserRefFlags

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Call /watchlist_screening/entity/list",
		Long:  "Capability: read. Lists entity watchlist screenings for a program.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"entity_watchlist_program_id": "<entity-program-id>"}
			if handled, err := maybeWriteInfo(cmd, info, monitorDocPath, template); handled || err != nil {
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
			if err := applyMonitorListFilters(cmd, body, "entity-program-id", programID, status, assignee, cursor); err != nil {
				return err
			}
			if err := applyMonitorUserRefFlags(cmd, body, userRefFlags); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--entity-program-id": {"entity_watchlist_program_id"},
			}); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/watchlist_screening/entity/list", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	userRefFlags = bindMonitorUserRefFlags(cmd)
	cmd.Flags().StringVar(&programID, "entity-program-id", "", "Entity watchlist program ID")
	cmd.Flags().StringVar(&status, "status", "", "Filter by screening status")
	cmd.Flags().StringVar(&assignee, "assignee", "", "Filter by assignee dashboard user ID")
	cmd.Flags().StringVar(&cursor, "cursor", "", "Pagination cursor")
	return cmd
}

func newMonitorEntityUpdateCmd() *cobra.Command {
	var screeningID, assignee, status string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags
	var searchFlags *monitorEntitySearchFlags
	var userRefFlags *monitorUserRefFlags

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Call /watchlist_screening/entity/update",
		Long:  "Capability: write. Updates an entity watchlist screening.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"entity_watchlist_screening_id": "<entity-screening-id>"}
			if handled, err := maybeWriteInfo(cmd, info, monitorDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "entity-watchlist-screening-id", screeningID, "entity_watchlist_screening_id"); err != nil {
				return err
			}
			if err := applyMonitorEntitySearchFlags(cmd, body, searchFlags); err != nil {
				return err
			}
			if err := applyMonitorUserRefFlags(cmd, body, userRefFlags); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "assignee", assignee, "assignee"); err != nil {
				return err
			}
			if err := applyStringFlag(cmd, body, "status", status, "status"); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--entity-watchlist-screening-id": {"entity_watchlist_screening_id"},
			}); err != nil {
				return err
			}
			if err := validateMonitorEntityUpdate(body); err != nil {
				return err
			}

			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/watchlist_screening/entity/update", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}

	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	searchFlags = bindMonitorEntitySearchFlags(cmd, "search")
	userRefFlags = bindMonitorUserRefFlags(cmd)
	cmd.Flags().StringVar(&screeningID, "entity-watchlist-screening-id", "", "Entity watchlist screening ID")
	cmd.Flags().StringVar(&assignee, "assignee", "", "Dashboard user ID to assign the screening to")
	cmd.Flags().StringVar(&status, "status", "", "Set screening status, e.g. pending_review, rejected, or cleared")
	return cmd
}

func newMonitorEntityHistoryListCmd() *cobra.Command {
	return newMonitorCursorListCmd(
		"history-list",
		"Call /watchlist_screening/entity/history/list",
		"Capability: read. Lists change history for an entity watchlist screening.",
		"/watchlist_screening/entity/history/list",
		"entity-watchlist-screening-id",
		"entity_watchlist_screening_id",
	)
}

func newMonitorEntityReviewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "review",
		Short: "Entity screening review commands",
		Long:  "Create and list reviews for an entity watchlist screening.",
	}
	cmd.AddCommand(newMonitorEntityReviewCreateCmd())
	cmd.AddCommand(newMonitorEntityReviewListCmd())
	return cmd
}

func newMonitorEntityReviewCreateCmd() *cobra.Command {
	var screeningID, comment string
	var confirmedHits, dismissedHits []string
	var info *commandInfoFlags
	var bodyFlags *requestBodyFlags

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Call /watchlist_screening/entity/review/create",
		Long:  "Capability: write. Creates a review for an entity watchlist screening.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			template := map[string]any{"entity_watchlist_screening_id": "<entity-screening-id>"}
			if handled, err := maybeWriteInfo(cmd, info, monitorDocPath, template); handled || err != nil {
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
			if err := applyStringFlag(cmd, body, "entity-watchlist-screening-id", screeningID, "entity_watchlist_screening_id"); err != nil {
				return err
			}
			if err := applyMonitorReviewFields(cmd, body, confirmedHits, dismissedHits, comment); err != nil {
				return err
			}
			if err := requireBodyFields(body, map[string][]string{
				"--entity-watchlist-screening-id": {"entity_watchlist_screening_id"},
			}); err != nil {
				return err
			}
			ctx, cancel := commandContext()
			defer cancel()
			resp, err := client.Call(ctx, "/watchlist_screening/entity/review/create", body)
			if err != nil {
				return err
			}
			return writeJSON(cmd, resp)
		},
	}
	info = bindInfoFlags(cmd)
	bodyFlags = bindBodyFlag(cmd)
	cmd.Flags().StringVar(&screeningID, "entity-watchlist-screening-id", "", "Entity watchlist screening ID")
	cmd.Flags().StringSliceVar(&confirmedHits, "confirmed-hit", nil, "Confirmed hit ID (repeatable)")
	cmd.Flags().StringSliceVar(&dismissedHits, "dismissed-hit", nil, "Dismissed hit ID (repeatable)")
	cmd.Flags().StringVar(&comment, "comment", "", "Review comment")
	return cmd
}

func newMonitorEntityReviewListCmd() *cobra.Command {
	return newMonitorCursorListCmd(
		"list",
		"Call /watchlist_screening/entity/review/list",
		"Capability: read. Lists reviews for an entity watchlist screening.",
		"/watchlist_screening/entity/review/list",
		"entity-watchlist-screening-id",
		"entity_watchlist_screening_id",
	)
}

func newMonitorEntityHitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hit",
		Short: "Entity screening hit commands",
		Long:  "List hits for an entity watchlist screening.",
	}
	cmd.AddCommand(newMonitorCursorListCmd(
		"list",
		"Call /watchlist_screening/entity/hit/list",
		"Capability: read. Lists hits for an entity watchlist screening.",
		"/watchlist_screening/entity/hit/list",
		"entity-watchlist-screening-id",
		"entity_watchlist_screening_id",
	))
	return cmd
}

func newMonitorEntityProgramCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "program",
		Short: "Entity watchlist program commands",
		Long:  "Inspect entity watchlist programs.",
	}
	cmd.AddCommand(newMonitorProgramGetCmd(
		"get",
		"Call /watchlist_screening/entity/program/get",
		"Capability: read. Retrieves an entity watchlist program.",
		"/watchlist_screening/entity/program/get",
		"entity-watchlist-program-id",
		"entity_watchlist_program_id",
	))
	cmd.AddCommand(newMonitorProgramListCmd(
		"list",
		"Call /watchlist_screening/entity/program/list",
		"Capability: read. Lists entity watchlist programs.",
		"/watchlist_screening/entity/program/list",
	))
	return cmd
}
