# ERHA - Eigner rolling hash algorithm

This rolling hash based file diffing algorithm will return when comparing original and an updated version of an input, it should return a description ("delta") which can be used to upgrade an original version of the file into the new file.

## Test 

```bash
$ make test
```

## How to run ?

### Build the binary

```bash
$ make build 
```

### ERHA file diff algorithm 
    // Note: It will not write the signature and delta of files into output files.
    // For this, please use subcommands of erha.
    $ go run main.go original-file signature-output-fule --chunk_size CHUNK_SIZE 

```bash
Example:
$ go run main.go 1.txt 2.txt --chunk_size 4
>  [ 0 - 4 ]  sai 
>  [ 4 - 8 ]  kuma
-  [ 5 - 9 ]  ruma
+  [ 8 - 21 ]  r and new one
```
    
### Generate signature
    $ go run main.go signature original-file signature-output-fule --chunk_size CHUNK_SIZE 
    // default CHUNK_SIZE is 2 

```bash
Example:
$ go run main.go signature 1.txt sig.txt --chunk_size 2
// or 
$ ./build/erha signature 1.txt sig.txt --chunk_size 2
// output
$ go run main.go signature 1.txt sig.txt --chunk_size 4      
INFO[2024-03-03T16:59:03+05:30] Signature is written to sig.txt
```

### Delta of the original and new file
    $ go run main.go delta signature-file new-file deleta-output-file --chunk_size CHUNK_SIZE 
    
```bash
$ go run main.go delta sig.txt 2.txt delta.txt
// or 
$ ./build/erha delta sig.txt 2.txt delta.txt

// Output 
$ go run main.go delta sig.txt 2.txt delta.txt --chunk_size 4
>  [ 0 - 4 ]  sai 
>  [ 4 - 8 ]  kuma
-  [ 5 - 9 ]  ruma
+  [ 8 - 21 ]  r and new one
INFO[2024-03-03T16:58:38+05:30] files delta is written to delta file :delta.txt
```