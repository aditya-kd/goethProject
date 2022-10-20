// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"math/big"

// 	// "github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/ethclient"
// )

// func main() {

// 	client, err := ethclient.Dial("http://127.0.0.1:7545")

// 	if err != nil {
// 		log.Fatal("Unable to reach client")
// 	}

// 	header, err := client.HeaderByNumber(context.Background(), nil)

// 	if err != nil {
// 		log.Fatal("Unable to get block header")
// 	}

// 	fmt.Println(header.Number.String())
// 	blockNumber := big.NewInt(1)
// 	block, err := client.BlockByNumber(context.Background(), blockNumber)

// 	if err != nil {
// 		log.Fatal("Unable to get block by number", blockNumber, err)
// 	}

// 	fmt.Println(block.Number().Uint64())
// 	fmt.Println(block.Time())
// 	fmt.Println(block.Difficulty())
// 	fmt.Println(block.Hash().Hex())

// 	for _, tx := range block.Transactions() {
// 		fmt.Println(tx.Hash().Hex())
// 	}
// }
