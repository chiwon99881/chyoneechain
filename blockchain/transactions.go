package blockchain

import (
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/chiwon99881/chyocoin/utils"
	"github.com/chiwon99881/chyocoin/wallet"
)

const (
	minerReward int = 50
)

// mempool은 블록체인에서 아직 확인되지 않은 즉, 블록에 아직 기록되지 않은 (곧 기록될) transactions를 의미
type mempool struct {
	Txs map[string]*Tx
	m   sync.Mutex
}

// Mempool Variable (언제나 같은 값을 가지는, 즉, 복사되는 메모리가 아닌 항상 같은 메모리의 값을 바라 보는 포인터)
// 얘는 왜 blockchain처럼 데이터베이스에서 복원하는 방식이 아닌 그저 변수에 할당하느냐 데이터베이스에 넣을 필요가 없음
// 언젠가 블록에 들어갈거고 그 전에 잠시 보관하기 위해 있는 메모리, 결과적으로 블록에 들어간다는 것은 데이터베이스에 이 멤풀 내 있는 트랜잭션들이
// 데이터 베이스에 들어간다는 것을 의미
var m *mempool
var memOnce sync.Once

// Mempool is function of return mempool
func Mempool() *mempool {
	memOnce.Do(func() {
		m = &mempool{
			Txs: make(map[string]*Tx),
		}
	})
	return m
}

// Tx struct
type Tx struct {
	ID        string   `json:"id"`
	Timestamp int      `json:"timestamp"`
	TxIns     []*TxIn  `json:"txIns"`
	TxOuts    []*TxOut `json:"txOuts"`
}

// TxIn Transaction에서 Input은 특정 사람이 가지는 지갑의 총 액수
type TxIn struct {
	TxID      string `json:"txID"`
	Index     int    `json:"index"`
	Signature string `json:"signature"`
}

// TxOut Transaction에서 Output은 그 총 액수를 재분배한 결과
type TxOut struct {
	Address string `json:"address"`
	Amount  int    `json:"amount"`
}

// UTxOut is unspent transaction output
type UTxOut struct {
	TxID   string `json:"txID"`
	Index  int    `json:"index"`
	Amount int    `json:"amount"`
}

func (t *Tx) getID() {
	t.ID = utils.Hash(t)
}

func (t *Tx) sign() {
	for _, txIn := range t.TxIns {
		txIn.Signature = wallet.Sign(t.ID, wallet.Wallet())
	}
}

func validate(t *Tx) bool {
	valid := true
	for _, txIn := range t.TxIns {
		prevTx := FindTx(Blockchain(), txIn.TxID)
		if prevTx == nil {
			valid = false
			break
		}
		address := prevTx.TxOuts[txIn.Index].Address
		valid = wallet.Verify(txIn.Signature, t.ID, address)
		if !valid {
			break
		}
	}
	return valid
}

func isOnMempool(uTxOut *UTxOut) bool {
	exists := false
	// Outer: this is label for loop
Outer:
	for _, tx := range Mempool().Txs {
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
	tx.sign()
	valid := validate(tx)
	if !valid {
		return nil, errors.New("this transaction invalid")
	}
	return tx, nil
}

func (m *mempool) AddTx(to string, amount int) (*Tx, error) {
	m.m.Lock()
	defer m.m.Unlock()
	tx, err := makeTx(wallet.Wallet().Address, to, amount)
	if err != nil {
		return nil, err
	}
	m.Txs[tx.ID] = tx
	return tx, nil
}

func (m *mempool) TxToConfirm() []*Tx {
	coinbaseTx := makeCoinbaseTx(wallet.Wallet().Address)
	var txs []*Tx
	for _, tx := range m.Txs {
		txs = append(txs, tx)
	}
	txs = append(txs, coinbaseTx)
	m.Txs = make(map[string]*Tx)
	return txs
}

// MempoolStatus is function of see mempool status.
func MempoolStatus(m *mempool, rw http.ResponseWriter) {
	m.m.Lock()
	defer m.m.Unlock()
	utils.HandleError(json.NewEncoder(rw).Encode(m.Txs))
}

// AddPeerTx is function of append peer transaction to mempool
func (m *mempool) AddPeerTx(tx *Tx) {
	m.m.Lock()
	defer m.m.Unlock()

	m.Txs[tx.ID] = tx
}
