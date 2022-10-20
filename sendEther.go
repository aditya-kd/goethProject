package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func makeTransaction() {
	client, err := ethclient.Dial("http://127.0.0.1:7545")

	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("466022db4e0e2a754837dd8d2f58de12331561ebb92d1e414ad5a3e9d81f2931")

	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error caasting public key to eECCDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	value := big.NewInt(3000000000000000000)
	gasLimit := uint64(210000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x8775BEc3b4ce3a082B7CBdEC5998016e37dC8fd2")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sent", value.String(), "wei")
	fmt.Println("To Address: ", toAddress.Hex())
	fmt.Println("Created Block Hash: ", signedTx.Hash())
}
func main() {

	var transactionCount = 15
	start := time.Now()
	for i := 1; i <= transactionCount; i++ {
		fmt.Println("\n")
		fmt.Println("Transaction ", i)
		makeTransaction()
	}
	fmt.Println("\n\n")
	fmt.Println("Total Transactions: ", transactionCount)
	log.Printf("Transaction Execution time: %s\n", time.Since(start))

}
