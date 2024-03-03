package delta

import (
	"os"

	"github.com/gobuffalo/logger"

	"github.com/gsk967/erha/fsoper"
	"github.com/gsk967/erha/signature"
)

// GenerateDelta
func GenerateDelta(sigFile, updatedFile string, deltaFile string) error {
	log := logger.NewLogger("info")
	sigData, err := fsoper.ReadFile(sigFile)
	if err != nil {
		log.Errorf("error while reading given signature file %s and err %v", sigFile, err)
		return err
	}

	sig, err := signature.UnmarshalJSON(sigData)
	if err != nil {
		return err
	}

	fsOper := fsoper.New(sig.ChunkSize)
	newFileReader, err := fsOper.Open(updatedFile)
	if err != nil {
		log.Errorf("error while reading new file %s and err %v", updatedFile, err)
		return err
	}

	dp := NewDeltaProcess(log, sig, newFileReader)
	delta := dp.BuildDelta()
	// show delta report
	GenerateDeltaReport(delta)

	// Write delta into deltaFile
	delteBytes, err := delta.MarshalJSON()
	if err != nil {
		return err
	}
	err = os.WriteFile(deltaFile, delteBytes, 0666)
	if err != nil {
		log.Errorf("err while writing the files delta into delta file : %s and err %v", deltaFile, err)
		return err
	}

	log.Infof("files delta is written to delta file :%s", deltaFile)

	return nil
}
