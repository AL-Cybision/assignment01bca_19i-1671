package assignment01bca_19i_1671

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

// function to change block transaction of the given block ref
func (b *block) changeBlock(transaction string) {
	b.transaction = transaction
}

func (ls *blockchain) verifyChain() bool {
	valid := true
	// fmt.Println(valid)

	for i := range ls.list {
		// var valid = true

		if i < (len(ls.list) - 1) {
			fmt.Println("ls.list[i].hash", ls.list[i].hash)
			fmt.Println("ls.list[i+1].previousHash", ls.list[i+1].previousHash)
			fmt.Println("ls.list[i].hash", ls.list[i].hash)
			fmt.Println(" calculateBlockHash(*ls.list[i])", calculateBlockHash(*ls.list[i]))
			if ls.list[i].hash != ls.list[i+1].previousHash || ls.list[i].hash != calculateBlockHash(*ls.list[i]) {

				// fmt.Println("Block is Invalid")
				valid = false
			} else {
				// fmt.Println("Block is Valid")
				// valid = true
			}
		} else {
			if ls.list[i].hash != calculateBlockHash(*ls.list[i]) {

				// fmt.Println("Block is Invalid")
				valid = false
			} else {
				// fmt.Println("Block is Valid")
				// valid = true
			}

		}

	}

	return valid
}

func calculateHash(stringToHash string) string { // function for calculating hash of a block
	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))

}

func calculateBlockHash(b block) string { // function for calculating hash of a block

	var str string
	str = b.transaction + strconv.Itoa(b.nonce) + b.previousHash
	b.hash = calculateHash(str)
	return b.hash

}

func (ls *blockchain) prev_hash() string {
	// fmt.Println(ls.list[ls.list.length-1].previousHash)
	return ls.list[len(ls.list)-1].hash
}

/*###############################	MAIN	#############################*/

// func main() {

// 	//   var blockchain []block

// 	//   blockchain[0]=*newBlock("Noman to AL-Cybision" , 1111 , "0")
// 	mniBlockChain := new(blockchain)
// 	mniBlockChain.createBlock("0", 101, "Noman to AL-Cybision")
// 	mniBlockChain.createBlock(mniBlockChain.prev_hash(), 102, "AL-Cybision to Rushad")
// 	mniBlockChain.createBlock(mniBlockChain.prev_hash(), 103, "Noman to Zaid")
// 	// mniBlockChain.listBlocks()
// 	// fmt.Println("Output After Block Changed")
// 	// mniBlockChain.list[1].changeBlock("AL-Cybision to Al-Rushad")
// 	mniBlockChain.listBlocks()
// 	if mniBlockChain.verifyChain() {
// 		fmt.Println("BlockChain is Valid")
// 	} else {
// 		fmt.Println("BlockChain is Not Valid")
// 	}

// 	//    const lastBlock = this.chain[this.chain.length - 1];

// }
