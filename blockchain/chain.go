package blockchain

import (
	"sync"

	"github.com/chiwon99881/chyocoin/db"
	"github.com/chiwon99881/chyocoin/utils"
)

const (
	defaultDifficulty  int = 2 //최초의 Difficulty
	difficultyInterval int = 5 // "5"개의 블록이 블록체인에 생성될 때마다 difficulty를 다시 계산
	blockInterval      int = 2 // 매 "2"분마다 블록 1개가 블록체인에 생성
	allowedRange       int = 2 // 딱 10분을 기준으로 Difficulty를 줄이고 높이고는 너무 엄격하니 플러스 마이너스 2분 간격
)

type blockchain struct {
	NewestHash        string `json:"newestHash"`
	Height            int    `json:"height"`
	CurrentDifficulty int    `json:"currentDifficulty"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock() {
	block := createBlock(b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	b.persist()
}

func (b *blockchain) Blocks() []*Block {
	var blocks []*Block
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

func (b *blockchain) recalculateDifficulty() int {
	allBlocks := b.Blocks()
	newestBlock := allBlocks[0]
	lastRecalculateBlock := allBlocks[difficultyInterval-1]
	actualTime := (newestBlock.Timestamp / 60) - (lastRecalculateBlock.Timestamp / 60)
	expectedTime := difficultyInterval * blockInterval
	if actualTime >= (expectedTime + allowedRange) {
		return b.CurrentDifficulty - 1
	} else if actualTime <= (expectedTime - allowedRange) {
		return b.CurrentDifficulty + 1
	}
	return b.CurrentDifficulty
}

func (b *blockchain) difficulty() int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		return b.recalculateDifficulty()
	} else {
		return b.CurrentDifficulty
	}
}

// Blockchain for initalize instance or return used blockchain
func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{
				Height: 0,
			}
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock()
			} else {
				b.restore(checkpoint)
			}
		})
	}
	return b
}
