package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var outputFile string

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion [output file]",
	Short: "Generates bash completion scripts",
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			outputFile := args[0]
			file, err := os.Open(outputFile)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			fmt.Println("Writing completion script to", outputFile)
			_ = rootCmd.GenBashCompletion(file)
		} else {
			_ = rootCmd.GenBashCompletion(os.Stdout)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
