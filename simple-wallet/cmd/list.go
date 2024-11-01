package cmd

import "github.com/spf13/cobra"

var listWalletsCmd = &cobra.Command{
	Use:   "lsit",
	Short: "lsit created wallets",
	Run:   func(cmd *cobra.Command, args []string) {},
}
