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

	controller := cfg.Controller
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
		if cfg.Controller.ControllerName == "" {
			log.Fatal("Missing controller definition")
		}

		err := generateJsonReceipt()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	generateCmd.AddCommand(generateReceiptCmd)
}
