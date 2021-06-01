package wallet

import "blockchain/core"

type Wallet struct {
	hash string
	amount float64
	outbound []core.Transaction
	inbound []core.Transaction
	private_key []byte
	public_key []byte
	wallet_address []byte
}

func create(blockchain *core.Blockchain) {
}