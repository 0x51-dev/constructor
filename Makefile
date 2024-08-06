.PHONY: test fmt

# Runs the tests.
test:
	go test -v -cover ./...

# Formats the code and runs the linter.
fmt:
	go mod tidy
	gofmt -s -w .
	goarrange run -r .
	golangci-lint run ./...
