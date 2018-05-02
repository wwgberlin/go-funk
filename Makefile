GO_FILES:=$(shell find . -name "*.go")

wavelet: $(GO_FILES)
	go build ./cmd/wavelet

test: $(GO_FILES)
	go test ./...

run-server: wavelet
	./wavelet -port 8080

