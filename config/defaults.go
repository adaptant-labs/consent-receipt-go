package config

import (
	"github.com/adaptant-labs/consent-receipt-go/api/keys"
	"log"
)

func (cfg *Configuration) SetDefaults() {
	if cfg.Config.SigningKey == "" {
		cfg.Config.SigningKey = "totally-secret-key"
	}

	if cfg.Config.PublicKeyFile != "" && cfg.Config.PrivateKeyFile != "" {
		var err error

		cfg.PrivateKey, cfg.PublicKey, err = keys.LoadKeys(cfg.Config.PrivateKeyFile, cfg.Config.PublicKeyFile)
		if err != nil {
			log.Fatal(err)
		}
	}
}
