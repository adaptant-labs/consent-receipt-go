package cmd

import (
	"fmt"
	"github.com/adaptant-labs/consent-receipt-go/api/keys"
	"log"

	"github.com/spf13/cobra"
)

// keygenCmd represents the keygen command
var keygenCmd = &cobra.Command{
	Use:   "keygen",
	Short: "Generate a new public/private signing keys",
	Run: func(cmd *cobra.Command, args []string) {
		err := keys.GenerateKeys()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Signing keys generated")
	},
}

func init() {
	rootCmd.AddCommand(keygenCmd)
}
