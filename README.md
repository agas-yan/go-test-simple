# go-test-simple

## How to run this project

Normal Test

`go test`

Verbose:

`go test -v`

With Cover

`go test -cover`

``` sh
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

* * *

## Init go project tutorial

1. How to go mod init

    `go mod init {path}`

    example:

    `go mod init github.com/agas-yan/go-test-simple`

2. Get project
    
    `go get -u github.com/gorilla/mux`

3. Install dependencies

    `go mod vendor`

