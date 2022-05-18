package net

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/grassrootseconomics/cic-go/provider"
)

type CicNet struct {
	provider   *provider.Provider
	tokenIndex common.Address
}

func NewCicNet(rpcProvider provider.Provider, tokenIndex common.Address) (*CicNet, error) {
	return &CicNet{
		provider:   &rpcProvider,
		tokenIndex: tokenIndex,
	}, nil
}
