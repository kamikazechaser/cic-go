package net

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lmittmann/w3/module/eth"
	"math/big"
)

func (c *CicNet) LastNonce(ctx context.Context, address common.Address) (uint64, error) {
	var nonce uint64

	err := c.ethClient.CallCtx(
		ctx,
		eth.Nonce(address, nil).Returns(&nonce),
	)
	if err != nil {
		return 0, err
	}

	return nonce, nil
}

func (c *CicNet) signAndCall(ctx context.Context, input []byte, txData WriteTx) (common.Hash, error) {
	var txHash common.Hash

	tx, err := types.SignNewTx(&txData.privateKey, c.kitabuSigner, &types.LegacyTx{
		To:       &txData.to,
		Nonce:    txData.nonce,
		Data:     input,
		Gas:      txData.gasLimit,
		GasPrice: big.NewInt(1),
	})

	err = c.ethClient.CallCtx(
		ctx,
		eth.SendTransaction(tx).Returns(&txHash),
	)
	if err != nil {
		return [32]byte{}, err
	}

	return txHash, nil
}
