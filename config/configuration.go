package config

import "github.com/adaptant-labs/consent-receipt-go/api"

type ConfigurationOptions struct {
	SigningKey	string `mapstructure:"signing-key"`
	PrivacyPolicyUrl	string `mapstructure:"privacy-policy"`
}

type Configuration struct {
	Config		ConfigurationOptions
	Controllers	[]api.DataController `mapstructure:"controller"`
}
