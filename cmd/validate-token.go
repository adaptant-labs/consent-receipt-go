package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// validateTokenCmd represents the token command
var validateTokenCmd = &cobra.Command{
	Use:   "token <JWT Token>",
	Short: "Validate a JWT-encoded Consent Receipt token",
	Args:	cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tokenString := args[0]
		fmt.Println(tokenString)
	},
}

func init() {
	validateCmd.AddCommand(validateTokenCmd)
}
