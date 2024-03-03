package delta

import (
	"bufio"
	"encoding/json"

	"github.com/gobuffalo/logger"
	"github.com/gsk967/erha/signature"
)

type DeltaProcess struct {
	signature   *signature.Signature
	reader      *bufio.Reader
	delta       *Delta
	buf         []byte
	pos         int64
	sum         int64
	visited     []byte
	chunkSize   int64
	foundHashes map[int64]bool
	log         logger.Logger
}

type SingleDelta struct {
	RChecksum   int64    `json:"rchecksum,omitempty"`
	MD5Checksum [16]byte `json:"-"`
	Start       int64    `json:"start,omitempty"`
	End         int64    `json:"end,omitempty"`
	DiffBytes   []byte   `json:"diff,omitempty"`
}

type Delta struct {
	Inserted []SingleDelta `json:"inserted,omitempty"`
	Deleted  []SingleDelta `json:"deleted,omitempty"`
	Copied   []SingleDelta `json:"copied,omitempty"`
}

func NewDelta() *Delta {
	return &Delta{
		Inserted: nil,
		Deleted:  nil,
		Copied:   nil,
	}
}

func (d *Delta) MarshalJSON() ([]byte, error) {
	y, err := json.Marshal(*d)
	if err != nil {
		return nil, err
	}
	return y, nil
}

func UnmarshalJSON(bytes []byte) (*Delta, error) {
	s := &Delta{}
	err := json.Unmarshal(bytes, s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
