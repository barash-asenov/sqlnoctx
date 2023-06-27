# sqlnoctx

![](https://github.com/barash-asenov/sqlnoctx/workflows/CI/badge.svg)

`sqlnoctx` finds usage of non context database/sql functions

You should use `sqlnoctx` if you want to enforce context aware methods on `database/sql` package

Benefits of using context-aware methods in the database/sql package in Go:

- Cancellation and Timeouts: Control query duration and prevent performance issues.
- Graceful Shutdown: Cleanly cancel ongoing operations during application shutdown.
- Monitoring and Logging: Attach contextual information for better tracing and analysis.
- Testing and Mocking: Enable independent unit testing by mocking database behavior.
- Middleware Support: Facilitate modular and extensible middleware patterns.

## Usage


### sqlnoctx with go vet

go vet is a Go standard tool for analyzing source code.

1. Install sqlnoctx.
```sh
$ go install github.com/barash-asenov/sqlnoctx/cmd/sqlnoctx@latest
```

2. sqlnoctx execute
```sh
$ go vet -vettool=`which sqlnoctx` main.go
./main.go:17:15: should use PingContext instead
```

### sqlnoctx with golangci-lint (as a private plugin)

golangci-lint is a fast Go linters runner.

sqlnoctx currently is not a part of public golangci-lint linter. But can be enabled to project still quite easily.  
The big drawback is that;
> `For a private linter (which acts as a plugin) to work properly, the plugin as well as the golangci-lint binary needs to be built for the same environment.`

Also;

> `golangci-lint` must be build with `CGO_ENABLED`

Source: https://golangci-lint.run/contributing/new-linters/#how-to-add-a-private-linter-to-golangci-lint

**Note**: Plugin is currently compatible with golangci-lint v1.53.3

1. Build golangci-lint with CGO_ENABLED=1.
[golangci-lint - Install](https://golangci-lint.run/usage/install/)

2. Clone and build sqlnoctx plugin.
```bash
git clone git@github.com:barash-asenov/sqlnoctx.git
cd sqlnoctx
make plugin
```

It will output `sqlnoctx.so` which can be moved root of the project directory (same level as .golangci-lint) that will be linted

2. Setup .golangci.yml
```yaml:
# Add sqlnoctx as custom linter
linters-settings:
  custom:
    sqlnoctx:
      path: ./sqlnoctx.so
      description: Sql no context
      original-url: https://github.com/barash-asenov/sqlnoctx

linters:
  disable-all: true
  enable:
    - sqlnoctx
```

3. sqlnoctx execute
```sh
# Use .golangci.yml
$ golangci-lint run

# Only sqlnoctx execute
golangci-lint run --disable-all -E sqlnoctx
```

## Detection rules
- Executing following functions
  - `(*database/sql.DB).Exec`,
  - `(*database/sql.DB).Ping`,
  - `(*database/sql.DB).Prepare`,
  - `(*database/sql.DB).Query`,
  - `(*database/sql.DB).QueryRow`,

## Reference
- [database/sql](https://pkg.go.dev/database/sql)
- [noctx](https://github.com/sonatard/noctx)
