# sqlnoctx

![](https://github.com/barash-asenov/sqlnoctx/workflows/CI/badge.svg)

`sqlnoctx` finds usage of non context database/sql functions

You should use `noctx` if sending http request in your library.
Passing `context.Context` enables library user to cancel http request, getting trace information and so on.

## Usage


### noctx with go vet

go vet is a Go standard tool for analyzing source code.

1. Install noctx.
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
The big drawback is that `For a private linter (which acts as a plugin) to work properly, the plugin as well as the golangci-lint binary needs to be built for the same environment.`
Also `golangci-lint` must be build with `CGO_ENABLED`

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

3. noctx execute
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
