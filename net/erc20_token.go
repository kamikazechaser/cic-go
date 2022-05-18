package net

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
)

type ERC20Token struct {
	Name     string
	Symbol   string
	Decimals big.Int
}

func (c *CicNet) ERC20TokenInfo(ctx context.Context, tokenAddress common.Address) (ERC20Token, error) {
	var (
		tokenName     string
		tokenSymbol   string
		tokenDecimals big.Int
	)

	err := c.provider.EthClient.CallCtx(
		ctx,
		eth.CallFunc(w3.MustNewFunc("name()", "string"), tokenAddress).Returns(&tokenName),
		eth.CallFunc(w3.MustNewFunc("symbol()", "string"), tokenAddress).Returns(&tokenSymbol),
		eth.CallFunc(w3.MustNewFunc("decimals()", "uint256"), tokenAddress).Returns(&tokenDecimals),
	)
	if err != nil {
		return ERC20Token{}, err
	}

	return ERC20Token{
		Name:     tokenName,
		Symbol:   tokenSymbol,
		Decimals: tokenDecimals,
	}, nil
}
func (c *CicNet) BalanceOf(ctx context.Context, tokenAddress common.Address, accountAddress common.Address) (big.Int, error) {
	var balance big.Int

	err := c.provider.EthClient.CallCtx(
		ctx,
		eth.CallFunc(w3.MustNewFunc("balanceOf(address _account)", "uint256"), tokenAddress, accountAddress).Returns(&balance),
	)
	if err != nil {
		return big.Int{}, err
	}

	return balance, nil
}
