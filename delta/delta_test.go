package delta

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDelta(t *testing.T) {
	delta := NewDelta()

	require.NotNil(t, delta)
	require.Nil(t, delta.Inserted)
	require.Nil(t, delta.Copied)
	require.Nil(t, delta.Deleted)
}

func TestDelta_MarshalJSON(t *testing.T) {
	delta := NewDelta()

	_, err := delta.MarshalJSON()
	require.Nil(t, err)
}

func TestUnmarshalJSON(t *testing.T) {
	delta := NewDelta()

	deltaJson, _ := delta.MarshalJSON()

	_, err := UnmarshalJSON(deltaJson)

	require.Nil(t, err)

	_, err = UnmarshalJSON(nil)

	require.NotNil(t, err)
}
