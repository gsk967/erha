package signature

import (
	"os"

	"github.com/gobuffalo/logger"

	"github.com/gsk967/erha/fsoper"
)

// Generate ...
func Generate(file, signatureFile string, chunkSize int64) error {
	log := logger.NewLogger("info")
	fsOper := fsoper.New(chunkSize)
	reader, err := fsOper.Open(file)
	if err != nil {
		log.Fatalf("err while opening the file %s and err %v", file, err)
	}

	// Generate signature
	//
	gs := NewSigGenerator(*reader, chunkSize)
	sig := gs.GenSig()
	sigMarshal, err := sig.MarshalJSON()
	if err != nil {
		log.Fatalf("err while marshing the generated signature ,err %v", err)
	}

	// Writing signature to output file
	err = os.WriteFile(signatureFile, sigMarshal, 0666)
	if err != nil {
		log.Fatalf("Err while writting the signature to file %s and Err is %v", signatureFile, err)
	}

	log.Infof("Signature is written to %s", signatureFile)
	return nil
}
