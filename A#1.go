package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

/*###############################BLOCKS STRUCTURE#############################*/

type block struct {
	hash         string
	previousHash string
	nonce        int
	transaction  string
}

type blockchain struct {
	list []*block
}

func (ls *blockchain) createBlock(previousHash string, nonce int, transaction string) *block {
	blk := newBlock(previousHash, nonce, transaction)
	ls.list = append(ls.list, blk)
	return blk
}

/*###############################BLOCKCHAIN FUNCTIONS#############################*/
func newBlock(previousHash string, nonce int, transaction string) *block {
	var b block
	b.transaction = transaction
	b.nonce = nonce
	b.previousHash = previousHash
	var str string
	str = b.transaction + strconv.Itoa(b.nonce) + b.previousHash
	b.hash = calculateHash(str)
	return &b
	// A method to add new block. To keep things simple, you could provide a
	// sting of your choice as a transaction (e.g., “bob to alice”). Also, use
	// any integer value as a nonce. The CreateHash() method will provide you the
	// block Hash value.
}

func (ls *blockchain) listBlocks() {
	// A method to print all the blocks in a nice format showing block data such
	// as transaction, nonce, previous hash, current block hash
	for i := range ls.list {
		fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		fmt.Printf("Block Hash: %s\n", ls.list[i].hash)
		fmt.Printf("PreviousBlock Hash: %s\n", ls.list[i].previousHash)
		fmt.Printf("Block transaction: %s\n", ls.list[i].transaction)
		fmt.Printf("Block nonce: %d\n", ls.list[i].nonce)

	}
}

func (b *block) changeBlock(transaction string) {
	b.transaction = transaction
}

// function to change block transaction of the given block ref
// func (ls *blockchain) verifyChain() {

// 	// if len(ls.list) == 1 {
// 	var previousBlock block
// 	var currentBlock block
// 	for i := range ls.list {
// 		currentBlock = (*ls.list[i])
// 		if i == 1 {
// 			previousBlock = (*ls.list[i])
// 		}
// 		if currentBlock.hash != calculateBlockHash(currentBlock) || previousBlock.hash != currentBlock.previousHash {
// 			fmt.Println("BlockChain is not Valid")
// 		} else {
// 			fmt.Println("BlockChain is Valid")
// 		}
// 	}

// }

// function to verify blockchain in case any changes are made. This can be
// done in two different ways:

func calculateHash(stringToHash string) string { // function for calculating hash of a block
	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))

}

func calculateBlockHash(b block) string { // function for calculating hash of a block

	var str string
	str = b.transaction + strconv.Itoa(b.nonce) + b.previousHash
	b.hash = calculateHash(str)
	return str

}

func (ls *blockchain) prev_hash() string {
	// fmt.Println(ls.list[ls.list.length-1].previousHash)
	return ls.list[len(ls.list)-1].hash
}

/*###############################	MAIN	#############################*/

func main() {

	//   var blockchain []block

	//   blockchain[0]=*newBlock("Noman to AL-Cybision" , 1111 , "0")
	mniBlockChain := new(blockchain)
	mniBlockChain.createBlock("0", 101, "Noman to AL-Cybision")
	mniBlockChain.createBlock(mniBlockChain.prev_hash(), 102, "AL-Cybision to Rushad")
	mniBlockChain.createBlock(mniBlockChain.prev_hash(), 103, "Noman to Zaid")
	mniBlockChain.listBlocks()
	// mniBlockChain.verifyChain()

	//    const lastBlock = this.chain[this.chain.length - 1];

}
