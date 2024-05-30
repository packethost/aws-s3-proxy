// Package cmd handles the cli interface for the proxy
package cmd

import (
	"log"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	cfgFile string
	logger  *zap.SugaredLogger
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aws-s3-proxy",
	Short: "An http proxy to an S3 api",
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.s3-proxy.yaml)")
	rootCmd.PersistentFlags().Bool("debug", false, "Enable debug logging")
	viperBindFlag("logging.debug", rootCmd.PersistentFlags().Lookup("debug"))
	rootCmd.PersistentFlags().Bool("pretty", false, "Enable pretty (human readable) logging output")
	viperBindFlag("logging.pretty", rootCmd.PersistentFlags().Lookup("pretty"))
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	setupLogging()

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			logger.Fatalf("unable to find home directory: %v", err)
		}

		// Search config in home directory with name ".s3-proxy" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".s3-proxy")
	}

	// Check for ENV variables set
	// All ENV vars will be prefixed with "S3_PROXY_"
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("s3_proxy")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		logger.Infof("unable to read in config file: %v", err)
	}

	logger.Infof("using config file: %s", viper.ConfigFileUsed())
}

// viperBindFlag provides a wrapper around the viper pflag bindings that handles error checks
func viperBindFlag(name string, flag *pflag.Flag) {
	if err := viper.BindPFlag(name, flag); err != nil {
		logger.Fatalf("failed to bind flag: %v", err)
	}
}

// viperBindEnv provides a wrapper around the viper env var bindings that handles error checks
func viperBindEnv(input ...string) {
	if err := viper.BindEnv(input...); err != nil {
		logger.Fatalf("failed to bind environment variable: %v", err)
	}
}

func setupLogging() {
	cfg := zap.NewProductionConfig()
	if viper.GetBool("logging.pretty") {
		cfg = zap.NewDevelopmentConfig()
	}

	// Level is already at InfoLevel, so only change to DebugLevel
	// when the flag is set
	if viper.GetBool("logging.debug") {
		cfg.Level.SetLevel(zap.DebugLevel)
	}

	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	logger = l.Sugar()
}
