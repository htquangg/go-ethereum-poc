package commands

import (
	"context"
	"errors"

	"github.com/htquangg/go-ethereum-poc/cmd/commands/runner"
	"github.com/htquangg/go-ethereum-poc/config"

	"github.com/spf13/cobra"
)

func NewRootCommand(ctx context.Context) *cobra.Command {
	rootCommand := &cobra.Command{
		Use:   "aecli",
		Short: "Run the wallet service",

		PersistentPreRunE: config.Setup,
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("please specify a command: [run]")
		},
	}

	rootCommand.Flags().StringVarP(
		&config.Filename,
		"config",
		"c",
		"config/dev.config.yaml",
		"Relative path to the config file",
	)

	rootCommand.PersistentFlags().String("blockchain.pk", "", "Account private key")
	rootCommand.AddCommand(runner.New(ctx))

	return rootCommand
}
