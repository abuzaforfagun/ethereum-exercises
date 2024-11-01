package cmd

import "github.com/spf13/cobra"

var exportWalletsCmd = &cobra.Command{
	Use:   "export [wallet name] [filename]",
	Args:  cobra.ExactArgs(2),
	Short: "export a wallet to file",
	Run:   func(cmd *cobra.Command, args []string) {},
}
