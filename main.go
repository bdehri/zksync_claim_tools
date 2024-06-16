package main

import (
	"context"
	"flag"
	"math/big"
	"os"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	"github.com/zksync-sdk/zksync2-go/clients"
	"gopkg.in/yaml.v2"

	"github.com/bdehri/zksync_claim_tools/pkg/contracts"
	"github.com/bdehri/zksync_claim_tools/pkg/utils"
)

type Config struct {
	Wallets       []Wallet `yaml:"wallets"`
	TargetAddress string   `yaml:"target_address"`
	RpcUrl        string   `yaml:"rpc_url"`
	GatherAll     bool     `yaml:"gather_all"`
	Approve       bool     `yaml:"approve"`
	OnlyApprove   bool     `yaml:"only_approve"`
}

type Wallet struct {
	PrivateKey   string   `yaml:"private_key"`
	MerkleIndex  string   `yaml:"merkle_index"`
	MerkleProofs []string `yaml:"merkle_proofs"`
	TokenAmount  string   `yaml:"token_amount"`
}

var (
	distributorAddress = common.HexToAddress("0x903fA9b6339B52FB351b1319c8C0411C044422dF")
	tokenAddress       = common.HexToAddress("0x5A7d6b2F92C77FAD6CCaBd7EE0624E64907Eaf3E")
	EthereumProvider   = "https://rpc.ankr.com/eth_sepolia" // Sepolia testnet
	TargetAddress      common.Address
)

// Connect to zkSync network

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "/usr/src/multi-bot/config.yaml", "path to the YAML config file")
	flag.Parse()

	yamlData, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read config file")
	}

	// Parse YAML data into config struct
	var config Config
	err = yaml.Unmarshal(yamlData, &config)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse config file")
	}

	TargetAddress = common.HexToAddress(config.TargetAddress)

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

	waitgroup := sync.WaitGroup{}

	for _, wallet := range config.Wallets {
		waitgroup.Add(1)
		go func(wallet Wallet) {
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

			opts.GasTipCap = big.NewInt(300)

			merkleIndex := new(big.Int)
			merkleIndex.SetString(wallet.MerkleIndex, 10) // 10 is the base for decimal numbers

			tokenAmount := new(big.Int)
			tokenAmount.SetString(wallet.TokenAmount, 10)

			if config.OnlyApprove {
				_, err := utils.Approve(zkToken, client, opts, TargetAddress, tokenAmount)
				if err != nil {
					log.Error().Err(err).Msg("Failed to approve tokens")
					return
				}
				return
			}

			_, err = utils.Claim(distributor, client, opts, merkleIndex, tokenAmount, proof)
			if err != nil {
				log.Error().Err(err).Msg("Failed to claim tokens")
				return
			}

			if config.GatherAll {
				_, err := utils.Transfer(zkToken, client, opts, TargetAddress, tokenAmount)
				if err != nil {
					log.Error().Err(err).Msg("Failed to transfer tokens")
					return
				}
			}

			if config.Approve {
				_, err := utils.Approve(zkToken, client, opts, TargetAddress, tokenAmount)
				if err != nil {
					log.Error().Err(err).Msg("Failed to approve tokens")
					return
				}
			}
		}(wallet)
	}
	waitgroup.Wait()

	log.Info().Msg("All done")

}
