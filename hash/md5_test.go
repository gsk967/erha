package hash

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMD5Checksum(t *testing.T) {
	inp := []byte("asdasd")
	md5Hash := MD5Checksum(inp)
	require.Equal(t, hex.EncodeToString(md5Hash[:]), "a8f5f167f44f4964e6c998dee827110c")
}
