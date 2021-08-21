package blockchain

import (
	"sync"
	"testing"

	"github.com/chiwon99881/chyocoin/utils"
)

type fakeDB struct {
	fakeLoadChain func() []byte
	fakeFindBlock func() []byte
}

func (f fakeDB) FindBlock(hash string) []byte {
	return f.fakeFindBlock()
}

func (f fakeDB) LoadChain() []byte {
	return f.fakeLoadChain()
}

func (fakeDB) SaveBlock(hash string, data []byte) {}
func (fakeDB) SaveChain(data []byte)              {}
func (fakeDB) DeleteAllBlock()                    {}

func TestBlockchain(t *testing.T) {
	t.Run("Should create blockchain", func(t *testing.T) {
		dbStorage = fakeDB{
			fakeLoadChain: func() []byte {
				return nil
			},
		}
		blockchain := Blockchain()
		if blockchain.Height != 1 || blockchain.CurrentDifficulty != 2 {
			t.Error("a newborn blockchain height should be 1 and difficulty should be 2")
		}
	})

	t.Run("Should restore blockchain", func(t *testing.T) {
		once = *new(sync.Once)
		dbStorage = fakeDB{
			fakeLoadChain: func() []byte {
				bc := &blockchain{
					Height:            2,
					NewestHash:        "xxx",
					CurrentDifficulty: 2,
				}
				return utils.ToBytes(bc)
			},
		}
		blockchain := Blockchain()
		if blockchain.Height != 2 {
			t.Errorf("a restore blockchain height should be 2, but got %d", blockchain.Height)
		}
	})
}

func TestBlocks(t *testing.T) {
	t.Run("PrevHash exists", func(t *testing.T) {
		fakeBlocks := 0
		b := &Block{}
		dbStorage = fakeDB{
			fakeFindBlock: func() []byte {
				if fakeBlocks == 0 {
					b = &Block{
						Height:   1,
						PrevHash: "xx",
					}
				} else {
					b = &Block{
						Height: 1,
					}
				}
				fakeBlocks++
				return utils.ToBytes(b)
			},
		}
		bc := &blockchain{}
		blocks := Blocks(bc)
		if len(blocks) <= 1 {
			t.Error("if prevHash is not nil, blocks length should be more 1")
		}
	})

	t.Run("PrevHash not exists", func(t *testing.T) {
		dbStorage = fakeDB{
			fakeFindBlock: func() []byte {
				b := &Block{
					Height: 1,
				}
				return utils.ToBytes(b)
			},
		}
		bc := &blockchain{}
		blocks := Blocks(bc)
		if len(blocks) != 1 {
			t.Error("if prevHash is nil, blocks length should be 1")
		}
	})
}

func TestFindTx(t *testing.T) {
	t.Run("Tx not found", func(t *testing.T) {
		dbStorage = fakeDB{
			fakeFindBlock: func() []byte {
				b := &Block{
					Height:       1,
					Transactions: []*Tx{},
				}
				return utils.ToBytes(b)
			},
		}
		tx := FindTx(&blockchain{NewestHash: "x"}, "test")
		if tx != nil {
			t.Error("Tx should be not found.")
		}
	})

	t.Run("Tx should be found", func(t *testing.T) {
		dbStorage = fakeDB{
			fakeFindBlock: func() []byte {
				b := &Block{
					Height: 1,
					Transactions: []*Tx{
						{ID: "test"},
					},
				}
				return utils.ToBytes(b)
			},
		}
		tx := FindTx(&blockchain{NewestHash: "x"}, "test")
		if tx == nil {
			t.Error("Tx should be found.")
		}
	})
}

func TestGetDifficulty(t *testing.T) {
	blocks := []*Block{
		{PrevHash: "x"},
		{PrevHash: "x"},
		{PrevHash: "x"},
		{PrevHash: "x"},
		{PrevHash: ""},
	}
	fakeBlocks := 0
	dbStorage = fakeDB{
		fakeFindBlock: func() []byte {
			defer func() {
				fakeBlocks++
			}()
			return utils.ToBytes(blocks[fakeBlocks])
		},
	}
	type test struct {
		height int
		want   int
	}
	tests := []test{
		{height: 0, want: defaultDifficulty},
		{height: 2, want: defaultDifficulty},
		{height: 5, want: 3},
	}
	for _, tc := range tests {
		bc := &blockchain{Height: tc.height, CurrentDifficulty: defaultDifficulty}
		got := getDifficulty(bc)
		if got != tc.want {
			t.Errorf("difficulty should be %d, but got %d", tc.want, got)
		}
	}
}

func TestAddPeerBlock(t *testing.T) {
	bc := &blockchain{
		Height:            1,
		CurrentDifficulty: 1,
		NewestHash:        "xx",
	}
	nb := &Block{
		Difficulty: 2,
		Hash:       "test",
		Transactions: []*Tx{
			{ID: "test"},
		},
	}
	Mempool().Txs["test"] = &Tx{}
	bc.AddPeerBlock(nb)
	if bc.Height != 2 || bc.CurrentDifficulty != nb.Difficulty {
		t.Error("peer chain's height should be current height + 1 and current difficulty should be new block's difficulty.")
	}
}

func TestReplace(t *testing.T) {
	bc := &blockchain{
		Height:            1,
		CurrentDifficulty: 1,
		NewestHash:        "xx",
	}
	blocks := []*Block{
		{Difficulty: 3, Hash: "test"},
		{Difficulty: 3, Hash: "test"},
	}
	bc.Replace(blocks)
	if bc.CurrentDifficulty != blocks[0].Difficulty || bc.Height != len(blocks) || bc.NewestHash != blocks[0].Hash {
		t.Error("blockchain should be replaced.")
	}
}
