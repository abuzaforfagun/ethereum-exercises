package cmd

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/abuzaforfagun/ethereum-exercises/simple-wallet/cryptography"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

func exportWalletsCmd(db *sql.DB) *cobra.Command {

	return &cobra.Command{
		Use:   "export [wallet name] [filename]",
		Args:  cobra.ExactArgs(2),
		Short: "export a wallet to file",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				fmt.Println("Error 'name' and 'file name' argument is required")
				cmd.Help()
				os.Exit(1)
			}

			walletName := args[0]
			fileName := args[1]
			err := export(db, walletName, fileName)
			if err != nil {
				fmt.Println("Unable to export. %v", err)
				os.Exit(1)
			}

			fmt.Println("Private key is exported")
		},
	}
}

func export(db *sql.DB, walletName string, filename string) error {
	walletName = strings.ToLower(walletName)
	sql := "SELECT name, address, private_key FROM wallet WHERE lower(name) = ?"
	row := db.QueryRow(sql, walletName)

	var name, address string
	var private_key []byte

	row.Scan(&name, &address, &private_key)

	decryptedPrivateKey, err := cryptography.Decrypt(string(private_key))
	if err != nil {
		return err
	}

	privateKey, err := crypto.ToECDSA([]byte(decryptedPrivateKey))
	addrs := crypto.PubkeyToAddress(privateKey.PublicKey)
	fmt.Println(addrs.Hex())
	if err != nil {
		return err
	}
	_ = privateKey

	privateKeyBytes := crypto.FromECDSA(privateKey)

	privateKeyHex := hex.EncodeToString(privateKeyBytes)

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(privateKeyHex)
	if err != nil {
		return err
	}

	return nil
}
