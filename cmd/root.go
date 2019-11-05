package cmd

import (
  "fmt"
  "github.com/adaptant-labs/consent-receipt-go/config"
  homedir "github.com/mitchellh/go-homedir"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
  "log"
  "os"
)

var (
  cfgFile string
  cfg = &config.Configuration{}
)

var rootCmd = &cobra.Command{
  Use:   "consent-receipt-go",
  Short: "Utilities for working with Consent Receipts",
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)

  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.consent-receipt-go.toml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
  if cfgFile != "" {
    // Use config file from the flag.
    viper.SetConfigFile(cfgFile)
  } else {
    // Find home directory.
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    // Search config in home directory with name ".consent-receipt-go" (without extension).
    viper.AddConfigPath(home)
    viper.SetConfigName(".consent-receipt-go")
  }

  viper.AutomaticEnv() // read in environment variables that match

  // If a config file is found, read it in.
  if err := viper.ReadInConfig(); err == nil {
    fmt.Fprintf(os.Stderr, "Using config file: %s\n", viper.ConfigFileUsed())
    err = viper.Unmarshal(cfg)
    if err != nil {
      log.Fatal(err)
    }

    // Set any other defaults that may be missing from the config file
    cfg.SetDefaults()
  }
}
