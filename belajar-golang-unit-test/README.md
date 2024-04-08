# README

To run the test:

```sh
# cd into specific package/directory, then
go test
# or this
go test -v
# or run specific one
go test -v -run=TestHelloWorld

# to run all unit tests in a project
go run ./...

# to run all unit tests + benchmark
go test -v -bench=.

# to skip unit tests and only run benchmark, specify filter that match no unit tests
go test -v -run=nothingness -bench=.

# to run specific benchmark
go test -v -run=nothingness -bench=BenchmarkHelloWorld

# to run all benchmark
go test -v -bench=. ./...
```
