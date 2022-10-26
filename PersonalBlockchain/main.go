// package main

// import(
// 	"crypto/md5"
// 	"crypto/sha256"
// 	"encoding/hex"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/gorilla/mux"
// )

// type Block struct{
// 	Pos int
// 	Data BookCheckout
// 	TimeStamp string
// 	Hash string 
// 	PrevHash string
// }

// type BookCheckout struct {
// 	BookID string		`json:"book_id"` 
// 	User string 		`json:"user"`
// 	CheckoutDate string `json:"checkout_date"`
// 	IsGenesis bool  	`json:"is_genesis"`
// }

// type Book struct{

// 	ID string		`json:"id"`
// 	Title string	`json:"title"`
// 	Author string	`json:"author"`
// 	PublishDate string	`json:"publish_date"`
// 	ISBN string			`json:"isbn"`
// }

// type BlockChain struct{
// 	blocks[]*Block
// }

// func (b *Block) generateHash(){

// 	bytes, _ := json.Marshal(b.Data)

// 	data := string(b.Pos) + b.TimeStamp + string(bytes) + b.PrevHash

// 	hash := sha256.New()
// 	hash.Write([]byte(data))
// 	b.Hash = hex.EncodeToString(hash.Sum(nil))

// }

// func CreateBlock(prevBlock *Block, checkoutItem BookCheckout ) *Block{

// 	block := &Block{}
// 	block.Pos = prevBlock.Pos + 1
// 	block.TimeStamp = time.Now().String()
// 	block.PrevHash = prevBlock.Hash
// 	block.Data = checkoutItem
// 	block.generateHash()

// 	return block
// }

// func (bc *BlockChain)AddBlock(data BookCheckout){
// 	prevBlock := bc.blocks[len(bc.blocks)-1]

// 	block := CreateBlock(prevBlock, data)

// 	if validBlock(block, prevBlock){
// 		bc.blocks = append(bc.blocks, block)
// 	}

// }

// func validBlock(block, prevBlock *Block) bool{
// 	if prevBlock.Hash != block.PrevHash{
// 		return false
// 	}

// 	if !block.validateHash(block.Hash){
// 		return false
// 	}
// 	if prevBlock.Pos+1 != block.Pos{
// 		return false
// 	}

// 	return true
// }
// func (b* Block) validateHash(hash string) bool{
// 	b.generateHash()
// 	if b.Hash != hash{
// 		return false
// 	}
// 	return true
// }
// func writeBlock(w http.ResponseWriter, r *http.Request){
// 	var checkoutitem BookCheckout;

// 	if err := json.NewDecoder(r.Body).Decode(&checkoutitem); err != nil{
// 		w.WriteHeader(http.StatusInternalServerError)
// 		log.Printf("Could not write the block %v", err)
// 		w.Write([]byte("Could not write the block"))
// 		return 
// 	}

// 	BlockChain.AddBlock(checkoutitem)

// 	resp, err := json.MarshalIndent(checkoutitem, "", " ")
// 	if err != nil{
// 		w.WriteHeader(http.StatusInternalServerError)
// 		log.Printf("Could not Marshal Payload: %v", err)
// 		w.Write([]byte("Could not write block"))
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(resp)

// }


// func newBook(w http.ResponseWriter, r *http.Request){
// 	var book Book 

// 	if err := json.NewDecoder(r.Body).Decode(&book); err != nil{
// 		w.WriteHeader(http.StatusInternalServerError)
// 		log.Printf("Could not create : %v", err)
// 		w.Write([]byte("could not create new book"))
// 		return 
// 	}

// 	h := md5.New()
// 	io.WriteString(h, book.ISBN+book.PublishDate)
// 	book.ID = fmt.Sprintf("%x", h.Sum(nil))

// 	resp, err := json.MarshalIndent(book, "", " ")
// 	if err != nil{
// 		w.WriteHeader(http.StatusInternalServerError)
// 		log.Printf("Could not Marshal payload %v", err)
// 		w.Write([]byte("Could Not save book data"))
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(resp) 

// }
// var Blockchain *Blockchain

// func NewBlokcchian() *Blockchain{
// 	return &Blockchain([]*Block{GenesisBlock()})
// }
// func GenesisBlock() *Block{
// 	return CreateBlock(&Block(), BookCheckout{IsGenesis: true})
// }

// func getBlockchain(w http.ResponseWriter, r *http.Request){
// 	jbytes, err := json.MarshalIndent(BlockChain.blocks, "", " ")
// 	if err != nil{
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(err)
// 		return
// 	}
// 	io.WriteString(w, string(jbytes))
// }
// func main(){

