package main

import (
	"fmt"
	"log"
)

func main() {
	blockchain := generate_blockchain();
	block,err := create_block(blockchain, proof_of_work(get_latest_block(blockchain).proof))

	if err != nil {
		log.Fatal(err.Error())
	}else{
		fmt.Printf("Successfully mined block %i", block.index)
	}


}
