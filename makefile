test-rules:
	go test -count=1 ./...

test-lint:
	go run cmd/loglint/main.go ./testdata/