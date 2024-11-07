package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

var (
	infuraUrlBase           string = "wss://mainnet.infura.io/ws/v3/"
	gasForSingleTransaction int    = 21000
)

type Response struct {
	EstimatedCostInEth string  `json:"estimated_cost"`
	EstimatedCostInUSD float64 `json:"estimated_cost_in_usd"`
}

type CoinGeckoResponse struct {
	Ethereum CurrencyResponse `json:"ethereum"`
}

type CurrencyResponse struct {
	USD float64 `json:"usd"`
}

func main() {
	infuraKey := os.Getenv("infuraKey") //4cdbd4667d854012af3a26102d205242
	infuraKey = "4cdbd4667d854012af3a26102d205242"

	if infuraKey == "" {
		log.Panic("Please add 'infuraKey' to the enviornment variable")
	}

	router := gin.Default()

	url := infuraUrlBase + infuraKey

	ethClient, err := ethclient.Dial(url)
	if err != nil {
		log.Panic("Unable to connect ehtereum", err)
	}

	router.GET("/single-transaction", func(ctx *gin.Context) {
		estimatedEth, err := getEstimate(context.Background(), ethClient, gasForSingleTransaction)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		costInUsd, err := getPriceInUsd(context.Background(), estimatedEth)
		costInUsdPlain, _ := costInUsd.Float64()

		response := &Response{
			EstimatedCostInEth: fmt.Sprintf("%.10f", estimatedEth),
			EstimatedCostInUSD: costInUsdPlain,
		}

		ctx.JSON(http.StatusOK, response)
	})

	router.GET("manual/:unit", func(ctx *gin.Context) {
		unitOfGasParam := ctx.Param("unit")

		unitOfGas, err := strconv.ParseInt(unitOfGasParam, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Please provide `unit`"})
			return
		}

		if unitOfGas < 1 {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Please provide `unit`"})
			return
		}

		estimatedEth, err := getEstimate(context.Background(), ethClient, int(unitOfGas))

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		cost := fmt.Sprintf("%.10f", estimatedEth)

		costInUsd, err := getPriceInUsd(context.Background(), estimatedEth)
		costInUsdPlain, _ := costInUsd.Float64()

		response := &Response{
			EstimatedCostInEth: cost,
			EstimatedCostInUSD: costInUsdPlain,
		}

		ctx.JSON(http.StatusOK, response)
	})

	log.Println("running on port 8080")
	router.Run(":8080")

}

func getEstimate(ctx context.Context, client *ethclient.Client, numberOfGasUnit int) (*big.Float, error) {
	suggestedGaspriceInWei, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}
	suggestedGasTip, err := client.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, err
	}

	totalEstimatedGasPriceInWei := new(big.Int).Add(suggestedGasTip, suggestedGaspriceInWei)
	gasPriceEth := new(big.Float).Quo(new(big.Float).SetInt(totalEstimatedGasPriceInWei), big.NewFloat(1e18))
	totalGasPrice := new(big.Float).Mul(gasPriceEth, new(big.Float).SetInt64(int64(numberOfGasUnit)))
	return totalGasPrice, nil
}

func getPriceInUsd(ctx context.Context, eth *big.Float) (*big.Float, error) {
	api := "https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd"
	response, err := http.Get(api)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("unable to get price")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var responsePayload CoinGeckoResponse

	json.Unmarshal(body, &responsePayload)
	price := new(big.Float).Mul(eth, big.NewFloat(float64(responsePayload.Ethereum.USD)))
	return price, nil
}
