package cmd

import (
	"fmt"
	"github.com/adaptant-labs/consent-receipt-go/api"
	"github.com/adaptant-labs/consent-receipt-go/api/category"
	"github.com/adaptant-labs/consent-receipt-go/api/keys"
	"github.com/adaptant-labs/consent-receipt-go/api/purpose"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var (
	serviceName string
	jurisdiction string
	subjectId string
	terminationPeriod string

	sensitiveNumsStr string
	scopes []string

	purposeNums []int
	purposes []*api.Purpose

	categoryNumsStr []string
	categoryGroups [][]category.DataCategory

	consentReceipt *api.ConsentReceipt
)

func prepareConsentReceipt() *api.ConsentReceipt {
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

	if jurisdiction != "" {
		cr.Jurisdiction = strings.ReplaceAll(jurisdiction, ",", " ")
	} else {
		cr.GenerateJurisdictions()
	}

	if cr.PolicyUrl == "" {
		cr.PolicyUrl = cfg.Config.PrivacyPolicyUrl
	}

	cr.SubjectID = subjectId

	if len(sensitiveNumsStr) > 0 {
		cr.Sensitive = true

		sc := categoriesFromNumString(sensitiveNumsStr)
		for _, v := range sc {
			cr.AddSensitiveCategory(v)
		}
	}

	if cfg.PublicKey != nil {
		cr.PublicKey, _ = keys.Fingerprint(*cfg.PublicKey)
	}

	return cr
}

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

		numCategories := len(categoryGroups)
		numScopes := len(scopes)
		numPurposes := len(purposeNums)

		if numCategories != numPurposes {
			return fmt.Errorf("number of data category specifications (%d) must match number of defined purposes (%d)", numCategories, numPurposes)
		}

		if numScopes > numPurposes {
			return fmt.Errorf("number of scope specifications (%d) cannot exceed number of defined purposes (%d)", numScopes, numPurposes)
		}

		for k, v := range purposeNums {
			spec := purpose.PurposeSpecification(v)
			p := api.NewPurpose(spec, categoryGroups[k], primary, terminationPeriod)
			if k < numScopes {
				p.Scopes = strings.ReplaceAll(scopes[k], ",", " ")
			}
			purposes = append(purposes, p)
			primary = false
		}

		consentReceipt = prepareConsentReceipt()
		return nil
	},
}

func init() {
	generateCmd.PersistentFlags().StringVar(&serviceName, "service", "testing", "Name of the service to generate a receipt for")
	generateCmd.PersistentFlags().StringVar(&subjectId, "id", "", "Unique identifier for the PII Principal / Data Subject")
	generateCmd.PersistentFlags().StringVar(&jurisdiction, "jurisdiction", "DE", "List of jurisdicitions in ISO 3166-1 alpha-2 format")
	generateCmd.PersistentFlags().StringVar(&terminationPeriod, "termination", api.DefaultTermination, "Termination period")
	generateCmd.PersistentFlags().IntSliceVarP(&purposeNums, "purposes", "p", []int{ purpose.CoreFunction.Number() }, "List of purposes to include")
	generateCmd.PersistentFlags().StringArrayVarP(&categoryNumsStr, "categories", "c", []string{ strconv.Itoa(category.Biographical.Number()) }, "List of data categories to include, per purpose")
	generateCmd.PersistentFlags().StringArrayVarP(&scopes, "scopes", "s", nil, "List of auth scopes to include, per purpose")
	generateCmd.PersistentFlags().StringVar(&sensitiveNumsStr, "sensitive", "", "List of sensitive data categories used by service")

	rootCmd.AddCommand(generateCmd)
}
