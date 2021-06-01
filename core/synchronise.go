package core

import (
	"bytes"
	"errors"
)

// Synchronization of the blockchain is what makes
// the decentralised network so powerful
// Every full node in the network has it's local copy of the entire blockchain

// If there were a harmful node in the network, then other nodes will have to vote it out,
// almost like a democracy

// The luar network operates on trust, that each full node on the
// network abides by the policies put in place

// This however does cause points of vulnerability
// 	1. Blockchain becomes corrupted
// 	2. Nodes could poison blocks (send viruses in the block)
// 	3. Nodes can manipulate their local blockchain and enforce it on others
// 	4. Nodes can re-spend transactions
// 	5. Updates and Forks can be forced onto existing nodes causing bugs/failures
// 	6. If an attacker takes over 50% of the network, they can prevent miners from spending

// 	We tackle each of these points on a local and remote level
// 	1. Verify the local blockchains integrity using block hashes then attempt to synchronise
// 		a. If local blockchain has become corrupted
// 			i. Confirm this with all other neighbour nodes
// 			ii. Delete the blockchain from the point of corruption and re-synchronize
// 	2. Only accept incoming blocks that are of the correct format
// 	3. Choose to accept or reject blocks based on
// 		a. The new blocks follow from the old blocks
// 		b. How old the blocks are
// 		c. How trusted the node is (we issue levels of trust based on digital signatures)
// 	4. Every full node must validate transactions to ensure money cannot be re-spent
// 	5. Updates will never be forcefully pushed,
//	   nodes will have to choose to update to the newest version of the chain
// 	   Nodes which operate on old chains shouldn't accept any blocks of later versions
// 	   But should accept those of earlier versions (hence back support)
// 	6. If an attacker gains close to 50% of the hashrate of the entire network then other nodes can vote
//     for regulations (aka. restrictions) to be placed on the nodes. This vote only requires 40% of the
// 	   nodes to agree upon.
//     Restrictions can include
//  		- Limited block creation / mining
//  		- Capped staking ability
//  		- Ability to extend chain by 1 million blocks (aka. print 1,000,000 new coins) (this can only ever be done once)

func is_local_chain_valid(blockchain *Blockchain) bool {
	var idx uint64
	for idx = 0; idx < blockchain.length-1; idx++ {
		if !is_valid_block_link(&blockchain.blocks[idx], &blockchain.blocks[idx+1]) {
			return false
		}
	}
	return true
}

func verify_chains_match(local *Blockchain, remote *Blockchain) (int, error) {
	if local.length > remote.length {
		return -1, errors.New("local blockchain is ahead of remote")
	} else if local.length < remote.length {
		return -1, errors.New("local blockchain is behind remote")
	} else {
		var idx uint64
		for idx = 0; idx < local.length; idx++ {
			if !bytes.Equal(local.blocks[idx].previous_hash, remote.blocks[idx].previous_hash) {
				return int(idx), errors.New("local blockchain is possibly corrupted at this block")
			}
		}
		return -1, nil
	}
}
