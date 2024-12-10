all: build clean

build:
	GOARCH=amd64 GOOS=linux go build -o farmapi cmd/farmapi.go

clean:
	go clean
	go mod tidy
	gofmt -w -s cmd/ internal/
