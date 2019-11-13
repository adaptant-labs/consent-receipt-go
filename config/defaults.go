package config

import "github.com/adaptant-labs/consent-receipt-go/api/keys"

func (cfg *Configuration) SetDefaults() {
	if cfg.Config.SigningKey == "" {
		cfg.Config.SigningKey = "totally-secret-key"
	}

	if cfg.Config.PublicKeyFile != "" && cfg.Config.PrivateKeyFile != "" {
		cfg.PrivateKey, cfg.PublicKey, _ = keys.InitKeys(cfg.Config.PrivateKeyFile, cfg.Config.PublicKeyFile)
	}
}
