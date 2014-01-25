package graph

import (
	crand "crypto/rand"
	mrand "math/rand"
	"strconv"
)

const idLenBytes = 16 // 128 bits

// Attempts to use cyrpto/rand;
// falls back to math/rand
func GenerateID() TaskID {
	var b [idLenBytes]byte
	if !fillCrypto(b[:]) {
		fillMath(b[:])
	}
	id := ""
	for _, b := range b {
		id += strconv.FormatInt(int64(b), 10)
	}
	return TaskID(id)
}

// Returns true on success, otherwise false
func fillCrypto(b []byte) bool {
	for len(b) > 0 {
		n, err := crand.Read(b)
		if err != nil {
			return false
		}
		b = b[n:]
	}
	return true
}

func fillMath(b []byte) {
	for i := range b {
		b[i] = byte(mrand.Uint32())
	}
}
