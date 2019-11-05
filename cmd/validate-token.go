package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/cobra"
	"os"
)

var dumpToken bool

func dumpJwtToken(token *jwt.Token) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "\t")
	encoder.Encode(claims)
}

func parseJwtToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return signingKey, nil
	})
}

// validateTokenCmd represents the token command
var validateTokenCmd = &cobra.Command{
	Use:   "token <JWT Token>",
	Short: "Validate a JWT-encoded Consent Receipt token",
	Args:	cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token, err := parseJwtToken(args[0])
		if err != nil || !token.Valid {
			fmt.Println("Token is invalid")
		} else {
			if dumpToken {
				dumpJwtToken(token)
			} else {
				fmt.Println("Token is valid")
			}
		}
	},
}

func init() {
	validateTokenCmd.Flags().BoolVarP(&dumpToken, "dump", "d", false, "dump token contents")
	validateCmd.AddCommand(validateTokenCmd)
}
