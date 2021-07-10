buildDateTime = $(shell date '+%Y-%m-%d %H:%M:%S')
gitCommitCode = $(shell git rev-list --full-history --all --abbrev-commit --max-count 1)
goVersion = $(shell go version)

run:	build
	 ./deployments/server
build:	apidoc
	go build -tags "apidoc manager" -o ./deployments/server cmd/main.go
air:
	air
apidoc:
	apidoc -i api -f ".go" -o web/apidoc

all:	apidoc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -tags manager -ldflags "-X 'main.buildDateTime=$(buildDateTime)' -X 'main.gitCommitCode=$(gitCommitCode)' -X 'main.goVersion=${goVersion}' -s -w" -o ./deployments/server cmd/main.go
	upx -9 ./deployments/server
