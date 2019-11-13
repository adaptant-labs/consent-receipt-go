package config

import (
	"crypto/rsa"
	"github.com/adaptant-labs/consent-receipt-go/api"
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
