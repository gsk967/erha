package cmd

import (
	"github.com/gobuffalo/logger"

	"github.com/gsk967/erha/delta"
	"github.com/gsk967/erha/fsoper"
	"github.com/gsk967/erha/signature"
)

// RunERHA will process the signatures of given original file and compare the new file with original file signature
func RunERHA(originalFile, newFile string, chunkSize int64) error {
	log := logger.NewLogger("info")
	fsOper := fsoper.New(chunkSize)
	reader, err := fsOper.Open(originalFile)
	if err != nil {
		log.Fatalf("err while opening the file %s and err %v", originalFile, err)
	}

	// Generate signature
	//
	gs := signature.NewSigGenerator(*reader, chunkSize)
	sig := gs.GenSig()

	newFileReader, err := fsOper.Open(newFile)
	if err != nil {
		log.Fatalf("err while opening the new  file %s and err %v", newFile, err)
	}

	// Delta for files
	//
	dp := delta.NewDeltaProcess(log, sig, newFileReader)
	filesDelta := dp.BuildDelta()
	// show delta report
	delta.GenerateDeltaReport(filesDelta)

	return nil
}
