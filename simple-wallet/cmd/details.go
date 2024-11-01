package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func detailsWalletCmd(db *sql.DB) *cobra.Command {
	return &cobra.Command{
		Use:   "details [wallet name]",
		Args:  cobra.ExactArgs(1),
		Short: "display the wallet details",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Please type the wallet name")
				cmd.Help()
				os.Exit(1)
			}

			walletName := args[0]
			name, address, private_key := details(db, walletName)
			fmt.Println("Name: \t\t\t", name)
			fmt.Println("Address: \t\t", address)
			fmt.Println("Encrypted Private Key: \t", string(private_key))
		},
	}
}

func details(db *sql.DB, walletName string) (string, string, []byte) {
	walletName = strings.ToLower(walletName)
	sql := "SELECT name, address, private_key FROM wallet WHERE lower(name) = ?"
	row := db.QueryRow(sql, walletName)

	var name, address string
	var private_key []byte

	row.Scan(&name, &address, &private_key)

	return name, address, private_key

}
