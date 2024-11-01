package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func removeWalletCmd(db *sql.DB) *cobra.Command {
	return &cobra.Command{
		Use:   "remove [wallet name]",
		Args:  cobra.ExactArgs(1),
		Short: "remove wallet",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Please provide the wallet name")
				cmd.Help()
				os.Exit(1)
			}
			walletName := args[0]

			err := remove(db, walletName)
			if err != nil {
				fmt.Printf("Unable to delete the wallet. %v", err)
			}
			fmt.Printf("Wallet [%s] deleted sucessfully!\n", walletName)
		},
	}
}

func remove(db *sql.DB, name string) error {
	name = strings.ToLower(name)

	var exists bool
	getWalletSql := "SELECT 1 FROM wallet WHERE name = ?"

	row := db.QueryRow(getWalletSql, name)
	row.Scan(&exists)
	if !exists {
		return errors.New("wallet does not exist")
	}

	sql := "DELETE FROM wallet WHERE name = ?"

	_, err := db.Exec(sql, name)

	return err
}
