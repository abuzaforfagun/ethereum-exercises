package cmd

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func listWalletsCmd(db *sql.DB) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list created wallets",
		Run: func(cmd *cobra.Command, args []string) {
			displayWallets(db)
		},
	}
}

func displayWallets(db *sql.DB) {
	fmt.Println("Name \t Address")
	sql := "SELECT name, address FROM wallet"

	rows, err := db.Query(sql)
	if err != nil {
		log.Fatalf("unable to query the database %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var name, address string

		rows.Scan(&name, &address)
		fmt.Printf("%s \t %s", name, address)
	}
}
