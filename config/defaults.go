package config

func (cfg *Configuration) SetDefaults() {
	if cfg.Config.SigningKey == "" {
		cfg.Config.SigningKey = "totally-secret-key"
	}
}
