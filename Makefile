test:
	bash -c "set -m; bash '$(CURDIR)/scripts/test.sh'"

lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0 run

format:
	go mod tidy
	go fmt ./...
	go run github.com/daixiang0/gci@latest write \
		--skip-generated \
		-s standard -s default \
		-s "prefix(github.com/a-novel-kit/test-utils)" \
		.
	go run mvdan.cc/gofumpt@latest -l -w .

PHONY: test lint format
