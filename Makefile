test:
	go test -race ./...
.PHONY: test

plugin:
	go build -buildmode=plugin plugin/sqlnoctx.go
.PHONY: plugin
