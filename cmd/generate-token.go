package cmd

import (
	"fmt"
	"github.com/adaptant-labs/consent-receipt-go/api"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/cobra"
	"log"
)

func generateJwtToken() (string, error) {
	service := api.NewServiceSinglePurpose("testing", api.DefaultPurpose())

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

	cr.GenerateJurisdictions()

	if cr.PolicyUrl == "" {
		cr.PolicyUrl = cfg.Config.PrivacyPolicyUrl
	}

	// Create the Claims
	claims := cr.GenerateClaims()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(cfg.Config.SigningKey))
	if err != nil {
		return "", err
	}

	return ss, nil
}

// generateTokenCmd represents the token command
var generateTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Generate a new JWT token",
	Run: func(cmd *cobra.Command, args []string) {
		if len(cfg.Controllers) == 0 {
			log.Fatal("Missing controller definition(s)")
		}

		token, err := generateJwtToken()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(token)
	},
}

func init() {
	generateCmd.AddCommand(generateTokenCmd)
}
