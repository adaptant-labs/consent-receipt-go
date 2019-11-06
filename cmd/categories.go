package cmd

import (
	"fmt"
	"github.com/adaptant-labs/consent-receipt-go/api/category"
	"github.com/spf13/cobra"
	"strings"
)

// categoriesCmd represents the categories command
var categoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "List available categories of data",
	Run: func(cmd *cobra.Command, args []string) {
		categories := category.DumpDataCategories()
		lines := strings.Join(categories[:], "\n")
		fmt.Println(lines)
	},
}

func init() {
	rootCmd.AddCommand(categoriesCmd)
}
