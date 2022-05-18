package net

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
)

func (c *CicNet) EntryCount(ctx context.Context) (big.Int, error) {
	var tokenCount big.Int

	err := c.provider.EthClient.CallCtx(
		ctx,
		eth.CallFunc(w3.MustNewFunc("entryCount()", "uint256"), c.tokenIndex).Returns(&tokenCount),
	)
	if err != nil {
		return big.Int{}, err
	}

	return tokenCount, nil
}

func (c *CicNet) AddressAtIndex(ctx context.Context, index *big.Int) (string, error) {
	var address common.Address

	err := c.provider.EthClient.CallCtx(
		ctx,
		eth.CallFunc(w3.MustNewFunc("entry(uint256 _idx)", "address"), c.tokenIndex, index).Returns(&address),
	)
	if err != nil {
		return "", err
	}

	return address.String(), nil
}
