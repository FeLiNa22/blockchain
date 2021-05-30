package blockchain

import (
	"crypto/sha256"
	_ "crypto/sha256"
	"fmt"
	_ "fmt"
	_ "math"
	"time"
)

type Transaction struct {
	sender   string
	receiver string
	amount   float64
}

type Block struct {
	index         int64
	timestamp     int64
	transactions  []Transaction
	proof         int64
	previous_hash string
}

type Blockchain struct {
	length               int64
	blocks               []Block
	current_transactions []Transaction
}

/*
	Create a new Block in the Blockchain
	:param proof: <int> The proof given by the Proof of Work algorithm
	:param previous_hash: (Optional) <str> Hash of previous Block
	:return: <dict> New Block
*/
func new_block(blockchain *Blockchain, proof int64, previous_hash string) Block {

	// create the new block
	block := Block{
		blockchain.length, time.Now().UnixNano(), blockchain.current_transactions, proof,
		previous_hash,
	}

	// empty current transactions
	blockchain.current_transactions = nil

	// append block the chain
	blockchain.blocks = append(blockchain.blocks, block)

	return block
}

/*
	Creates a new transaction to go into the next mined Block
	:param sender: <str> Address of the Sender
	:param recipient: <str> Address of the Recipient
	:param amount: <int> Amount
	:return: <int> The index of the Block that will hold this transaction
*/
func new_transaction(blockchain *Blockchain, sender string, receiver string, amount float64) int64 {
	transaction := Transaction{sender, receiver, amount}
	blockchain.current_transactions = append(blockchain.current_transactions, transaction)
	return get_last(blockchain).index + 1
}

/*
   Creates a SHA-256 hash of a Block
   :param block: <dict> Block
   :return: <str>
*/
func hash(block *Block) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", block)))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func get_last(blockchain *Blockchain) Block {
	return blockchain.blocks[len(blockchain.blocks)-1]
}

func consensus(blockchain *Blockchain) {

}
