<h1 align="center">
  <a href="https://github.com/zoobc/zoobc-core">
    ZooBC Core
  </a>
</h1>
<p align="center">
  <a href="https://circleci.com/gh/zoobc/zoobc-core">
    <img src="https://circleci.com/gh/zoobc/zoobc-core.svg?style=svg&circle-token=cdd770bcb30a201696bb10e76ed15504cf235a9f" alt="CircleCI"/>
  </a>
</p>

Zoobc-core is the main node application to run the zoobc blockchain. This repository consist of the main node application and the `command line interface` tools to help with development, which is located in the `cmd/` directory.

### Environment

- [golang](https://golang.org/doc/install), currently using go.1.12
- [go dep](https://golang.github.io/dep/docs/installation.html), currently using v.0.5.1
- [go-swagger](https://github.com/go-swagger/go-swagger) optional. Used as tools to document the rpc endpoint.
- [gopherbadger](https://github.com/jpoles1/gopherbadger) optional. Used to calculate total test coverage.
- [protoc](https://github.com/protocolbuffers/protobuf), optional as we are pushing the generated go file to the repo.
- [protoc-gen-go](https://github.com/golang/protobuf), optional as we are pushing the generated go file to the repo.

### Install

- clone the repository.
- run `dep ensure -v --vendor-only` to install the dependencies.
- run `git submodule update --init --recursive --remote` to update / fetch submodule.

### Run

- Main node application
  ```
  go run main.go
  ```
  Flags:
  1. `debug` (bool): enters debug mode with capabilities like `prometheus monitoring` (uses port defined by `monitoringPort` in the config file).
  2. `config-postfix` (string): defines which config file with defined postfix to use. It will use the config file in `./resource/config{postfix}`.toml.
  3. `config-path` (string): defines the directory that will hold the resources and configs. by default it will use `./resource`. This will be useful particularly for robust resource placements, for example while used in Zoobc-Testing-Framework.
- Command line tools
  ```
  cd cmd
  go run main.go --help
  go run main.go tx generate -t registerNode
  ```

### Build

```
go build -o zoobc
```

### Tests

- #### unit test

  - run all tests `go test ./...`
  - run all test with coverage report `go test ./... -coverprofile=cover.out && go tool cover -html=cover.out`

- #### Lint
  - run `golangci-lint ./...` to check any linting error in the changes.
  - remember to run tests, and lint before submitting PR.

### Swagger

- install
- pull newest `schema` and run `./compile-go.sh` to recompile the go file and produce swagger definition from it.

### Contributing

please refer to [contribute.md](contribute.md) and [code of conduct](code_of_conduct.md).

### GRPC web proxy for browser

[link](https://github.com/improbable-eng/grpc-web/tree/master/go/grpcwebproxy)

```
nohup grpcwebproxy --backend_addr=localhost:7000 --run_tls_server=false --allow_all_origins --server_http_debug_port=7001 --server_http_max_write_timeout 1h &
```
