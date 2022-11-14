package assignment02

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

var ID_Count = 0
var Nonce_val = 0

type Transaction struct {
	TransactionID string
	Sender        string
	Receiver      string
	Amount        int
}

type Block struct {
	Nonce       int
	BlockData   []Transaction
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

type Blockchain struct {
	ChainHead *Block
}

func GenerateNonce(blockData []Transaction) int {

	Nonce_val++
	return Nonce_val
}

func CalculateHash(blockData []Transaction, nonce int) string {
	dataString := ""
	for i := 0; i < len(blockData); i++ {
		dataString += (blockData[i].TransactionID + blockData[i].Sender +
			blockData[i].Receiver + strconv.Itoa(blockData[i].Amount)) + strconv.Itoa(nonce)
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(dataString)))
}

func NewBlock(blockData []Transaction, chainHead *Block) *Block {
	block := new(Block)
	block.Nonce = GenerateNonce(blockData)
	block.BlockData = blockData
	block.PrevPointer = chainHead
	block.CurrentHash = CalculateHash(block.BlockData, block.Nonce)

	if chainHead == nil {
		block.PrevHash = ""

	} else {
		block.PrevHash = chainHead.CurrentHash
	}

	return block
}

func ListBlocks(chainHead *Block) {

	fmt.Println("WHOLE BLOCKCHAIN")

	currentNode := chainHead
	if currentNode == nil {
		fmt.Println("nothing to display")
		return
	}
	fmt.Println(strings.Repeat("=", 25))
	fmt.Printf("%+v\n", *currentNode)
	for currentNode.PrevPointer != nil {
		currentNode = currentNode.PrevPointer
		fmt.Println(strings.Repeat("=", 25))
		fmt.Printf("%+v\n", *currentNode)
	}
}

func DisplayTransactions(blockData []Transaction) {
	for i, x := range blockData {
		fmt.Printf("%s Transaction Details  %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))

		fmt.Printf(" TRANSACTION ID: %s \n SENDER :  %s \n RECIEVER :  %s \n AMOUNT:  %d \n \n ", x.TransactionID, x.Sender, x.Receiver, x.Amount)
	}
}

func NewTransaction(sender string, receiver string, amount int) Transaction {
	ID_Count++
	block := new(Transaction)
	block.TransactionID = strconv.Itoa(ID_Count)
	block.Sender = sender
	block.Receiver = receiver
	block.Amount = amount
	return *block
}
