# Unit Test in Go

## Run Tests

- go test unit-test/services

### Run specific func

- go test -run "AddMany"

### Show logs

- go test -v

### Benchmark

- go test -bench=.

### Subtest

- go test -run TestMultiply/2x1 -v
