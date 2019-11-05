package cmd

import (
	"fmt"
	"github.com/adaptant-labs/consent-receipt-go/api"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/cobra"
	"log"
)

func generateJwtToken() (string, error) {
	purpose := api.NewPurpose("testing", true, "n/a")
	service := api.NewServiceSinglePurpose("testing", purpose)

	address := api.NewPostalAddress("DE", "Deisenhofen", "82041", "Bahnhofstr. 36")
	controller := api.NewDataController("Adaptant Solutions AG", "Max Musterman", "compliance@adaptant.io", "49-00-00000000", address)

	cr := api.NewConsentReceipt()
	cr.AddDataController(controller)
	cr.AddService(service)

	// Create the Claims
	claims := cr.GenerateClaims()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)
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
