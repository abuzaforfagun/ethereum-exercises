package cmd

import "github.com/spf13/cobra"

var createWalletCmd = &cobra.Command{
	Use:   "create [wallet name]",
	Args:  cobra.ExactArgs(1),
	Short: "create a new wallet",
	Long:  "create a new wallet and return the wallet address",
	Run:   func(cmd *cobra.Command, args []string) {},
}
