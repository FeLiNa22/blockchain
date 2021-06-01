package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

type Block struct {
	version       uint8
	index         uint64
	timestamp     uint64
	transactions  []Transaction
	proof         uint64
	previous_hash []byte
}

/*
   Creates a SHA-256 hash of a Block
   :param block: <dict> Block
   :return: <str>
*/
func hash(block *Block) []byte {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", *block)))
	return h.Sum(nil)
}

/*
	Validates if two blocks are linked correctly
	:param block: <dict> Block
	:param block: <dict> Block
	:return <bool> true if blocks are linked
*/
func is_valid_block_link(prev_block *Block, new_block *Block) bool {
	return bytes.Equal(hash(prev_block), new_block.previous_hash)
}

/*
	Dump the structure to a block file
*/
func Dump(block *Block) {
	f, err := os.Create(fmt.Sprintf("blocks/%q.bin", block.index))
	if err != nil {
		log.Fatal("Couldn't create file")
	}

	err = binary.Write(f, binary.LittleEndian, block)
	if err != nil {
		log.Fatal("Write failed")
	}

	err = f.Close()
	if err != nil {
		log.Fatal("Couldn't save block to file")
	}
}

/*
	reconstruct the block from a binary file
*/
func Load(index uint64) *Block {
	f, err := os.Open(fmt.Sprintf("blocks/%q.bin", index))
	if err != nil {
		log.Fatal("Couldn't open file")
	}

	block := Block{}
	err = binary.Read(f, binary.LittleEndian, block)
	if err != nil {
		log.Fatal("Couldn't read file")
	}

	err = f.Close()
	if err != nil {
		log.Fatal("Couldn't save block to file")
	}

	return &block
}
