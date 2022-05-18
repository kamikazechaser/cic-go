package net

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/grassrootseconomics/cic-go/provider"
	"github.com/lmittmann/w3/module/eth"
)

func (c *CicNet) LastNonce(ctx context.Context, address common.Address) (uint64, error) {
	var nonce uint64

	err := c.provider.EthClient.CallCtx(
		ctx,
		eth.Nonce(address, nil).Returns(&nonce),
	)
	if err != nil {
		return 0, err
	}

	return nonce, nil
}

func (c *CicNet) signAndCall(ctx context.Context, input []byte, txData provider.WriteTx) (common.Hash, error) {
	var txHash common.Hash

	tx, err := types.SignNewTx(&txData.PrivateKey, c.provider.Signer, &types.LegacyTx{
		To:       &txData.To,
		Nonce:    txData.Nonce,
		Data:     input,
		Gas:      txData.GasLimit,
		GasPrice: big.NewInt(1),
	})

	err = c.provider.EthClient.CallCtx(
		ctx,
		eth.SendTransaction(tx).Returns(&txHash),
	)
	if err != nil {
		return [32]byte{}, err
	}

	return txHash, nil
}
