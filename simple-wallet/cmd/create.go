package cmd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

func createWalletCmd(db *sql.DB) *cobra.Command {
	return &cobra.Command{
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
			address := createWallet(db, walletName)
			fmt.Println("Account created! Address: ", address)
		},
	}
}

func createWallet(db *sql.DB, name string) string {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		fmt.Println("Unable to generate the key")
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	addressHex := address.Hex()

	sql := "INSERT INTO wallet (name, address, private_key) VALUES (?, ?, ?)"

	_, err = db.Exec(sql, name, addressHex, privateKey.D.Bytes())

	if err != nil {
		fmt.Println("Unable to store the generated key")
		return ""
	}
	return addressHex
}
