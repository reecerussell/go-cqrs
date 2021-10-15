[![Go Report Card](https://goreportcard.com/badge/github.com/reecerussell/go-cqrs)](https://goreportcard.com/badge/github.com/reecerussell/go-cqrs)
[![codecov](https://codecov.io/gh/reecerussell/go-cqrs/branch/master/graph/badge.svg)](https://codecov.io/gh/reecerussell/go-cqrs)
[![Go Docs](https://godoc.org/github.com/reecerussell/go-cqrs?status.svg)](https://godoc.org/github.com/reecerussell/go-cqrs)

# CQRS

This is a basic commanding package with interfaces used to command and query data, which aid the separation of concerns. Commands and queries are executed using Executors such as `QueryExecutor` or `CommandExecutor` which enforce every command to be executed the same.
