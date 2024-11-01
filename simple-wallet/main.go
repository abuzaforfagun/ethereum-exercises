package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/abuzaforfagun/ethereum-exercises/simple-wallet/cmd"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := initDatabase()
	defer db.Close()
	cmd.Execute(db)
}

func initDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./wallet.db")
	if err != nil {
		log.Fatalf("unable to open database %v", err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS wallet (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"address" TEXT,
		"private_key" BLOB
	);`

	_, err = db.ExecContext(context.Background(), createTableSQL)
	if err != nil {
		log.Fatalf("unable to create table")
	}

	return db
}
