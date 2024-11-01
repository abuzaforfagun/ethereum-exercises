package cmd

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

var createWalletCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "create a new wallet",
	Long:  "create a new wallet and return the wallet address",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error 'name' argument is required")
			cmd.Help()
			os.Exit(1)
		}
		walletName := args[0]
		address := createWallet(walletName)
		fmt.Println("Account created! Address: ", address)
	},
}

func createWallet(name string) string {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		fmt.Println("Unable to generate the key")
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	return address.Hex()
}
