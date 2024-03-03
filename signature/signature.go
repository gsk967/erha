package signature

import (
	"bufio"
	"io"

	"github.com/gsk967/erha/hash"
	"github.com/gsk967/erha/utils"
)

type SigGenerator struct {
	reader bufio.Reader
	buf    []byte
	curPos int64
	size   int64
}

func NewSigGenerator(reader bufio.Reader, chunkSize int64) *SigGenerator {
	return &SigGenerator{
		reader: reader,
		buf:    make([]byte, chunkSize),
		curPos: 0,
		size:   chunkSize,
	}
}

func (sg *SigGenerator) GenSig() *Signature {
	s := &Signature{
		Checksums: nil,
		ChunkSize: sg.size,
	}

	for {
		rhash, md5Hash, err := sg.NextChunkHashes()
		if err != nil {
			if err == io.EOF {
				return s
			}
			utils.PanicErr(err)
		}
		ch := &Checksum{
			RChecksum:   rhash,
			MD5Checksum: md5Hash,
			Start:       sg.curPos - sg.size,
			End:         sg.curPos,
		}
		ch.Content = append(ch.Content, sg.buf...)
		s.Checksums = append(s.Checksums, ch)
	}
}

func (sg *SigGenerator) NextChunkHashes() (int64, [16]byte, error) {
	n, err := sg.reader.Read(sg.buf)
	if err != nil {
		return 0, [16]byte{}, err
	}
	sg.curPos += int64(n)
	rhash := hash.Checksum(sg.buf)
	return rhash, hash.MD5Checksum(sg.buf), nil
}
