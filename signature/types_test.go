package signature

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMarshalJson(t *testing.T) {
	sig := Signature{
		Checksums: nil,
		ChunkSize: 1,
	}
	r, err := sig.MarshalJSON()
	require.NoError(t, err)
	require.NotNil(t, r)
}

func TestUnmarshalJson(t *testing.T) {
	sig := Signature{
		Checksums: nil,
		ChunkSize: 1,
	}
	r, err := sig.MarshalJSON()
	require.NoError(t, err)
	nsig, err := UnmarshalJSON(r)
	require.NoError(t, err)
	require.Equal(t, nsig.ChunkSize, sig.ChunkSize)
}
