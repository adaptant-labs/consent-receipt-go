package config

import "github.com/adaptant-labs/consent-receipt-go/api"

type ConfigurationOptions struct {
	SigningKey	string `mapstructure:"signing-key"`
}

type Configuration struct {
	Config		ConfigurationOptions
	Controller	api.DataController `mapstructure:"controller"`
}
