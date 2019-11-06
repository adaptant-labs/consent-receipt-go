package cmd

import (
	"fmt"
	"github.com/adaptant-labs/consent-receipt-go/api/purpose"
	"github.com/spf13/cobra"
	"strings"
)

// purposesCmd represents the purposes command
var purposesCmd = &cobra.Command{
	Use:   "purposes",
	Short: "List available purposes of processing",
	Run: func(cmd *cobra.Command, args []string) {
		purposes := purpose.DumpPurposeSpecifications()
		lines := strings.Join(purposes[:], "\n")
		fmt.Println(lines)
	},
}

func init() {
	rootCmd.AddCommand(purposesCmd)
}
