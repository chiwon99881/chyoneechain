package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
)

// Block struct
type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

type blockchain struct {
	blocks []*Block
}

var b *blockchain
var once sync.Once

func (b *Block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	totalBlock := len(GetBlockchain().blocks)
	if totalBlock == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlock-1].Hash
}

func createBlock(data string) *Block {
	newBlock := Block{data, "", getLastHash(), len(GetBlockchain().blocks) + 1}
	newBlock.calculateHash()
	return &newBlock
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

// GetBlockchain for initalize instance or return used blockchain
func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis")
		})
	}
	return b
}

func (b *blockchain) AllBlocks() []*Block {
	return b.blocks
}

func (b *blockchain) Getblock(height int) (*Block, error) {
	currentMaxHeight := len(b.blocks)
	if height > currentMaxHeight {
		return nil, errors.New("height is too larger than current max height")
	}
	return b.blocks[height-1], nil
}
