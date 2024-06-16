package cmd

import (
	"context"
	"math/big"
	"os"
	"sync"

	"github.com/bdehri/zksync_claim_tools/pkg/contracts"
	"github.com/bdehri/zksync_claim_tools/pkg/interfaces"
	"github.com/bdehri/zksync_claim_tools/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/zksync-sdk/zksync2-go/clients"
	"gopkg.in/yaml.v2"
)

func NewApproveCmd() *cobra.Command {
	var configPath, targetAddress string
	var approveCmd = &cobra.Command{
		Use:   "approve",
		Short: "approve zkSync tokens to an address",
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

			zkToken, err := contracts.NewErc20(tokenAddress, client)
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to create zkToken contract")
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

					// Start configuring transaction parameters
					opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
					if err != nil {
						log.Error().Err(err).Msg("Failed to create transactor")
					}

					opts.GasTipCap = big.NewInt(300)

					merkleIndex := new(big.Int)
					merkleIndex.SetString(wallet.MerkleIndex, 10) // 10 is the base for decimal numbers

					tokenAmount := new(big.Int)
					tokenAmount.SetString(wallet.TokenAmount, 10)

					_, err = utils.Approve(zkToken, client, opts, TargetAddress, tokenAmount)
					if err != nil {
						log.Error().Err(err).Msg("Failed to approve tokens")
						return
					}
				}(wallet)
			}
			waitgroup.Wait()
		},
	}

	approveCmd.Flags().StringVar(&configPath, "config", "onfig.yaml", "path to the YAML config file")
	approveCmd.Flags().StringVar(&targetAddress, "target", "", "target address to approve tokens to")
	approveCmd.MarkFlagRequired("target")
	approveCmd.MarkFlagRequired("config")
	return approveCmd
}
