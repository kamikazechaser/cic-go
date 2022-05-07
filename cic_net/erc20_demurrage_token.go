package cic_net

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
	"math/big"
)

type DemurrageTokenInfo struct {
	DemurrageAmount         big.Int
	DemurrageTimestamp      big.Int
	MinimumParticipantSpend big.Int
	ResolutionFactor        big.Int
	PeriodStart             big.Int
	PeriodDuration          big.Int
	TaxLevel                big.Int
	ActualPeriod            big.Int
	RedistributionCount     big.Int
	IsDemurrageToken        bool
}

func (c *CicNet) DemurrageTokenInfo(ctx context.Context, tokenAddress common.Address) (DemurrageTokenInfo, error) {
	var (
		demurrageAmount         big.Int
		demurrageTimestamp      big.Int
		minimumParticipantSpend big.Int
		resolutionFactor        big.Int
		periodStart             big.Int
		periodDuration          big.Int
		taxLevel                big.Int
		actualPeriod            big.Int
		redistributionCount     big.Int
	)

	err := c.ethClient.CallCtx(
		ctx,
		eth.CallFunc(w3.MustNewFunc("demurrageAmount()", "uint128"), tokenAddress).Returns(&demurrageAmount),
		eth.CallFunc(w3.MustNewFunc("demurrageTimestamp()", "uint256"), tokenAddress).Returns(&demurrageTimestamp),
		eth.CallFunc(w3.MustNewFunc("minimumParticipantSpend()", "uint256"), tokenAddress).Returns(&minimumParticipantSpend),
		eth.CallFunc(w3.MustNewFunc("resolutionFactor()", "uint256"), tokenAddress).Returns(&resolutionFactor),
		eth.CallFunc(w3.MustNewFunc("periodStart()", "uint256"), tokenAddress).Returns(&periodStart),
		eth.CallFunc(w3.MustNewFunc("periodDuration()", "uint256"), tokenAddress).Returns(&periodDuration),
		eth.CallFunc(w3.MustNewFunc("taxLevel()", "uint256"), tokenAddress).Returns(&taxLevel),
		eth.CallFunc(w3.MustNewFunc("actualPeriod()", "uint256"), tokenAddress).Returns(&actualPeriod),
		eth.CallFunc(w3.MustNewFunc("redistributionCount()", "uint256"), tokenAddress).Returns(&redistributionCount),
	)

	if err != nil {
		return DemurrageTokenInfo{}, err
	}

	return DemurrageTokenInfo{
		DemurrageAmount:         demurrageAmount,
		DemurrageTimestamp:      demurrageTimestamp,
		MinimumParticipantSpend: minimumParticipantSpend,
		ResolutionFactor:        resolutionFactor,
		PeriodStart:             periodStart,
		PeriodDuration:          periodDuration,
		TaxLevel:                taxLevel,
		ActualPeriod:            actualPeriod,
		RedistributionCount:     redistributionCount,
	}, nil
}

func (c *CicNet) BaseBalanceOf(ctx context.Context, tokenAddress common.Address, accountAddress common.Address) (big.Int, error) {
	var balance big.Int

	err := c.ethClient.CallCtx(
		ctx,
		eth.CallFunc(w3.MustNewFunc("baseBalanceOf(address _account)", "uint256"), tokenAddress, accountAddress).Returns(&balance),
	)
	if err != nil {
		return big.Int{}, err
	}

	return balance, nil
}
