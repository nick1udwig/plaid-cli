package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func applyPIAddressFlags(cmd *cobra.Command, body map[string]any, street []string, city, postalCode, country string, streetFlag, cityFlag, postalFlag, countryFlag string, path ...string) error {
	shouldSet := anyFlagChanged(cmd, streetFlag, cityFlag, postalFlag, countryFlag) ||
		((len(street) > 0 || city != "" || postalCode != "" || country != "") && !bodyHasValue(body, path...))
	if !shouldSet {
		return nil
	}

	address := map[string]any{}
	if len(street) > 0 {
		address["street"] = street
	}
	if city != "" {
		address["city"] = city
	}
	if postalCode != "" {
		address["postal_code"] = postalCode
	}
	if country != "" {
		address["country"] = country
	}
	if len(address) == 0 {
		return nil
	}
	return setBodyValue(body, address, path...)
}

func validatePIAddress(body map[string]any, label string, path ...string) error {
	if !bodyHasValue(body, path...) {
		return nil
	}

	required := map[string][]string{
		fmt.Sprintf("%s.street", label):      append(append([]string{}, path...), "street"),
		fmt.Sprintf("%s.city", label):        append(append([]string{}, path...), "city"),
		fmt.Sprintf("%s.postal_code", label): append(append([]string{}, path...), "postal_code"),
		fmt.Sprintf("%s.country", label):     append(append([]string{}, path...), "country"),
	}
	return requireBodyFields(body, required)
}

func applyPIBACSFlags(cmd *cobra.Command, body map[string]any, account, sortCode, accountFlag, sortCodeFlag string, path ...string) error {
	shouldSet := anyFlagChanged(cmd, accountFlag, sortCodeFlag) ||
		((account != "" || sortCode != "") && !bodyHasValue(body, path...))
	if !shouldSet {
		return nil
	}

	bacs := map[string]any{}
	if account != "" {
		bacs["account"] = account
	}
	if sortCode != "" {
		bacs["sort_code"] = sortCode
	}
	if len(bacs) == 0 {
		return nil
	}
	return setBodyValue(body, bacs, path...)
}

func validatePIBACS(body map[string]any, label string, path ...string) error {
	if !bodyHasValue(body, path...) {
		return nil
	}

	required := map[string][]string{
		fmt.Sprintf("%s.account", label):   append(append([]string{}, path...), "account"),
		fmt.Sprintf("%s.sort_code", label): append(append([]string{}, path...), "sort_code"),
	}
	return requireBodyFields(body, required)
}

func applyPIAmountFlags(cmd *cobra.Command, body map[string]any, currency, value, currencyFlag, valueFlag string, path ...string) error {
	if err := applyStringFlag(cmd, body, currencyFlag, currency, append(path, "currency")...); err != nil {
		return err
	}
	if strings.TrimSpace(value) == "" && !cmd.Flags().Changed(valueFlag) {
		return nil
	}
	return applyDecimalStringFlag(cmd, body, valueFlag, value, append(path, "value")...)
}

func validatePIAmount(body map[string]any, label string, path ...string) error {
	if !bodyHasValue(body, path...) {
		return nil
	}

	required := map[string][]string{
		fmt.Sprintf("%s.currency", label): append(append([]string{}, path...), "currency"),
		fmt.Sprintf("%s.value", label):    append(append([]string{}, path...), "value"),
	}
	return requireBodyFields(body, required)
}

func applyPIConsentPeriodicFlags(cmd *cobra.Command, body map[string]any, currency, value, interval, alignment string) error {
	shouldSet := anyFlagChanged(cmd, "periodic-amount-currency", "periodic-amount-value", "periodic-interval", "periodic-alignment") ||
		((currency != "" || value != "" || interval != "" || alignment != "") && !bodyHasValue(body, "constraints", "periodic_amounts"))
	if !shouldSet {
		return nil
	}

	entry := map[string]any{}
	if currency != "" || value != "" {
		amount := map[string]any{}
		if currency != "" {
			amount["currency"] = currency
		}
		if strings.TrimSpace(value) != "" {
			parsed, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
			if err != nil {
				return fmt.Errorf("parse --periodic-amount-value: %w", err)
			}
			amount["value"] = parsed
		}
		if len(amount) > 0 {
			entry["amount"] = amount
		}
	}
	if interval != "" {
		entry["interval"] = interval
	}
	if alignment != "" {
		entry["alignment"] = alignment
	}
	if len(entry) == 0 {
		return nil
	}
	return setBodyValue(body, []any{entry}, "constraints", "periodic_amounts")
}

func validatePIConsentPeriodicFlags(body map[string]any) error {
	if !bodyHasValue(body, "constraints", "periodic_amounts") {
		return nil
	}

	entry, ok, err := firstObjectFromArrayPath(body, "constraints", "periodic_amounts")
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("constraints.periodic_amounts must contain at least one object")
	}

	rawAmount, ok := entry["amount"]
	if !ok {
		return fmt.Errorf("constraints.periodic_amounts[0].amount is required")
	}
	if _, ok := entry["interval"]; !ok {
		return fmt.Errorf("constraints.periodic_amounts[0].interval is required")
	}
	if _, ok := entry["alignment"]; !ok {
		return fmt.Errorf("constraints.periodic_amounts[0].alignment is required")
	}

	amountBody, ok := rawAmount.(map[string]any)
	if !ok {
		return fmt.Errorf("constraints.periodic_amounts[0].amount must be an object")
	}
	return validatePIAmount(map[string]any{"amount": amountBody}, "constraints.periodic_amounts[0].amount", "amount")
}

func firstObjectFromArrayPath(body map[string]any, path ...string) (map[string]any, bool, error) {
	raw, ok := bodyValue(body, path...)
	if !ok {
		return nil, false, nil
	}

	switch typed := raw.(type) {
	case []any:
		if len(typed) == 0 {
			return nil, false, nil
		}
		entry, ok := typed[0].(map[string]any)
		if !ok {
			return nil, false, fmt.Errorf("%s must contain objects", strings.Join(path, "."))
		}
		return entry, true, nil
	case []map[string]any:
		if len(typed) == 0 {
			return nil, false, nil
		}
		return typed[0], true, nil
	default:
		return nil, false, fmt.Errorf("%s must be an array", strings.Join(path, "."))
	}
}
