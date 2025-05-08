run_http:
	@export PATH="$$PATH:$$(go env GOPATH)/bin" && go run cmd/main.go http

wire:
	@export PATH="$$PATH:$$(go env GOPATH)/bin" && wire ./cmd/app && echo "✅ done!"