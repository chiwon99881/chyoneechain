package blockchain

import (
	"time"

	"github.com/chiwon99881/chyocoin/utils"
)

const (
	minerReward int = 50
)

// Tx struct
type Tx struct {
	ID        string   `json:"id"`
	Timestamp int      `json:"timestamp"`
	TxIns     []*TxIn  `json:"txIns"`
	TxOuts    []*TxOut `json:"txOuts"`
}

func (t *Tx) getID() {
	t.ID = utils.Hash(t)
}

// TxIn Transaction에서 Input은 특정 사람이 가지는 지갑의 총 액수
type TxIn struct {
	Owner  string
	Amount int
}

// TxOut Transaction에서 Output은 그 총 액수를 재분배한 결과
type TxOut struct {
	Owner  string
	Amount int
}

func makeCoinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{"COINBASE", minerReward},
	}
	txOuts := []*TxOut{
		{address, minerReward},
	}
	tx := Tx{
		ID:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getID()
	return &tx
}
