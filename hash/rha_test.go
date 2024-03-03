package hash

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	inp            = []byte{3, 3, 1, 4}
	chunkSize      = int64(3)
	firstChunkSum  = int64(195841)
	secondChunkSum = int64(195334)
)

func TestCheckSum(t *testing.T) {
	checksum := Checksum(inp[:3])
	require.Equal(t, checksum, firstChunkSum)
}

func TestChunkSlide(t *testing.T) {
	checksum := Checksum(inp[:3])
	newCS := ChunkSlide(checksum, inp[0], inp[len(inp)-1], chunkSize)
	require.Equal(t, newCS, secondChunkSum)
}
