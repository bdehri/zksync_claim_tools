package main

import (
	"github.com/rs/zerolog/log"

	"github.com/bdehri/zksync_claim_tools/pkg/cmd"
)

// Connect to zkSync network
func main() {

	rootCommand := cmd.NewRootCmd()
	if err := rootCommand.Execute(); err != nil {
		log.Fatal().Err(err).Msg("Failed to execute command")
	}

}
