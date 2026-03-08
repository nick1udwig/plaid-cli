package cmd

import (
	"strings"
	"testing"
)

func TestBuildTransferRequirementSubmissionFromFlags(t *testing.T) {
	t.Parallel()

	t.Run("builds submission entry", func(t *testing.T) {
		t.Parallel()

		entry, ok, err := buildTransferRequirementSubmissionFromFlags(transferRequirementSubmissionFlags{
			requirementType: "BUSINESS_NAME",
			value:           "Example LLC",
		})
		if err != nil {
			t.Fatalf("buildTransferRequirementSubmissionFromFlags() error = %v", err)
		}
		if !ok {
			t.Fatal("buildTransferRequirementSubmissionFromFlags() ok = false, want true")
		}
		if got := entry["requirement_type"]; got != "BUSINESS_NAME" {
			t.Fatalf("requirement_type = %#v, want BUSINESS_NAME", got)
		}
	})

	t.Run("requires requirement type when any field is set", func(t *testing.T) {
		t.Parallel()

		_, _, err := buildTransferRequirementSubmissionFromFlags(transferRequirementSubmissionFlags{
			value: "Example LLC",
		})
		if err == nil {
			t.Fatal("buildTransferRequirementSubmissionFromFlags() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "--requirement-type") {
			t.Fatalf("error = %q, want requirement-type guidance", err)
		}
	})
}
