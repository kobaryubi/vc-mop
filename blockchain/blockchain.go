package blockchain

import (
	"time"
)

type Block struct {
	prevHash     [32]byte
	transactions []string
	timestamp    int64
}

func NewBlock(prevHash [32]byte, transactions []string) *Block {
	b := &Block{}

	b.prevHash = prevHash
	b.transactions = transactions
	b.timestamp = time.Now().UnixNano()

	return b
}
