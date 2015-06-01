default: deps test cover benchmark

vet:
	go vet ./...

deps:
	go get -d -v ./...
	go list -f '{{range .TestImports}}{{.}} {{end}}' ./... | xargs -n1 go get -d

test:
	go test ./...

cover:
	go test ./... --cover

benchmark:
	go test ./... -bench .

.PHONY: deps vet test cover benchmark