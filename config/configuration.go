package config

import "github.com/adaptant-labs/consent-receipt-go/api"

type Configuration struct {
	Controller api.DataController `mapstructure:"controller"`
}
