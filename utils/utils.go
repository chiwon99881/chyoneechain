package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

// HandleError function for handle error
func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// ToBytes is anything convert to bytes
// i interface{} means receive all types you want passed.
func ToBytes(i interface{}) []byte {
	var aBuffer bytes.Buffer
	// gob은 byte를 encode / decode할 수 있는 package
	encoder := gob.NewEncoder(&aBuffer)
	HandleError(encoder.Encode(i))
	return aBuffer.Bytes()
}

func FromBytes(i interface{}, data []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	HandleError(decoder.Decode(i))
}

func Hash(i interface{}) string {
	s := fmt.Sprintf("%v", i)
	hash := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hash)
}
