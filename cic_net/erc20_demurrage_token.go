package cic_net

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
	"math/big"
)

type DemurrageToken struct {
	Name     string
	Symbol   string
	Decimals big.Int
}

func (c *CicNet) TokenInfo(ctx context.Context, tokenAddress common.Address) (DemurrageToken, error) {
	var (
		tokenName     string
		tokenSymbol   string
		tokenDecimals big.Int
	)

	err := c.ethClient.CallCtx(
		ctx,
		eth.CallFunc(w3.MustNewFunc("name()", "string"), tokenAddress).Returns(&tokenName),
		eth.CallFunc(w3.MustNewFunc("symbol()", "string"), tokenAddress).Returns(&tokenSymbol),
		eth.CallFunc(w3.MustNewFunc("decimals()", "uint256"), tokenAddress).Returns(&tokenDecimals),
	)
	if err != nil {
		return DemurrageToken{}, err
	}

	return DemurrageToken{
		Name:     tokenName,
		Symbol:   tokenSymbol,
		Decimals: tokenDecimals,
	}, nil
}
