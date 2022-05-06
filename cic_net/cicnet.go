package cic_net

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
)

type CicNet struct {
	ethClient  *w3.Client
	tokenIndex common.Address
}

func NewCicNet(rpcEndpoint string, tokenIndex common.Address) (*CicNet, error) {
	ethClient, err := w3.Dial(rpcEndpoint)
	if err != nil {
		return nil, err
	}

	return &CicNet{
		ethClient:  ethClient,
		tokenIndex: tokenIndex,
	}, nil
}

func (c *CicNet) Close() error {
	err := c.ethClient.Close()
	if err != nil {
		return err
	}

	return nil
}
