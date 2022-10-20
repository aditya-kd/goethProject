// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/ethclient"
// )
// func main(){
// 	conn, err := ethclient.Dial("http://127.0.0.1:7545")

// 	if err != nil{
// 		log.Fatal("Failed to connect to Eth Node", err)
// 	}
// 	account := common.HexToAddress("0x69c6de81fc1f1e31f1b3c775119a4ac9b7b83a3b3765ce6ff3b2fc6f1642ddd9")

// 	balance, err := conn.BalanceAt(context.Background(), account, nil)
// 	if err !=  nil {
// 		log.Fatal("Unable to get balance")
// 	}else{
// 		fmt.Println(balance)
// 	}
// }