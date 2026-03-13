build:
  CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/script-mcp ./main.go

build-cli:
  CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/script-cli ./cmd/cli/

dev:
  go run main.go --sse_port 3004

install:
  go install ./...

install-cli:
  go install ./cmd/cli/
