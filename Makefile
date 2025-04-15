.PHONY: lint
lint:
	@golangci-lint run -c .golangci.yaml ./...

.PHONY: test
test:
	@go test -v ./... \
		-test.parallel 2 \
		-test.timeout 5s

.PHONY: docs
docs:
	@godoc -http=:8080
