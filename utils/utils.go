package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"strings"
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

// FromBytes is function of byte to string
func FromBytes(i interface{}, data []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	HandleError(decoder.Decode(i))
}

// Hash is function of any value to hash
func Hash(i interface{}) string {
	s := fmt.Sprintf("%v", i)
	hash := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hash)
}

// Splitter is function of split string
func Splitter(s, sep string, i int) string {
	r := strings.Split(s, sep)
	if len(r)-1 < i {
		return ""
	}
	return r[i]
}
