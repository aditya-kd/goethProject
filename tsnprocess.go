// package main

// import(
// 	"fmt"
// 	"log"
// 	"context"
// 	"crypto/ecdsa"
// 	"math/big"

// 	"github.com/ethereum/go-ethereum/ethclient"
// 	"github.com/ethereum/go-ethereum/core/types"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/crypto"
// 	"github.com/ethereum/go-ethereum/params"
// 	"github.com/ethereum/go-ethereum/rlp"
// )
// func main()
// {
// 	client, err := ethclient.Dial("https://mainnet.infura.io/v3/f82b411fd41040bc90c1e9f38c3961c3")
// 	fmt.Println("This is the start of this Program")
// 	client, err := nil{
// 		log.Fatal("Unable to connect")
// 	}
// 	privateKey, err := crypto.HexToECDSA("Hash")

// 	if err != nil{
// 		log.Fatal(err)
// 	}
// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

// 	if !ok{
// 		log.Fatal("Unalble to cast public key to ECDSA")
// 	}

// 	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
// 	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)

// 	if err != nil{
// 		log.Fatal("Unable to get nonce")
// 	}
// 	fmt.Println("nonce", nonce)
// 	toAddress := common.HexToAddress("Hash")
// 	value := big.NewInt(100)

// 	gasFeeCap, gasTip, gasLimit := big.NewInt(233), big.NewInt(234), big.NewInt(324)

// 	var data []byte
// 	tx := types.NewTx(&types.DynamicFeeTx{
// 		Nonce: nonce,
// 		GasFeeCap: gasFeeCap,
// 		GasTipCap: gasTipCap,
// 		Gas: gas,
// 		To: &toAddress,
// 		Value: value,
// 		Data: data
// 	}
// 	config, block := params.RinkebyChainConfig, params.RinkebyChainConfig.LondonBlock
// 	signer := types.MakeSigner(config, block)
// 	signedTx, err := types.SignTx(tx, signer, privateKey)
// 	if err != nil{
// 		log.Fatal("Unable to sign Tx")
// 	}

// 	hash := signedTx.Hash().Bytes()
// 	raw, err := rlp.EncodeToBytes(signedTx)

// 	fmt.Println("hash:%x\n0x%x", hash, raw)
// 	err = client.SendTransaction(context.Background(), signedTx)

// 	if err != nil{
// 		log.Fatal("Unable to submit transaction")
// 	}

// 	fmt.Printf("tx sent: ", signedTx.Hash().Hex())
	
// }
