package core

import (
	"testing"
)

func TestInitialChainIsGenesisBlock(t *testing.T) {
	blockchain := Generate_blockchain()
	if blockchain.length != 1 && blockchain.blocks[0].proof == 0{
		t.Errorf("Initial blockchain should be genesis, instead got: %d", blockchain.length)
	}
}

func TestProofIsChecked(t *testing.T){
	blockchain := Generate_blockchain()
	prev_proof := get_latest_block(blockchain).proof
	new_proof := find_next_proof(prev_proof)
	if !is_valid_proof(prev_proof, new_proof) {
		t.Errorf("Couldn't find proof")
	}
}
