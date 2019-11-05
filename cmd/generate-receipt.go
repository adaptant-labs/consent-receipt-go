package cmd

import (
	"encoding/json"
	"github.com/adaptant-labs/consent-receipt-go/api"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func generateJsonReceipt() error {
	controller := api.NewDataController()
	cr := api.NewConsentReceipt()
	cr.AddDataController(controller)

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "\t")

	return encoder.Encode(cr)
}

// generateReceiptCmd represents the receipt command
var generateReceiptCmd = &cobra.Command{
	Use:   "receipt",
	Short: "Generate a JSON-based Consent Receipt",
	Run: func(cmd *cobra.Command, args []string) {
		err := generateJsonReceipt()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	generateCmd.AddCommand(generateReceiptCmd)
}
