package balance

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/grassrootseconomics/cic-go/provider"
)

type BatchBalance struct {
	provider      *provider.Provider
	batchContract common.Address
}

func NewBatchBalance(rpcProvider *provider.Provider, batchContract common.Address) (*BatchBalance, error) {
	return &BatchBalance{
		provider:      rpcProvider,
		batchContract: batchContract,
	}, nil
}
