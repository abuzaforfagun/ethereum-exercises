package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/abuzaforfagun/ethereum-exercises/simple-wallet/cryptography"
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
			address, err := createWallet(db, walletName)
			if err != nil {
				fmt.Println("Unable to create wallet.", err)
				return
			}
			fmt.Println("Wallet created! Address: ", address)
		},
	}
}

func createWallet(db *sql.DB, name string) (string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", errors.New("unable to get encryption private key")
	}

	encryptedPrivateKey, err := cryptography.Encrypt(string(crypto.FromECDSA(privateKey)))

	if err != nil {
		return "", err
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	addressHex := address.Hex()

	sql := "INSERT INTO wallet (name, address, private_key) VALUES (?, ?, ?)"

	_, err = db.Exec(sql, name, addressHex, encryptedPrivateKey)

	if err != nil {
		return "", errors.New("unable to store the generated key")
	}
	return addressHex, nil
}
