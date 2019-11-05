package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/adaptant-labs/consent-receipt-go/api"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

func validateJsonReceipt(filename string) bool {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var claims api.ConsentReceiptClaims

	err = json.Unmarshal(data, &claims)
	if err != nil {
		return false
	}

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
		if valid {
			fmt.Println("Receipt is valid");
		} else {
			fmt.Println("Receipt is invalid")
		}
	},
}

func init() {
	validateCmd.AddCommand(validateReceiptCmd)
}
