package balance

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
)

func (c *BatchBalance) TokensBalance(ctx context.Context, owner common.Address, tokens []common.Address) ([]*big.Int, error) {
	var balancesResults []*big.Int

	err := c.ethClient.CallCtx(
		ctx,
		eth.CallFunc(w3.MustNewFunc("tokensBalance(address owner, address[] contracts)", "uint256[]"), c.batchContract, owner, tokens).Returns(&balancesResults),
	)
	if err != nil {
		return nil, err
	}

	return balancesResults, nil
}
