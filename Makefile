GO_FILES:=$(shell find . -name "*.go")

go-funk: $(GO_FILES)
	go build -o $@ .

.PHONY: test
test: $(GO_FILES)
	go test -v ./...

.PHONY: run-server
run-server: go-funk test
	./go-funk -port 8080

