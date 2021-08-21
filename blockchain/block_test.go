package blockchain

import (
	"reflect"
	"testing"

	"github.com/chiwon99881/chyocoin/utils"
)

func TestCreateBlock(t *testing.T) {
	dbStorage = fakeDB{}
	Mempool().Txs["test"] = &Tx{}
	b := createBlock("x", 1, 1)
	if reflect.TypeOf(b) != reflect.TypeOf(&Block{}) {
		t.Error("createBlock() should return an instance of a block struct")
	}
}

func TestFindBlock(t *testing.T) {
	t.Run("Block not found", func(t *testing.T) {
		dbStorage = fakeDB{
			fakeFindBlock: func() []byte {
				return nil
			},
		}
		_, err := FindBlock("xx")
		if err == nil {
			t.Error("the block should not be found.")
		}
	})
	t.Run("Block not found", func(t *testing.T) {
		dbStorage = fakeDB{
			fakeFindBlock: func() []byte {
				b := &Block{
					Height: 1,
				}
				return utils.ToBytes(b)
			},
		}
		block, _ := FindBlock("xx")
		if block.Height != 1 {
			t.Error("the block height should be 1.")
		}
	})
}
