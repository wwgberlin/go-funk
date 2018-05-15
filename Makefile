GO_FILES:=$(shell find . -name "*.go")

wavelet: $(GO_FILES)
	go build .

.PHONY: test
test: $(GO_FILES)
	go test -v ./...

.PHONY: run-server
run-server: go-funk test
	./go-funk -port 8080

