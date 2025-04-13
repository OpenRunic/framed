.PHONY: lint
lint:
	@golangci-lint run -c .golangci.yaml ./...

.PHONY: test
test:
	@go test -v ./... \
		-test.parallel 2 \
		-test.timeout 5s
