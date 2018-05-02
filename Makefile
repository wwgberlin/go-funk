GO_FILES:=$(shell find . -name "*.go")

wavelet: $(GO_FILES)
	go build ./cmd/wavelet

.PHONY: test
test: $(GO_FILES)
	go test ./...

.PHONY: run-server
run-server: wavelet test
	./wavelet -port 8080

