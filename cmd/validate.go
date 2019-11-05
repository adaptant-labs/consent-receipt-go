package cmd

import (
	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate a Consent Receipt",
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
