package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

var (
	infuraUrlBase string = "wss://mainnet.infura.io/ws/v3/"
)

func main() {
	smartContractAddress := "0x514910771af9ca656af840dff83e8264ecf986ca" // Chainlink token
	logEvents(smartContractAddress)
}

func logEvents(address string) {
	infuraKey := os.Getenv("infuraKey")

	if infuraKey == "" {
		log.Panic("Please add 'infuraKey' to the enviornment variable")
	}

	url := infuraUrlBase + infuraKey

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Panic("Unable to connect ehtereum", err)
	}

	headerCh := make(chan *types.Header)
	client.SubscribeNewHead(context.Background(), headerCh)
	linkAddress := common.HexToAddress(address)

	logCh := make(chan types.Log)
	eventSignature := "Transfer(address,address,uint256)"

	eventHash := crypto.Keccak256Hash([]byte(eventSignature))

	client.SubscribeFilterLogs(context.Background(), ethereum.FilterQuery{
		Addresses: []common.Address{
			linkAddress,
		},
		Topics: [][]common.Hash{
			[]common.Hash{
				eventHash,
			},
		},
	}, logCh)

	for logItem := range logCh {
		if len(logItem.Topics) > 0 && logItem.Topics[0] != eventHash {
			continue
		}
		fmt.Println("Block#: ", logItem.BlockNumber)
		fmt.Println("Block Hash: ", logItem.BlockHash)

		block, err := client.BlockByHash(context.Background(), logItem.BlockHash)
		if err != nil {
			log.Println("unable to get block", err)
			continue
		}
		fmt.Println("Time: ", time.Unix(int64(block.Time()), 0))

		fmt.Println("Transaction Hash: ", logItem.TxHash)

		tx, _, err := client.TransactionByHash(context.Background(), logItem.TxHash)
		if err != nil {
			log.Println("unable to get the transaction")
			continue
		}

		sender, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
		if err != nil {
			log.Println("unable to get the signer")
		}
		fmt.Println("From: ", sender.Hex())
		fmt.Println("To: ", tx.To().Hex())
		amount := new(big.Int)
		amount.SetBytes(logItem.Data)
		fmt.Println("Amount: ", weiToEther(amount))
	}
}

func weiToEther(wei *big.Int) *big.Float {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)
	return f.Quo(fWei.SetInt(wei), big.NewFloat(params.Ether))
}
