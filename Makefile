.PHONY: fmt test vet lint check

fmt:
	go fmt ./...

test:
	go test ./... -v

vet:
	go vet ./...

lint:
	golangci-lint run

check: fmt vet lint test
