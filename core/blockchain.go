package core

import (
	"bytes"
	_ "crypto/sha256"
	"errors"
	_ "fmt"
	"log"
	_ "math"
	"time"
)

var VERSION uint8 = 1

type Blockchain struct {
	version              uint8
	length               uint64
	blocks               []Block
	current_transactions []Transaction
}

/*
	Create a fork of the original blockchain
*/
func Generate_blockchain() *Blockchain {
	blockchain := Blockchain{VERSION, 0, []Block{}, []Transaction{}}
	block, err := create_genesis_block(&blockchain)
	if err != nil {
		log.Fatal("Could not initiate the blockchain")
	} else {
		blockchain.blocks = append(blockchain.blocks, *block)
		// increment length of the chain
		blockchain.length++
	}
	return &blockchain
}

/*
	Creates the initial block in the chain
*/
func create_genesis_block(blockchain *Blockchain) (*Block, error) {
	// check blockchain is not already initiated
	if blockchain.length != 0 {
		return nil, errors.New("blockchain has already been initiated")
	}
	// create the initial block
	block := create_block(blockchain, 1, []byte{1})
	return block, nil
}

/*
	Adds a block to the local Blockchain
	:param block: <Block> The block to be added to the blockchain
	:return: <error> Any errors generated when adding the block
*/
func add_block(blockchain *Blockchain, new_block *Block) error {
	prev_block := get_latest_block(blockchain)

	// check new_block version is compatible with local blockchain version
	if blockchain.version < new_block.version {
		return errors.New("new_block is incompatible with this nodes blockchain")
	}

	// check hash of prev_block and new_block match
	if bytes.Equal(hash(prev_block), new_block.previous_hash) {
		return errors.New("new_block hash does not match with the block header")
	}

	// check the proof in the new new_block is valid
	if !is_valid_proof(prev_block.proof, new_block.proof) {
		return errors.New("new_block proof is invalid")
	}

	// check the new_block time stamp is not more than 1 hour into the future
	if !(prev_block.timestamp <= new_block.timestamp+(1000*60*60)) {
		return errors.New("new_block timestamp is made before the head new_block")
	}

	// append new_block to the chain
	blockchain.blocks = append(blockchain.blocks, *new_block)

	// increment length of the chain
	blockchain.length++

	// empty current transactions
	blockchain.current_transactions = nil

	return nil
}

/*
	Create a new Block
	:param proof: <int> The proof given by the Proof of Work algorithm
	:return: <dict> New Block
*/
func create_block(blockchain *Blockchain, proof uint64, prev_hash []byte) *Block {

	// create the new block
	block := Block{
		blockchain.version,
		blockchain.length,
		uint64(time.Now().UnixNano()),
		blockchain.current_transactions,
		proof,
		prev_hash,
	}

	return &block
}

/*
	Creates a new transaction to go into the next mined Block
	:param sender: <str> Address of the Sender
	:param recipient: <str> Address of the Recipient
	:param amount: <int> Amount
	:return: <int> The index of the Block that will hold this transaction
*/
func create_transaction(blockchain *Blockchain, sender string, receiver string, amount float64) uint64 {
	transaction := Transaction{sender, receiver, amount}
	blockchain.current_transactions = append(blockchain.current_transactions, transaction)
	return get_latest_block(blockchain).index + 1
}

/*
	Gets the last block in the blockchain
	:return: <Block> The Block at the end of the chain
*/
func get_latest_block(blockchain *Blockchain) *Block {
	return &blockchain.blocks[blockchain.length-1]
}
