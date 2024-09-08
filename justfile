up:
	go mod download

test:
  go test -v ./...

build:
  go build -o ./bin/ ./cmd/go-wc/*.go

aider:
	ANTHROPIC_API_KEY=$(cat .anthropic_key) aider --sonnet