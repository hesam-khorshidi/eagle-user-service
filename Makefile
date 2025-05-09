run_http:
	@export PATH="$$PATH:$$(go env GOPATH)/bin" && go run cmd/main.go http

wire:
	@export PATH="$$PATH:$$(go env GOPATH)/bin" && wire ./cmd/app && wire ./cmd/migrator && echo "✅ done!"

build:
	@export PATH="$$PATH:$$(go env GOPATH)/bin" && go build -o user-service cmd/main.go