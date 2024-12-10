all: build clean

build:
	go mod tidy
	GOARCH=amd64 GOOS=linux go build -o farmapi cmd/farmapi.go

clean:
	go clean
	gofmt -w -s cmd/ internal/
