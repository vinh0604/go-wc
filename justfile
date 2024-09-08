up:
	go mod download

test:
  go test -v ./...

aider:
	ANTHROPIC_API_KEY=$(cat .anthropic_key) aider --sonnet