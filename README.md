### Common Commands

```bash
# format code, VSC should be doing this automatically
go fmt

# run tests
go test

# run tests and show coverage
go test -cover

# run tests and output coverage file
go test -coverprofile=coverage.out

# pretty print coverage file
go tool cover -html=coverage.out

# AIO coverage file and pretty print
go test -coverprofile=coverage.out && go tool cover -html=coverage.out

# get some module
go get "github.com/google/go-cmp/cmp"

# cleanup go.mod for direct references
go mod tidy
```
