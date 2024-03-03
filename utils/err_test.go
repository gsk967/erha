package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErrCheck(t *testing.T) {
	err := errors.New("new error")
	rerr := ErrCheck(err)
	require.NotNil(t, rerr)
}

func TestPanicErr(t *testing.T) {
	err := errors.New("new error")
	require.Panics(t, func() { PanicErr(err) }, "PanicErr did  panic with error")
}
