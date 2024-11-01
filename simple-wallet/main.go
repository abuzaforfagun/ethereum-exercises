package main

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/abuzaforfagun/ethereum-exercises/simple-wallet/cmd"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := initDatabase()
	verifySecretKey()
	defer db.Close()
	cmd.Execute(db)
}

func verifySecretKey() {
	secretKey := os.Getenv("wallet_secret_key")
	if secretKey == "" {
		key := make([]byte, 16)

		_, err := rand.Read(key)
		if err != nil {
			log.Fatal(err)
		}
		hexKey := hex.EncodeToString(key)
		os.Setenv("wallet_secret_key", string(hexKey))
		fmt.Printf("A random secret is added to the enviornment\n", key)
	}
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
