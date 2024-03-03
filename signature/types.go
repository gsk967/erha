package signature

import "encoding/json"

// Signature of the file
type Signature struct {
	Checksums []*Checksum `json:"checksums"`
	ChunkSize int64       `json:"chunk_size"`
}

func (s *Signature) MarshalJSON() ([]byte, error) {
	return json.Marshal(*s)
}

func UnmarshalJSON(bytes []byte) (*Signature, error) {
	var s Signature
	err := json.Unmarshal(bytes, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// Checksum
type Checksum struct {
	RChecksum   int64    `json:"r_checksum"`
	MD5Checksum [16]byte `json:"md5_checksum"`
	Start       int64    `json:"start"`
	End         int64    `json:"end"`
	Content     []byte   `json:"content"`
}
