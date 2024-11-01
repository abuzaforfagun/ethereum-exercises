package cmd

import (
	"database/sql"
	"log"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "simple-wallet",
	Short: "simple-wallet is a cli tool to play with ethereum wallet",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute(db *sql.DB) {
	rootCommand.AddCommand(createWalletCmd(db))
	rootCommand.AddCommand(detailsWalletCmd(db))
	rootCommand.AddCommand(listWalletsCmd(db))
	rootCommand.AddCommand(removeWalletCmd(db))
	rootCommand.AddCommand(exportWalletsCmd(db))

	if err := rootCommand.Execute(); err != nil {
		log.Panicf("Unable to execute command %v", err)
	}
}
