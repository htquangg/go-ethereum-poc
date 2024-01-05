package runner

import (
	"context"
	"errors"

	"github.com/spf13/cobra"
)

func New(ctx context.Context)*cobra.Command {
	runCommand := &cobra.Command{
		Use: "run",
		Short: "Run contract methods in the blockchain",
		RunE: func(cmd *cobra.Command, args []string) error{
			return errors.New("please specify a subcommand: [balance, transfer]")
		},
	}

	runCommand.AddCommand(NewBalanceRunner(ctx))

	return runCommand
}
