package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var (
	distributorAddress = common.HexToAddress("0x66Fd4FC8FA52c9bec2AbA368047A0b27e24ecfe4")
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
