package db

import (
	"fmt"
	"os"

	"github.com/chiwon99881/chyocoin/utils"
	bolt "go.etcd.io/bbolt"
)

const (
	dbName       = "blockchain"
	dataBucket   = "data"
	blocksBucket = "blocks"
	checkpoint   = "checkpoint"
)

type DB struct{}

func (DB) FindBlock(hash string) []byte {
	return findBlock(hash)
}

func (DB) SaveBlock(hash string, data []byte) {
	saveBlock(hash, data)
}

func (DB) DeleteAllBlock() {
	emptyBlocks()
}

func (DB) SaveChain(data []byte) {
	saveChain(data)
}

func (DB) LoadChain() []byte {
	return loadChain()
}

var db *bolt.DB

func getDbName() string {
	port := os.Args[2][6:]
	return fmt.Sprintf("%s_%s.db", dbName, port)
}

// DB initialize
func InitDB() {
	if db == nil {
		// Open creates and opens a database at the given path.
		// If the file does not exist then it will be created automatically.
		dbPointer, err := bolt.Open(getDbName(), 0600, nil)
		db = dbPointer
		utils.HandleError(err)
		err = db.Update(func(t *bolt.Tx) error {
			_, err := t.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleError(err)
			_, err = t.CreateBucketIfNotExists([]byte(blocksBucket))
			return err
		})
		utils.HandleError(err)
	}
}

// Close is function of close database
func Close() {
	db.Close()
}

// SaveBlock is saved of block in blockchain DB
func saveBlock(hash string, data []byte) {
	err := db.Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte(hash), data)
		return err
	})
	utils.HandleError(err)
}

// SaveBlockchain is saved of chain in blockchain DB
func saveChain(data []byte) {
	err := db.Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		err := bucket.Put([]byte(checkpoint), data)
		return err
	})
	utils.HandleError(err)
}

// Checkpoint is function of current blockchain pointer
func loadChain() []byte {
	var data []byte
	db.View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		data = bucket.Get([]byte(checkpoint))
		return nil
	})
	return data
}

// Block is function of get one block
func findBlock(hash string) []byte {
	var data []byte
	db.View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blocksBucket))
		data = bucket.Get([]byte(hash))
		return nil
	})
	return data
}

// EmptyBlocks is function of delete bucket and data
func emptyBlocks() {
	db.Update(func(t *bolt.Tx) error {
		utils.HandleError(t.DeleteBucket([]byte(blocksBucket)))
		_, err := t.CreateBucket([]byte(blocksBucket))
		utils.HandleError(err)
		return nil
	})
}
