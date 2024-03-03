.PHONY: test
test:
	go test -v ./...

.PHONY: tidy
tidy:
	go mod tidy

clean:
	rm -rf ${PWD}/build

.PHONY: build
build:
	go build -o ${PWD}/build/erha .

.PHONY: run
run: tidy
	echo "ℹ️ It will not write signature and delta of files into output files, Please use sub-commands for output"
	go run main.go 1.txt 2.txt