package cmd

import (
	"fmt"
	"github.com/adaptant-labs/consent-receipt-go/api"
	"github.com/dgrijalva/jwt-go"

	"github.com/spf13/cobra"
)

func generateJwtToken() (string, error) {
	controller := api.NewDataController()
	cr := api.NewConsentReceipt()
	cr.AddDataController(controller)

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
			panic(err)
		}

		fmt.Println(token)
	},
}

func init() {
	generateCmd.AddCommand(generateTokenCmd)
}
