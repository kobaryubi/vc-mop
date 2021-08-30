package blockchain

import (
	"time"
)

// Block ブロックチェーンのブロックを定義する型
type Block struct {
	prevHash     [32]byte
	transactions []string
	timestamp    int64
}

// NewBlock ブロックを作成する関数
func NewBlock(prevHash [32]byte, transactions []string) *Block {
	b := &Block{}

	b.prevHash = prevHash
	b.transactions = transactions
	b.timestamp = time.Now().UnixNano()

	return b
}
