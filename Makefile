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

.PHONY: publish
publish:
	GOPROXY=proxy.golang.org go list -m github.com/OpenRunic/framed@${tag}
