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
