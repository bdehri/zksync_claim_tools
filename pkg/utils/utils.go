package utils

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"github.com/zksync-sdk/zksync2-go/clients"
	zksyncTypes "github.com/zksync-sdk/zksync2-go/types"

	"github.com/bdehri/zksync_claim_tools/pkg/contracts"
)

func Claim(distributor *contracts.ZkSyncMerkleDistributor, client clients.Client, opts *bind.TransactOpts, merkleIndex *big.Int, tokenAmount *big.Int, proof [][32]byte) (bool, error) {
	i := 0
	result := false
	var err error
	var tx *types.Transaction
	var receipt *zksyncTypes.Receipt
	for i < 5 && !result {
		i++

		tx, err = distributor.Claim(opts, merkleIndex, tokenAmount, proof)
		if err != nil {
			log.Error().Err(err).Msg("Failed to send claim, retrying...")
			continue
		}

		log.Info().Str("tx_hash", tx.Hash().Hex()).Msg("Sent claim")

		receipt, err = client.WaitMined(context.Background(), tx.Hash())
		if err != nil {
			log.Error().Err(err).Msg("Failed to wait for tx, retrying...")
			continue
		}
		log.Info().Str("tx_hash", receipt.TxHash.Hex()).Msg("Tx mined")
		result = true

	}
	return result, nil
}

func Approve(token *contracts.Erc20, client clients.Client, opts *bind.TransactOpts, targetAddress common.Address, amount *big.Int) (bool, error) {
	i := 0
	result := false
	var err error
	var tx *types.Transaction
	var receipt *zksyncTypes.Receipt
	for i < 5 && !result {
		i++

		tx, err = token.Approve(opts, targetAddress, amount)
		if err != nil {
			log.Error().Err(err).Msg("Failed to send approve, retrying...")
			continue
		}

		log.Info().Str("tx_hash", tx.Hash().Hex()).Msg("Sent approve")

		receipt, err = client.WaitMined(context.Background(), tx.Hash())
		if err != nil {
			log.Error().Err(err).Msg("Failed to wait for tx, retrying...")
			continue
		}
		log.Info().Str("tx_hash", receipt.TxHash.Hex()).Msg("Tx mined")
		break
	}
	return result, nil
}

func Transfer(token *contracts.Erc20, client clients.Client, opts *bind.TransactOpts, targetAddress common.Address, amount *big.Int) (bool, error) {
	i := 0
	result := false
	var err error
	var tx *types.Transaction
	var receipt *zksyncTypes.Receipt
	for i < 5 && !result {
		i++
		tx, err = token.Transfer(opts, targetAddress, amount)
		if err != nil {
			log.Error().Err(err).Msg("Failed to send transfer, retrying...")
			continue
		}

		log.Info().Str("tx_hash", tx.Hash().Hex()).Msg("Sent transfer")

		receipt, err = client.WaitMined(context.Background(), tx.Hash())
		if err != nil {
			log.Error().Err(err).Msg("Failed to wait for tx, retrying...")
			continue
		}
		log.Info().Str("tx_hash", receipt.TxHash.Hex()).Msg("Tx mined")
		break
	}
	return result, err
}
