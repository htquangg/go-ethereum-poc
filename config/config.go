package config

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const envPrefix = "AE"

var (
	Filename string

	App AppConfig

	// envVars list of environment variables read by app. The name should match with a struct field.
	// The dots will be replaced by underscores, it will be capitalized and the envPrefix will be added
	// i.e: blockchain.pk => AE_BLOCKCHAIN_PK
	envVars = []string{
		"blockchain.pk",
	}
)

type (
	AppConfig struct {
		Blockchain BlockChainConfig
		Contract   ContractConfig
	}

	BlockChainConfig struct {
		HTTP       string `mapstructure:"http"`
		WS         string `mapstructure:"ws"`
		PrivateKey string `mapstructure:"pk"`
	}

	ContractConfig struct {
		Address   string `mapstructure:"address"`
		GasLimit  int64  `mapstructure:"gas_limit"`
		GasPrice  int64  `mapstructure:"gas_price"`
		WeiFounds int64  `mapstructure:"wei_founds"`
	}
)

// The precedence to override a configuration is: flag -> environment variable -> configuration field
func Setup(cmd *cobra.Command, _ []string) error {
	v := viper.New()
	v.SetConfigFile(Filename)
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")

	v.SetEnvPrefix(envPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		_ = v.BindPFlag(f.Name, cmd.Flags().Lookup(f.Name))
	})

	for _, env := range envVars {
		_ = v.BindPFlag(env, cmd.Flags().Lookup(env))
	}

	err = v.Unmarshal(&App)
	if err != nil {
		return err
	}

	return nil
}
