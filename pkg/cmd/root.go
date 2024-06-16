package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var (
	distributorAddress = common.HexToAddress("0x903fA9b6339B52FB351b1319c8C0411C044422dF")
	tokenAddress       = common.HexToAddress("0x5A7d6b2F92C77FAD6CCaBd7EE0624E64907Eaf3E")
	EthereumProvider   = "https://ethereum-rpc.publicnode.com" // Sepolia testnet
	TargetAddress      common.Address
)

func NewRootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "zksync_claim_tools",
		Short: "A tool to claim zkSync tokens and do erc20 calls",
	}
	return rootCmd
}
