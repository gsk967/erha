package hash

// Pattern matching algorithm
// https://www.geeksforgeeks.org/rabin-karp-algorithm-for-pattern-searching/

import (
	"math"
)

const (
	D   = 255
	MOD = math.MaxInt64
)

// Checksum will generate checmsum for given chunk
// s = length of chunk
// x, y = 0,0
// for i in 1...s
//	x = (text[i-1]*pow(d,s-i)) % MOD
//	y = (x+y)% MOD

func Checksum(chunk []byte) int64 {
	x, y := int64(0), int64(0)
	s := len(chunk)
	for i := 1; i <= s; i++ {
		x = (int64(chunk[i-1]) * int64(math.Pow(D, float64(s-i)))) % MOD
		y = (x + y) % MOD
	}
	return y
}

// ChunkSlide will geneate new checksum
// It will remove previous byte and add new bytes
// hash = ((hash - (lv*pow(d,s-1))) * d + rv) % MOD
func ChunkSlide(y int64, left, right byte, size int64) int64 {
	lv, rv := int64(left), int64(right)
	y = ((y-(lv*int64(math.Pow(D, float64(size-1)))))*D + rv) % MOD
	return y
}
