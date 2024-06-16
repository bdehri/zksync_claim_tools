package utils

import (
	"context"
	"math/big"

	"github.com/bdehri/zksync_claim_tools/pkg/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"github.com/zksync-sdk/zksync2-go/clients"
)

func Claim(distributor *contracts.ZkSyncMerkleDistributor, client clients.Client, opts *bind.TransactOpts, merkleIndex *big.Int, tokenAmount *big.Int, proof [][32]byte) (bool, error) {
	tx, err := distributor.Claim(opts, merkleIndex, tokenAmount, proof)
	if err != nil {
		return false, err
	}

	log.Info().Str("tx_hash", tx.Hash().Hex()).Msg("Sent claim")

	receipt, err := client.WaitMined(context.Background(), tx.Hash())
	if err != nil {
		return false, err
	}
	log.Info().Str("tx_hash", receipt.TxHash.Hex()).Msg("Tx mined")

	return true, nil
}

func Approve(token *contracts.Erc20, client clients.Client, opts *bind.TransactOpts, targetAddress common.Address, amount *big.Int) (bool, error) {
	tx, err := token.Approve(opts, targetAddress, amount)
	if err != nil {
		return false, err
	}

	log.Info().Str("tx_hash", tx.Hash().Hex()).Msg("Sent approve")

	receipt, err := client.WaitMined(context.Background(), tx.Hash())
	if err != nil {
		return false, err
	}
	log.Info().Str("tx_hash", receipt.TxHash.Hex()).Msg("Tx mined")

	return true, nil
}

func Transfer(token *contracts.Erc20, client clients.Client, opts *bind.TransactOpts, targetAddress common.Address, amount *big.Int) (bool, error) {
	tx, err := token.Transfer(opts, targetAddress, amount)
	if err != nil {
		return false, err
	}

	log.Info().Str("tx_hash", tx.Hash().Hex()).Msg("Sent transfer")

	receipt, err := client.WaitMined(context.Background(), tx.Hash())
	if err != nil {
		return false, err
	}
	log.Info().Str("tx_hash", receipt.TxHash.Hex()).Msg("Tx mined")

	return true, nil
}
