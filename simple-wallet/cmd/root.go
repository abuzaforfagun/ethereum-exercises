package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "simple-wallet",
	Short: "simple-wallet is a cli tool to play with ethereum wallet",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	rootCommand.AddCommand(createWalletCmd)
	rootCommand.AddCommand(displayWalletCmd)
	rootCommand.AddCommand(listWalletsCmd)
	rootCommand.AddCommand(removeWalletCmd)
	rootCommand.AddCommand(exportWalletsCmd)

	if err := rootCommand.Execute(); err != nil {
		log.Panicf("Unable to execute command %v", err)
	}
}
