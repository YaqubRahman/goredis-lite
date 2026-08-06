# Mini Redis (Go + gRPC)

A simple in-memory key-value store built in Go, with a gRPC API. Supports `Set`, `Get`, and `Delete` over the network.

Made so I could learn Go, gRPC, and Protocol Buffers.

## Stack

- **Go** - server and client
- **gRPC** - service communication
- **Protocol Buffers** - API contract

## Project structure

```
├── grpc-service.proto   # the API contract
├── proto/               # generated Go code (do not edit)
├── main.go              # server
├── server.go            # Get/Set/Delete implementation
└── client/main.go       # example client
```

## Running it

Start the server:

```bash
go run .
```

In a second terminal, run the client:

```bash
go run ./client
```

The client sets a key, reads it back, and prints the result.

## Regenerating the proto code

After editing `grpc-service.proto`, regenerate the Go code:

```bash
protoc --go_out=. --go-grpc_out=. grpc-service.proto
```

## Roadmap

- [ ] Thread-safe map (concurrent access)
- [ ] TTL / key expiry
- [ ] Persistence to disk
