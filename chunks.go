package main

// func process() error {
// 	fsOper := fsoper.New(types.DefaultChunkSize)

// 	reader, err := fsOper.Open("1.txt")
// 	if err != nil {
// 		return err
// 	}
// 	// Read chunks from file
// 	block := make([]byte, types.DefaultChunkSize)
// 	// md5Hashes := make([]byte, types.DefaultChunkSize)
// 	for {
// 		// Add chunks to buffer
// 		bytesRead, err := reader.Read(block)
// 		// Stop if not bytes read or end to file
// 		if bytesRead == 0 || err == io.EOF {
// 			break
// 		}

// 		md5Hash := hash.MD5Checksum(block)
// 		fmt.Println(bytesRead, " hash ", md5Hash)
// 	}

// 	return nil
// }

// func CalMyHash(chunk []byte) {
// 	for i := 0; i < len(chunk); i++ {
// 		println("index i ", i, " => ", uint64(chunk[i]))
// 	}
// }
