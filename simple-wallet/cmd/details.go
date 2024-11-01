package cmd

import "github.com/spf13/cobra"

var displayWalletCmd = &cobra.Command{
	Use:   "details [wallet name]",
	Args:  cobra.ExactArgs(1),
	Short: "display the wallet details",
	Run:   func(cmd *cobra.Command, args []string) {},
}
