default: fmt install generate

build:
	go build -v ./...

install: build
	go install -v ./...

generate:
	cd tools; go generate ./...

fmt:
	gofmt -s -w -e .

test:
	go test -v -cover -timeout=120s -parallel=10 ./...

.PHONY: fmt test build install generate
