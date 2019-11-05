package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func validateJsonReceipt(filename string) bool {
	return true
}

// validateReceiptCmd represents the receipt command
var validateReceiptCmd = &cobra.Command{
	Use:   "receipt <JSON file>",
	Short: "Validate a JSON-based Consent Receipt",
	Args:	cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		if !FileExists(filename) {
			log.Fatal("File does not exist")
		}

		valid := validateJsonReceipt(filename)
		if (valid) {
			fmt.Println("Receipt is valid");
		} else {
			fmt.Println("Receipt is invalid")
		}
	},
}

func init() {
	validateCmd.AddCommand(validateReceiptCmd)
}
