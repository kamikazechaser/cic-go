package provider

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lmittmann/w3"
)

const (
	kitabuMainnetChainId = 6060
)

type Provider struct {
	EthClient *w3.Client
	Signer    types.Signer
}

type WriteTx struct {
	From       common.Address
	To         common.Address
	GasLimit   uint64
	Nonce      uint64
	PrivateKey ecdsa.PrivateKey
}

func NewRpcProvider(rpcEndpoint string) (*Provider, error) {
	ethClient, err := w3.Dial(rpcEndpoint)
	if err != nil {
		return nil, err
	}

	return &Provider{
		EthClient: ethClient,
		Signer:    types.NewEIP155Signer(big.NewInt(kitabuMainnetChainId)),
	}, nil
}

func (c *Provider) CLose() error {
	err := c.EthClient.Close()
	if err != nil {
		return err
	}

	return nil
}
