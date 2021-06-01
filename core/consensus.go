package core

import (
	"crypto/sha256"
)

var DIFFICULTY = 2

func merge_proofs(first_proof uint64, second_proof uint64) []byte{
	var b = make([]byte, 16)
	b[0] = byte(first_proof >> 56)
	b[1] = byte(first_proof >> 48)
	b[2] = byte(first_proof >> 40)
	b[3] = byte(first_proof >> 32)
	b[4] = byte(first_proof >> 24)
	b[5] = byte(first_proof >> 16)
	b[6] = byte(first_proof >> 8)
	b[7] = byte(first_proof)
	b[8] = byte(second_proof >> 56)
	b[9] = byte(second_proof >> 48)
	b[10] = byte(second_proof >> 40)
	b[11] = byte(second_proof >> 32)
	b[12] = byte(second_proof >> 24)
	b[13] = byte(second_proof >> 16)
	b[14] = byte(second_proof >> 8)
	b[15] = byte(second_proof)
	return b
}

func is_valid_proof(prev_proof uint64, new_proof uint64) bool {
	// create hash of merged proofs
	h := sha256.New()
	h.Write(merge_proofs(prev_proof, new_proof))
	var slice = h.Sum(nil)[0:DIFFICULTY]
	// check if first DIFFICULTY number of bytes are 0
	for _, elem := range slice {
		if elem != 0 {
			return false
		}
	}
	return true
}

func find_next_proof(prev_proof uint64) uint64 {
	var new_proof uint64 = 0
	for ; !is_valid_proof(prev_proof, new_proof); new_proof++ {}
	return new_proof
}
