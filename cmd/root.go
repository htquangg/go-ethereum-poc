package cmd

import (
	"os"
	"strings"

	"github.com/htquangg/go-ethereum-poc/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "aecli",
	Short:   "",
	Long:    "",
	Version: "",
}

func init() {
	cobra.OnInitialize(initLog)
	rootCmd.PersistentFlags().StringP("log-level", "l", "info", "log level (trace, debug, info, warn, error, fatal)")
	rootCmd.PersistentPreRunE = config.Setup
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initLog() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	ll, err := rootCmd.Flags().GetString("log-level")
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	switch strings.ToLower(ll) {
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "err", "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}
