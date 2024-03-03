package delta

import (
	"bufio"
	"io"

	"github.com/gobuffalo/logger"

	"github.com/gsk967/erha/hash"
	"github.com/gsk967/erha/signature"
	"github.com/gsk967/erha/utils"
)

func NewDeltaProcess(log logger.Logger, sig *signature.Signature, reader *bufio.Reader) *DeltaProcess {
	if sig.ChunkSize == 0 {
		log.Fatalf("invalid chunksize ", sig.ChunkSize)
	}
	return &DeltaProcess{
		reader:      reader,
		signature:   sig,
		buf:         make([]byte, sig.ChunkSize),
		delta:       NewDelta(),
		chunkSize:   sig.ChunkSize,
		visited:     make([]byte, 0),
		pos:         0,
		foundHashes: make(map[int64]bool),
		log:         log,
	}
}

func (dp *DeltaProcess) BuildDelta() *Delta {
	rChecksumsMap := map[int64][16]byte{}
	for _, checksum := range dp.signature.Checksums {
		rChecksumsMap[checksum.RChecksum] = checksum.MD5Checksum
	}

	for {
		err := dp.Roll(rChecksumsMap)
		if err != nil {
			if err == io.EOF {
				break
			}
			utils.PanicErr(err)
		}
		dp.foundHashes[dp.sum] = true
	}

	for _, chunk := range dp.signature.Checksums {
		var d SingleDelta
		if _, ok := dp.foundHashes[chunk.RChecksum]; !ok {
			d = SingleDelta{
				RChecksum:   chunk.RChecksum,
				MD5Checksum: chunk.MD5Checksum,
				Start:       chunk.Start,
				End:         chunk.End,
			}
			if len(chunk.Content) != 0 {
				d.DiffBytes = append(d.DiffBytes, chunk.Content...)
			}
			dp.delta.Deleted = append(dp.delta.Deleted, d)
		}
	}

	return dp.delta
}

func (dp *DeltaProcess) Roll(checksums map[int64][16]byte) error {
	n, err := dp.reader.Read(dp.buf)
	dp.pos += int64(n)
	if err != nil {
		if err == io.EOF {
			dp.SetRemainingBytes()
		}
		return err
	}

	dp.sum = hash.Checksum(dp.buf)
	if md5Sum, ok := checksums[dp.sum]; ok {
		if ok := dp.AppendDiff(md5Sum); ok {
			return nil
		}
	}

	dp.visited = append(dp.visited, dp.buf...)

	for {
		err = dp.Next()
		dp.foundHashes[dp.sum] = true

		if err != nil {
			if err == io.EOF {
				dp.SetRemainingBytes()
				return err
			}
			dp.log.Printf("err while logging the window %s", err.Error())
		}

		if md5Sum, ok := checksums[dp.sum]; ok {
			if dp.visited != nil {
				dp.delta.Inserted = append(dp.delta.Inserted, SingleDelta{
					Start:     dp.pos - int64(len(dp.visited)),
					End:       dp.pos,
					DiffBytes: dp.visited,
				})
				dp.visited = nil
			}
			dp.AppendDiff(md5Sum)
			return nil
		}
		dp.visited = append(dp.visited, dp.buf[len(dp.buf)-1])
	}
}

func (dp *DeltaProcess) AppendDiff(md5Hash [16]byte) bool {
	newMD5Hash := hash.MD5Checksum(dp.buf)

	if md5Hash == newMD5Hash {
		d := SingleDelta{
			RChecksum:   dp.sum,
			MD5Checksum: newMD5Hash,
			Start:       dp.pos - dp.chunkSize,
			End:         dp.pos,
		}

		if len(dp.buf) != 0 {
			d.DiffBytes = append(d.DiffBytes, dp.buf...)
		}
		dp.delta.Copied = append(dp.delta.Copied, d)
		return true
	}

	return false
}

func (dp *DeltaProcess) Next() error {
	left := dp.buf[0]
	right := make([]byte, 1)

	n, err := dp.reader.Read(right)
	dp.pos += int64(n)
	if err != nil {
		return err
	}

	dp.sum = hash.ChunkSlide(dp.sum, left, right[0], dp.chunkSize)
	dp.buf = append(dp.buf[1:], right[0])
	return nil
}

func (dp *DeltaProcess) SetRemainingBytes() {
	if len(dp.visited) != 0 {
		dp.delta.Inserted = append(dp.delta.Inserted, SingleDelta{
			Start:     dp.pos - int64(len(dp.visited)),
			End:       dp.pos,
			DiffBytes: dp.visited,
		})
	}
}
