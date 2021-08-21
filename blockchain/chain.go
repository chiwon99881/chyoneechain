package blockchain

import (
	"encoding/json"
	"net/http"
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
	m                 sync.Mutex
}

type storage interface {
	FindBlock(hash string) []byte
	SaveBlock(hash string, data []byte)
	SaveChain(data []byte)
	LoadChain() []byte
	DeleteAllBlock()
}

var b *blockchain
var once sync.Once

var dbStorage storage = db.DB{}

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) AddBlock() *Block {
	b.m.Lock()
	defer b.m.Unlock()
	block := createBlock(b.NewestHash, b.Height+1, getDifficulty(b))
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	persistBlockchain(b)
	return block
}

func persistBlockchain(b *blockchain) {
	dbStorage.SaveChain(utils.ToBytes(b))
}

// Blocks is get all block in blockchain
func Blocks(b *blockchain) []*Block {
	b.m.Lock()
	defer b.m.Unlock()
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

// Txs is find all transactions in blockchain
func Txs(b *blockchain) []*Tx {
	var txs []*Tx
	for _, block := range Blocks(b) {
		txs = append(txs, block.Transactions...)
	}
	return txs
}

// FindTx is find one transaction
func FindTx(b *blockchain, targetID string) *Tx {
	for _, tx := range Txs(b) {
		if tx.ID == targetID {
			return tx
		}
	}
	return nil
}

func recalculateDifficulty(b *blockchain) int {
	allBlocks := Blocks(b)
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

func getDifficulty(b *blockchain) int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		return recalculateDifficulty(b)
	} else {
		return b.CurrentDifficulty
	}
}

// UTxOutsByAddress is unspent transaction output by address
func UTxOutsByAddress(address string, b *blockchain) []*UTxOut {
	var uTxOuts []*UTxOut
	creatorTxs := make(map[string]bool)
	for _, block := range Blocks(b) {
		for _, tx := range block.Transactions {
			for _, input := range tx.TxIns {
				if input.Signature == "COINBASE" {
					break
				}
				if FindTx(b, input.TxID).TxOuts[input.Index].Address == address {
					creatorTxs[input.TxID] = true
				}
			}
			for index, output := range tx.TxOuts {
				if output.Address == address {
					if _, ok := creatorTxs[tx.ID]; !ok {
						uTxOut := &UTxOut{tx.ID, index, output.Amount}
						if !isOnMempool(uTxOut) {
							uTxOuts = append(uTxOuts, uTxOut)
						}
					}
				}
			}
		}
	}
	return uTxOuts
}

// BalanceByAddress is total balance by address
func BalanceByAddress(address string, b *blockchain) int {
	var amount int
	txOuts := UTxOutsByAddress(address, b)
	for _, txOut := range txOuts {
		amount += txOut.Amount
	}
	return amount
}

// Blockchain for initalize instance or return used blockchain
func Blockchain() *blockchain {
	once.Do(func() {
		b = &blockchain{
			Height: 0,
		}
		checkpoint := dbStorage.LoadChain()
		if checkpoint == nil {
			b.AddBlock()
		} else {
			b.restore(checkpoint)
		}
	})
	return b
}

// Status is function of see blockchain status
func Status(b *blockchain, rw http.ResponseWriter) {
	b.m.Lock()
	defer b.m.Unlock()
	utils.HandleError(json.NewEncoder(rw).Encode(b))
}

func (b *blockchain) Replace(newBlocks []*Block) {
	b.m.Lock()
	defer b.m.Unlock()
	b.CurrentDifficulty = newBlocks[0].Difficulty
	b.Height = len(newBlocks)
	b.NewestHash = newBlocks[0].Hash
	persistBlockchain(b)
	dbStorage.DeleteAllBlock()
	for _, block := range newBlocks {
		persistBlock(block)
	}
}

func (b *blockchain) AddPeerBlock(block *Block) {
	b.m.Lock()
	m.m.Lock()
	defer b.m.Unlock()
	defer m.m.Unlock()

	b.Height++
	b.CurrentDifficulty = block.Difficulty
	b.NewestHash = block.Hash

	persistBlockchain(b)
	persistBlock(block)

	for _, tx := range block.Transactions {
		_, ok := m.Txs[tx.ID]
		if ok {
			delete(m.Txs, tx.ID)
		}
	}
}
