package cmd

import "github.com/spf13/cobra"

const paymentInitiationDocPath = "docs/plaid/api/products/payment-initiation/index.md"

func newPaymentInitiationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "payment-initiation",
		Short: "Payment Initiation product commands",
		Long:  "Payment Initiation commands for UK and Europe payment recipients, payments, and payment consents.",
	}

	cmd.AddCommand(newPaymentInitiationRecipientCmd())
	cmd.AddCommand(newPaymentInitiationPaymentCmd())
	cmd.AddCommand(newPaymentInitiationConsentCmd())

	return cmd
}
