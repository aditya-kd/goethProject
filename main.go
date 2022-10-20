// // Generate Ethereum Wallet 
// package main

// import (
// 	// "context"
// 	"fmt"
// 	"log"

// 	// "github.com/ethereum/go-ethereum/common"
// 	// "github.com/ethereum/go-ethereum/ethclient"
// 	"github.com/ethereum/go-ethereum/common/hexutil"
// 	"github.com/ethereum/go-ethereum/crypto"
// )

// func main() {

// 	privateKey, err := crypto.GenerateKey()

// 	if err != nil {
// 		log.Fatal("Unable to generate new keys")
// 	}

// 	privateKeyBytes := crypto.FromECDSA(privateKey)
// 	fmt.Println(hexutil.Encode(privateKeyBytes))
// }
