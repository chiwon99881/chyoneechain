package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"strconv"

	"github.com/chiwon99881/chyocoin/db"
	"github.com/chiwon99881/chyocoin/utils"
)

// Block struct
type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

func (b *Block) toBytes() []byte {
	var blockBuffer bytes.Buffer
	// gob은 byte를 encode / decode할 수 있는 package
	encoder := gob.NewEncoder(&blockBuffer)
	utils.HandleError(encoder.Encode(b))
	return blockBuffer.Bytes()
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, b.toBytes())
}

func createBlock(data, prevHash string, height int) *Block {
	block := Block{
		Data:     data,
		Hash:     "",
		PrevHash: prevHash,
		Height:   height,
	}
	payload := block.Data + block.PrevHash + strconv.Itoa(block.Height)
	block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	block.persist()
	return &block
}
