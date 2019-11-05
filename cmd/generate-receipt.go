package cmd

import (
	"encoding/json"
	"github.com/adaptant-labs/consent-receipt-go/api"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func generateJsonReceipt() error {
	purpose := api.NewPurpose("testing", true, "n/a")
	service := api.NewServiceSinglePurpose("testing", purpose)

	address := api.NewPostalAddress("DE", "Deisenhofen", "82041", "Bahnhofstr. 36")
	address.Region = "BY"

	controller := api.NewDataController("Adaptant Solutions AG", "Max Musterman", "compliance@adaptant.io", "49-00-00000000", address)

	cr := controller.NewConsentReceipt()
	cr.AddService(service)

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
