package balance

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
)

type BatchBalance struct {
	ethClient     *w3.Client
	batchContract common.Address
}

func NewBatchBalance(rpcEndpoint string, batchContract common.Address) (*BatchBalance, error) {
	ethClient, err := w3.Dial(rpcEndpoint)
	if err != nil {
		return nil, err
	}

	return &BatchBalance{
		ethClient:     ethClient,
		batchContract: batchContract,
	}, nil
}

func (c *BatchBalance) Close() error {
	err := c.ethClient.Close()
	if err != nil {
		return err
	}

	return nil
}
