# go-toolbox

Small generic Go packages: a double-ended queue and slice/option helpers.

## Install

```bash
go get github.com/slashvar/go-toolbox@latest
```

Requires Go 1.25 or later.

## Packages

| Path                           | Content                                    |
|--------------------------------|--------------------------------------------|
| [`pkg/deque`](pkg/deque)       | Ring-buffer double-ended queue             |
| [`pkg/utils`](pkg/utils)       | Generic slice helpers and an optional type |
| [`examples/bfs`](examples/bfs) | Breadth-first tree traversal using `deque` |

## Development

```bash
go build ./...
go test ./...
golangci-lint run
go run ./examples/bfs
```

CI runs build and test on every push and pull request to `main`.

## License

BSD 2-Clause. See [LICENSE](LICENSE).
