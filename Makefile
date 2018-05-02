GO_FILES:=$(shell find . -name "*.go")

wavelet: $(GO_FILES)
	go build ./cmd/wavelet

run-server: wavelet
	./wavelet -port 8080

