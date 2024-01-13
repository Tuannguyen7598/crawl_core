# Go Crawl Core

## System diagram

![images](infrastructure.png)

## Features

- [x] Package manager: go mod
- [x] Routes: gorilla/mux
- [x] Environment variables, config.go
- [x] Base Module
- [x] Base Rest api
- [x] Base gRPC server, gRPC client
- [x] Database access layer
- [x] DDD architecture
- [x] Docker compose build infrastructure with replica mongo
- [x] Run air with development

## Folders structure

- `cmd/`:Controllers for web requests processing.
- `config/`: Config variables.
- `domain/`: Bussiness logic.
- `internal/`: Package internal of project.
- `pkg/`: Package of github.
- `tmp/`: file main and logging.
- `vendor/`: Packages used in application.

## How to use

Check Go version:

```bash
$ go version
```

### Download vendor packages

```bash
$ go mod download
```

### Run application

```bash
## Mode development
$ air

## Mode production
$ go run main.go
```

Check http://localhost:6666 to Http server
Check http://localhost:50052 to gRPC server

## Requirements

- Go 1.12+
