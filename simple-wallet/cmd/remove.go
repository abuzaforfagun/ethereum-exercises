package cmd

import "github.com/spf13/cobra"

var removeWalletCmd = &cobra.Command{
	Use:   "remove [wallet name]",
	Args:  cobra.ExactArgs(1),
	Short: "remove wallet",
	Run:   func(cmd *cobra.Command, args []string) {},
}
