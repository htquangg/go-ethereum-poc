package runner

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/htquangg/go-ethereum-poc/config"
	"github.com/htquangg/go-ethereum-poc/pkg/blockchain"
	"github.com/spf13/cobra"
)

var ErrInvalidBalanceAddress = errors.New("invalid balance address")

const TimeoutIn = 5 * time.Second

func NewBalanceRunner(ctx context.Context) *cobra.Command {
	var (
		of      string
		address string
	)

	balanceCommand := &cobra.Command{
		Use:   "balance",
		Short: "Get the balance of an address or contract",
		Run: func(_ *cobra.Command, _ []string) {
			err := runBalance(ctx, of, address)
			if err != nil {
				fmt.Printf("[ERROR][BALANCE] something went wrong: %T \n", err)
			}
		},
	}

	balanceCommand.Flags().StringVar(&of, "of", "", "Balance of: address, contract")
	balanceCommand.Flags().StringVar(&address, "address", "", "Target address")

	_ = balanceCommand.MarkFlagRequired("of")

	return balanceCommand
}

func runBalance(ctx context.Context, of string, address string) error {
	ctxCall, cancel := context.WithTimeout(ctx, TimeoutIn)
	defer cancel()

	client, err := ethclient.DialContext(ctxCall, config.App.Blockchain.HTTP)
	if err != nil {
		return err
	}

	runner := blockchain.NewBalance(
		client,
		config.App.Blockchain.PrivateKey,
		config.App.Contract.Address,
	)

	switch of {
	case blockchain.AddressBalance:
		if address == "" {
			return ErrInvalidBalanceAddress
		}
		balance, err := runner.GetBalance(ctx, address)
		if err != nil {
			return err
		}
		fmt.Printf("The contract balance is %f\n", balance)
	case blockchain.ContractBalance:
		balance, err := runner.GetContractBalance(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("The contract balance is %f\n", balance)
	}

	return nil
}
