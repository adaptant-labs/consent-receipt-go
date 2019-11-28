package cmd

import (
	"fmt"
	"github.com/adaptant-labs/consent-receipt-go/api/purpose"
	"github.com/muesli/reflow"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// describeCmd represents the describe command
var describeCmd = &cobra.Command{
	Use:   "describe <purposeID>",
	Short: "Describe a specific purpose of processing",
	Args:	cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		purposeId, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}

		purpose := purpose.PurposeSpecification(purposeId)

		fmt.Printf("\x1B[36m%s\x1B[0m\n\n", purpose.PurposeWithPrefix())

		wrapped := reflow.ReflowString(purpose.Description(), 80)
		fmt.Println(wrapped)
	},
}

func init() {
	purposesCmd.AddCommand(describeCmd)
}
