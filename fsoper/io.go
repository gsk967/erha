/*
FileSystem operations
*/
package fsoper

import (
	"bufio"
	"errors"
	"math"
	"os"
)

type FSOperations struct {
	chunksize int64
}

func New(chunksize int64) FSOperations {
	return FSOperations{
		chunksize: chunksize,
	}
}

func (fsOper FSOperations) Open(filePath string) (*bufio.Reader, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	// Get file info and get total file size
	fInfo, _ := f.Stat()
	fSize := fInfo.Size()
	// Calculate file chunks availables
	fChunks := fsOper.Chunks(fSize)
	// Check if at least two chunks are generated based on file size and block size
	if fChunks <= 1 {
		return nil, errors.New("at least 2 chunks are required for operations")
	}
	return bufio.NewReader(f), err
}

// Returning no of chunks based on file size
func (fsOper FSOperations) Chunks(fileSize int64) int {
	return int(math.Ceil(float64(fileSize) / float64(fsOper.chunksize)))
}

// ReadFile will given read file and return its contents
func ReadFile(fileName string) ([]byte, error) {
	return os.ReadFile(fileName)
}
