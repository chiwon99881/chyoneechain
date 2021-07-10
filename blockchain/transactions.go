package blockchain

import (
	"errors"
	"time"

	"github.com/chiwon99881/chyocoin/utils"
)

const (
	minerReward int = 50
)

// mempool은 블록체인에서 아직 확인되지 않은 즉, 블록에 아직 기록되지 않은 (곧 기록될) transactions를 의미
type mempool struct {
	Txs []*Tx
}

// Mempool Variable (언제나 같은 값을 가지는, 즉, 복사되는 메모리가 아닌 항상 같은 메모리의 값을 바라 보는 포인터)
// 얘는 왜 blockchain처럼 데이터베이스에서 복원하는 방식이 아닌 그저 변수에 할당하느냐 데이터베이스에 넣을 필요가 없음
// 언젠가 블록에 들어갈거고 그 전에 잠시 보관하기 위해 있는 메모리, 결과적으로 블록에 들어간다는 것은 데이터베이스에 이 멤풀 내 있는 트랜잭션들이
// 데이터 베이스에 들어간다는 것을 의미
var Mempool *mempool = &mempool{}

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
	TxID  string `json:"txID"`
	Index int    `json:"index"`
	Owner string `json:"owner"`
}

// TxOut Transaction에서 Output은 그 총 액수를 재분배한 결과
type TxOut struct {
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

// UTxOut is unspent transaction output
type UTxOut struct {
	TxID   string `json:"txID"`
	Index  int    `json:"index"`
	Amount int    `json:"amount"`
}

func isOnMempool(uTxOut *UTxOut) bool {
	exists := false
	// Outer: this is label for loop
Outer:
	for _, tx := range Mempool.Txs {
		for _, txIn := range tx.TxIns {
			if txIn.TxID == uTxOut.TxID && txIn.Index == uTxOut.Index {
				exists = true
				break Outer
			}
		}
	}
	return exists
}

func makeCoinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{"", -1, "COINBASE"},
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

func makeTx(from, to string, amount int) (*Tx, error) {
	if BalanceByAddress(from, Blockchain()) < amount {
		return nil, errors.New("not enough money")
	}
	var txOuts []*TxOut
	var txIns []*TxIn
	total := 0
	uTxOuts := UTxOutsByAddress(from, Blockchain())
	for _, uTxOut := range uTxOuts {
		if total >= amount {
			break
		}
		txIn := &TxIn{uTxOut.TxID, uTxOut.Index, from}
		txIns = append(txIns, txIn)
		total += uTxOut.Amount
	}
	if change := total - amount; change != 0 {
		changeTxOut := &TxOut{from, change}
		txOuts = append(txOuts, changeTxOut)
	}
	txOut := &TxOut{to, amount}
	txOuts = append(txOuts, txOut)
	tx := &Tx{
		ID:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getID()
	return tx, nil
}

func (m *mempool) AddTx(to string, amount int) error {
	tx, err := makeTx("chyoni", to, amount)
	if err != nil {
		return err
	}
	m.Txs = append(m.Txs, tx)
	return nil
}

func (m *mempool) TxToConfirm() []*Tx {
	coinbaseTx := makeCoinbaseTx("chyoni")
	txs := m.Txs
	txs = append(txs, coinbaseTx)
	m.Txs = nil
	return txs
}
