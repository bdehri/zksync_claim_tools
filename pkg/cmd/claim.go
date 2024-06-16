package cmd

import (
	"context"
	"math/big"
	"os"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/zksync-sdk/zksync2-go/clients"
	"gopkg.in/yaml.v2"

	"github.com/bdehri/zksync_claim_tools/pkg/contracts"
	"github.com/bdehri/zksync_claim_tools/pkg/interfaces"
	"github.com/bdehri/zksync_claim_tools/pkg/utils"
)

func NewClaimCmd() *cobra.Command {
	var approve, gather bool
	var targetAddress string
	var configPath string
	var gasPrice int64

	var claimCmd = &cobra.Command{
		Use:   "claim",
		Short: "claim zkSync tokens",
		Run: func(cmd *cobra.Command, args []string) {
			yamlData, err := os.ReadFile(configPath)
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read config file")
			}

			// Parse YAML data into config struct
			var config interfaces.Config
			err = yaml.Unmarshal(yamlData, &config)
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to parse config file")
			}

			TargetAddress = common.HexToAddress(targetAddress)

			client, err := clients.Dial(config.RpcUrl)
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to connect to zkSync provider")
				return
			}
			defer client.Close()

			chainID, err := client.ChainID(context.TODO())
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to get chain ID")
				return
			}

			// Connect to Ethereum network
			ethClient, err := ethclient.Dial(EthereumProvider)
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to connect to Ethereum provider")
			}
			defer ethClient.Close()

			distributor, err := contracts.NewZkSyncMerkleDistributor(distributorAddress, client)
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to create distributor contract")
			}

			zkToken, err := contracts.NewErc20(tokenAddress, client)
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to create zkToken contract")
			}

			windowStart, err := distributor.WINDOWSTART(nil)
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to get window start")
			}

			for {
				if windowStart.Int64()-time.Now().Unix() <= 2 {
					break
				}
				log.Info().Int64("window_start", windowStart.Int64()).Msg("Waiting for window start minus 2")
				time.Sleep(1 * time.Second)
			}

			waitgroup := sync.WaitGroup{}

			for _, wallet := range config.Wallets {
				waitgroup.Add(1)
				go func(wallet interfaces.Wallet) {
					defer waitgroup.Done()
					privateKey, err := crypto.ToECDSA(common.Hex2Bytes(wallet.PrivateKey))
					if err != nil {
						log.Error().Err(err).Msg("Failed to parse private key")
						return
					}

					proof := make([][32]byte, 0)
					for _, p := range wallet.MerkleProofs {
						proof = append(proof, [32]byte(common.Hex2BytesFixed(p, 32)))
					}

					// Start configuring transaction parameters
					opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
					if err != nil {
						log.Error().Err(err).Msg("Failed to create transactor")
					}

					opts.GasTipCap = big.NewInt(1000000000)
					opts.GasFeeCap = big.NewInt(11000000000)
					if gasPrice > 0 {
						opts.GasPrice = big.NewInt(gasPrice * 1e9)
					}

					merkleIndex := new(big.Int)
					merkleIndex.SetString(wallet.MerkleIndex, 10) // 10 is the base for decimal numbers

					tokenAmount := new(big.Int)
					tokenAmount.SetString(wallet.TokenAmount, 10)

					_, err = utils.Claim(distributor, client, opts, merkleIndex, tokenAmount, proof)
					if err != nil {
						log.Error().Err(err).Msg("Failed to claim tokens")
						return
					}

					if gather {
						_, err := utils.Transfer(zkToken, client, opts, TargetAddress, tokenAmount)
						if err != nil {
							log.Error().Err(err).Msg("Failed to transfer tokens")
							return
						}
					}

					if approve {
						_, err := utils.Approve(zkToken, client, opts, TargetAddress, tokenAmount)
						if err != nil {
							log.Error().Err(err).Msg("Failed to approve tokens")
							return
						}
					}
				}(wallet)
			}
			waitgroup.Wait()
		},
	}

	claimCmd.Flags().BoolVar(&approve, "approve", false, "approve tokens after claiming")
	claimCmd.Flags().BoolVar(&gather, "gather", false, "gather all tokens to single account")
	claimCmd.Flags().StringVar(&targetAddress, "target", "", "target address to send tokens to")
	claimCmd.Flags().StringVar(&configPath, "config", "config.yaml", "path to config file")
	claimCmd.Flags().Int64Var(&gasPrice, "gas-price", 0, "gas price for transactions in gwei")
	claimCmd.MarkFlagRequired("config")
	return claimCmd
}
