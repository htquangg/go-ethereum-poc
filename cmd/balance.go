package cmd

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/htquangg/go-ethereum-poc/config"
	"github.com/htquangg/go-ethereum-poc/pkg/blockchain"
	util "github.com/htquangg/go-ethereum-poc/pkg/utils"
	"github.com/htquangg/go-ethereum-poc/pkg/visualize"

	"github.com/spf13/cobra"
)

var balanceCmd = &cobra.Command{
	Use:   "balances",
	Short: "Used to create, read and update balances",
	Run: func(_ *cobra.Command, _ []string) {
	},
}

var balancesGetCmd = &cobra.Command{
	Example: `balances get <address A> <address B>..."`,
	Short:   "Used to retrieve balances by address",
	Use:     "get [addresses]",
	DisableFlagsInUseLine: true,
	Args:    cobra.MinimumNArgs(1),
	Run:     getBalancesByAddresses,
}

func getBalancesByAddresses(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	client, err := ethclient.Dial(config.App.Blockchain.HTTP)
	if err != nil {
		return
	}

	runner := blockchain.NewBalance(
		client,
		config.App.Blockchain.PrivateKey,
		config.App.Contract.Address,
	)

	balances := make(map[string]float64, len(args))
	for _, addr := range args {
		balance, err := runner.GetBalance(ctx, addr)
		if err != nil {
			util.HandleError(err, "Unable get balance")
		}
		balances[addr] = balance
	}

	visualize.PrintAllSecretDetails(balances)
}

func init() {
	rootCmd.AddCommand(balanceCmd)

	balanceCmd.AddCommand(balancesGetCmd)
}
