# A simple Golang web setup for practicing CI/CD with Github Actions

- [Golang](https://go.dev/).
- [Github Actions](https://github.com/features/actions).

Build binary and run test locally:
```sh
go build -o bin/web ./cmd/web
go test -cover -v ./cmd/web
```

Run [`golangci-lint`](https://golangci-lint.run/) locally:
```sh
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.0.2
export PATH="$(go env GOPATH)/bin:$PATH"
golangci-lint --version
golangci-lint run
```

Build Docker image:
```sh
docker build -t go-web-cicd:latest .
```
