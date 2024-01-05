package blockchain

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	AddressBalance  = "address"
	ContractBalance = "contract"
)

type Balance struct {
	client          *ethclient.Client
	privateKey      string
	contractAddress string
}

func NewBalance(client *ethclient.Client, priateKey string, contractAddress string) *Balance {
	return &Balance{
		client:          client,
		privateKey:      priateKey,
		contractAddress: contractAddress,
	}
}

func (b *Balance) GetContractBalance(ctx context.Context) (float64, error) {
	return b.GetBalance(ctx, b.contractAddress)
}

func (b *Balance) GetBalance(ctx context.Context, address string) (float64, error) {
	value, err := b.client.BalanceAt(ctx, common.HexToAddress(address), nil)
	if err != nil {
		return 0, err
	}

	ethValue, _ := weiToEther(value).Float64()

	return ethValue, nil
}
