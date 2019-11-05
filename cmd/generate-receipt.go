package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// generateReceiptCmd represents the receipt command
var generateReceiptCmd = &cobra.Command{
	Use:   "receipt",
	Short: "Generate a JSON-based Consent Receipt",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("receipt called")
	},
}

func init() {
	generateCmd.AddCommand(generateReceiptCmd)
}
