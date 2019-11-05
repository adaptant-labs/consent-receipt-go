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

	controller := cfg.Controller
	cr := controller.NewConsentReceipt()
	cr.AddService(service)

	// Create the Claims
	claims := cr.GenerateClaims()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(cfg.Config.SigningKey)
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
		if cfg.Controller.ControllerName == "" {
			log.Fatal("Missing controller definition")
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
