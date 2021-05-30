package blockchain

type Wallet struct {
	hash string
	amount float64
	outbound []Transaction
	inbound []Transaction
}

func create(blockchain *Blockchain) {
}