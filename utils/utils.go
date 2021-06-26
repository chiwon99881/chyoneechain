package utils

import (
	"bytes"
	"encoding/gob"
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