// 	BlokChain = NewBlockchain()

// 	r := mux.NewRouter()
// 	r.HandleFunc("/", getBlockchain).Methods("GET")
// 	r.HandleFunc("/", writeBlock).Methods("POST")
// 	r.HandleFunc("/new", newBook).Methods("POST")


// 	go func(){
// 		for _, block := range BlockChain.blocks{
// 			fmt.Printf("Prev.Hash: %x\n", block.PrevHash)
// 			bytes, _ := json.MarshalIndent(block.Data, "", " ")
// 			fmt.Printf("Data:%v\n", string(bytes))
// 			fmt.Printf("Hash:%v\n", block.Hash)
// 			fmt.Println()
// 		}
// 	}()
// 	log.Println("Listening on port 3000")
// 	log.Fatal(http.ListenAndServe(":3000", r))

// }


package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Block struct {
	Pos       int
	Data      BookCheckout
	Timestamp string
	Hash      string
	PrevHash  string
}

type BookCheckout struct {
	BookID       string `json:"book_id"`
	User         string `json:"user"`
	CheckoutDate string `json:"checkout_date"`
	IsGenesis    bool   `json:"is_genesis"`
}

type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishDate string `json:"publish_date"`
	ISBN        string `json:"isbn"`
}

func (b *Block) generateHash() {
	// get string val of the Data
	bytes, _ := json.Marshal(b.Data)
	// concatenate the dataset
	data := string(rune(b.Pos)) + b.Timestamp + string(bytes) + b.PrevHash
	hash := sha256.New()
	hash.Write([]byte(data))
	b.Hash = hex.EncodeToString(hash.Sum(nil))
}

func CreateBlock(prevBlock *Block, checkoutItem BookCheckout) *Block {
	block := &Block{}
	block.Pos = prevBlock.Pos + 1
	block.Timestamp = time.Now().String()
	block.Data = checkoutItem
	block.PrevHash = prevBlock.Hash
	block.generateHash()

	return block
}

type Blockchain struct {
	blocks []*Block
}

var BlockChain *Blockchain

func (bc *Blockchain) AddBlock(data BookCheckout) {

	prevBlock := bc.blocks[len(bc.blocks)-1]

	block := CreateBlock(prevBlock, data)

	if validBlock(block, prevBlock) {
		bc.blocks = append(bc.blocks, block)
	}
}

func GenesisBlock() *Block {
	return CreateBlock(&Block{}, BookCheckout{IsGenesis: true})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}

func validBlock(block, prevBlock *Block) bool {

	if prevBlock.Hash != block.PrevHash {
		return false
	}

	if !block.validateHash(block.Hash) {
		return false
	}

	if prevBlock.Pos+1 != block.Pos {
		return false
	}
	return true
}

func (b *Block) validateHash(hash string) bool {
	b.generateHash()
	return b.Hash == hash
}

func getBlockchain(w http.ResponseWriter, r *http.Request) {
	jbytes, err := json.MarshalIndent(BlockChain.blocks, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	io.WriteString(w, string(jbytes))
}

func writeBlock(w http.ResponseWriter, r *http.Request) {
	var checkoutItem BookCheckout
	if err := json.NewDecoder(r.Body).Decode(&checkoutItem); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not write Block: %v", err)
		w.Write([]byte("could not write block"))
		return
	}

	BlockChain.AddBlock(checkoutItem)
	resp, err := json.MarshalIndent(checkoutItem, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not marshal payload: %v", err)
		w.Write([]byte("could not write block"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func newBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not create: %v", err)
		w.Write([]byte("could not create new Book"))
		return
	}

	h := md5.New()
	io.WriteString(h, book.ISBN+book.PublishDate)
	book.ID = fmt.Sprintf("%x", h.Sum(nil))

	resp, err := json.MarshalIndent(book, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not marshal payload: %v", err)
		w.Write([]byte("could not save book data"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func main() {

	BlockChain = NewBlockchain()

	r := mux.NewRouter()
	r.HandleFunc("/", getBlockchain).Methods("GET")
	r.HandleFunc("/", writeBlock).Methods("POST")
	r.HandleFunc("/new", newBook).Methods("POST")

	go func() {

		for _, block := range BlockChain.blocks {
			fmt.Printf("Prev. hash: %x\n", block.PrevHash)
			bytes, _ := json.MarshalIndent(block.Data, "", " ")
			fmt.Printf("Data: %v\n", string(bytes))
			fmt.Printf("Hash: %x\n", block.Hash)
			fmt.Println()
		}

	}()
	log.Println("Listening on port 3000")

	log.Fatal(http.ListenAndServe(":3000", r))
}