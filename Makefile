build:
	GOOS=linux GOARCH=amd64 go build -o ./bin/gitlab-variable-linux-amd64 -ldflags="-s -w"  ./cmd/gitlab-variables/main.go
build_arm64:
	go build -o ./bin/gitlab-variable-arm64 -ldflags="-s -w"  ./cmd/gitlab-variables/main.go
