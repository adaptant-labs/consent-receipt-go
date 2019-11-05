package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// validateReceiptCmd represents the receipt command
var validateReceiptCmd = &cobra.Command{
	Use:   "receipt",
	Short: "Validate a JSON-based Consent Receipt",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("receipt called")
	},
}

func init() {
	validateCmd.AddCommand(validateReceiptCmd)
}
