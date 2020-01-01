package config

import (
	"crypto/rsa"
	"github.com/adaptant-labs/consent-receipt-go/api"
	"github.com/dgrijalva/jwt-go"
)

type ConfigurationOptions struct {
	SigningKey       string `mapstructure:"signing-key"`
	PrivateKeyFile   string `mapstructure:"private-key-file"`
	PublicKeyFile    string `mapstructure:"public-key-file"`
	PrivacyPolicyUrl string `mapstructure:"privacy-policy"`
}

type Configuration struct {
	Config      ConfigurationOptions
	Controllers []api.DataController `mapstructure:"controller"`

	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
}

func (c Configuration) GenerateJwtToken(cr api.ConsentReceipt) (string, error) {
	// Create the Claims
	claims := cr.GenerateClaims()

	var signedString string
	var err error

	if c.PrivateKey != nil {
		token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
		signedString, err = token.SignedString(c.PrivateKey)
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedString, err = token.SignedString([]byte(c.Config.SigningKey))
		if err != nil {
			return "", err
		}
	}

	return signedString, err
}
