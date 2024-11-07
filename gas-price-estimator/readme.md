### Project Overview

_Project Name_: Gas price estimator

_Objective_: Develop an API to retrieves the current gas price from an Ethereum node and calculates the estimated cost of a transaction. This can be useful for users or developers who want a quick way to estimate the transaction cost on Ethereum, based on current gas prices and user-specified parameters.

## How to run

1. Retrieve `infura` API key, [see more](https://docs.infura.io/dashboard/create-api)
2. Set the value as envrionrment variable `infuraKey`
3. `cd gas-price-estimator`
4. From terminal `go run .`

## API Details

1. `/single-transaction` returns the estimated cost of sending a single transaction
2. `/manual/:unit` here `unit` is the number of gas it require to do the operations. And it returns the estimated cost of that operation.
