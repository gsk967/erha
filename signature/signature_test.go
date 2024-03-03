package signature

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/gsk967/erha/hash"
	"github.com/stretchr/testify/require"
)

func TestNewSigGenerator(t *testing.T) {
	data := []byte("hello world")
	chunkSize := int64(2)

	reader := bytes.NewReader(data)
	br := bufio.NewReader(reader)
	sg := NewSigGenerator(*br, int64(chunkSize))
	require.Equal(t, sg.size, chunkSize)
}

func TestGenSig(t *testing.T) {
	data := []byte("hello world")
	chunkSize := int64(2)

	reader := bytes.NewReader(data)
	br := bufio.NewReader(reader)
	sg := NewSigGenerator(*br, int64(chunkSize))
	sig := sg.GenSig()
	require.NotNil(t, sig.Checksums)
	require.Equal(t, hash.MD5Checksum(data[:chunkSize]), sig.Checksums[0].MD5Checksum)
}

func TestNextChunkHashes(t *testing.T) {
	data := []byte("hello world")
	chunkSize := int64(2)

	md5Sum := hash.MD5Checksum(data[:chunkSize])
	rhash := hash.Checksum(data[:chunkSize])

	reader := bytes.NewReader(data)
	br := bufio.NewReader(reader)
	sg := NewSigGenerator(*br, int64(chunkSize))
	hash, md5, err := sg.NextChunkHashes()
	require.Nil(t, err)
	require.Equal(t, hash, rhash)
	require.Equal(t, md5, md5Sum)
}
