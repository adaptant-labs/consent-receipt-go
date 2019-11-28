package cmd

import (
	"encoding/json"
	"github.com/adaptant-labs/consent-receipt-go/api"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func generateJsonReceipt() error {
	service := api.NewServiceMultiPurpose(serviceName, purposes)

	// The first controller is the primary controller
	controller := cfg.Controllers[0]
	cr := controller.NewConsentReceipt()

	cr.AddService(service)

	// Add any additional controllers
	if len(cfg.Controllers) > 1 {
		for _, iter := range cfg.Controllers[1:] {
			cr.AddDataController(&iter)
		}
	}

	cr.GenerateJurisdictions()

	if len(cr.SensitiveCategories) > 1 {
		cr.Sensitive = true
	}

	if cr.PolicyUrl == "" {
		cr.PolicyUrl = cfg.Config.PrivacyPolicyUrl
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "\t")

	return encoder.Encode(cr)
}

// generateReceiptCmd represents the receipt command
var generateReceiptCmd = &cobra.Command{
	Use:   "receipt",
	Short: "Generate a JSON-based Consent Receipt",
	Run: func(cmd *cobra.Command, args []string) {
		if len(cfg.Controllers) == 0 {
			log.Fatal("Missing controller definition(s)")
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
