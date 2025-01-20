all: build clean

build:
	go mod tidy
	GOARCH=amd64 GOOS=linux go build -o farmapi cmd/farmapi.go
	GOARCH=arm64 GOOS=linux CGO_ENABLED=1 go build -o farmapi_arm cmd/farmapi.go

clean:
	go clean
	gofmt -w -s cmd/ internal/
