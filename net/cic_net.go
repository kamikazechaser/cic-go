package net

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

type CicNet struct {
	ethClient    *w3.Client
	tokenIndex   common.Address
	kitabuSigner types.Signer
}

type WriteTx struct {
	from       common.Address
	to         common.Address
	gasLimit   uint64
	nonce      uint64
	privateKey ecdsa.PrivateKey
}

func NewCicNet(rpcEndpoint string, tokenIndex common.Address) (*CicNet, error) {
	ethClient, err := w3.Dial(rpcEndpoint)
	if err != nil {
		return nil, err
	}

	return &CicNet{
		ethClient:    ethClient,
		tokenIndex:   tokenIndex,
		kitabuSigner: types.NewEIP155Signer(big.NewInt(kitabuMainnetChainId)),
	}, nil
}

func (c *CicNet) Close() error {
	err := c.ethClient.Close()
	if err != nil {
		return err
	}

	return nil
}
