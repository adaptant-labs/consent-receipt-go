package cmd

import (
	"fmt"
	"github.com/adaptant-labs/consent-receipt-go/api"
	"github.com/adaptant-labs/consent-receipt-go/api/category"
	"github.com/adaptant-labs/consent-receipt-go/api/purpose"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var (
	serviceName string
	terminationPeriod string

	purposeNums []int
	purposes []*api.Purpose

	categoryNumsStr []string
	categoryGroups [][]category.DataCategory
)

func categoriesFromNumString(numStr string) []category.DataCategory {
	nums := strings.Split(numStr, ",")
	c := make([]category.DataCategory, len(nums))
	for i := range c {
		num, err := strconv.Atoi(nums[i])
		if err == nil {
			c[i] = category.DataCategory(num)
		}
	}
	return c
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a new Consent Receipt",
	TraverseChildren: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		primary := true

		for _, v := range categoryNumsStr {
			categoryGroups = append(categoryGroups, categoriesFromNumString(v))
		}

		if len(categoryGroups) != len(purposeNums) {
			return fmt.Errorf("Number of data category specifications (%d) must match number of defined purposes (%d)", len(categoryGroups), len(purposeNums))
		}

		for k, v := range purposeNums {
			spec := purpose.PurposeSpecification(v)
			p := api.NewPurpose(spec, categoryGroups[k], primary, terminationPeriod)
			purposes = append(purposes, p)
			primary = false
		}

		return nil
	},

}

func init() {
	generateCmd.PersistentFlags().StringVar(&serviceName, "service", "testing", "Name of the service to generate a receipt for")
	generateCmd.PersistentFlags().StringVar(&terminationPeriod, "termination", api.DefaultTermination, "Termination period")
	generateCmd.PersistentFlags().IntSliceVarP(&purposeNums, "purposes", "p", []int{ purpose.CoreFunction.Number() }, "List of purposes to include")
	generateCmd.PersistentFlags().StringArrayVarP(&categoryNumsStr, "categories", "c", []string{ string(category.Biographical.Number()) }, "List of data categories to include, per purpose")

	rootCmd.AddCommand(generateCmd)
}
