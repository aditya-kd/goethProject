// // Connect to Ethereum
// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/ethclient"
// )

// func main() {
// 	fmt.Println("Working")
// 	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/f82b411fd41040bc90c1e9f38c3961c3")

// 	if err != nil {
// 		log.Fatal("Failed to connect to Eth Node", err)
// 	}
// 	cntxt := context.Background()
// 	txn, pending, _ := conn.TransactionByHash(cntxt, common.HexToHash("0xdec8a1daf5cd4b8f2ec4f41373922a57121b72ac5d4b039a6252add7107eb217"))

// 	if pending {
// 		fmt.Println("Txn is pending", txn)
// 	} else {
// 		fmt.Println("Txn is not pending", txn)
// 	}
// }
