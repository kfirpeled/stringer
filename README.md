# Sample CLI app using cobra

```sh
❯ ./dist/stringer -h                   
stringer is a super fancy CLI (kidding)

One can use stringer to modify or inspect strings straight from the terminal

Usage:
  stringer [flags]
  stringer [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  inspect     Inspects a string
  reverse     Reverses a string

Flags:
  -h, --help      help for stringer
  -v, --version   version for stringer

Use "stringer [command] --help" for more information about a command.
```

# Run

```sh
go run main.go  -h
```

# Build

```sh
❯ go build -o ./dist/stringer -ldflags="-X 'github.com/kfirpeled/stringer/cmd/stringer.version=0.0.2'" main.go
```

# Testing

```sh
❯ go test -v ./...
?       github.com/kfirpeled/stringer   [no test files]
=== RUN   Test_ExecuteCommand
--- PASS: Test_ExecuteCommand (0.00s)
PASS
ok      github.com/kfirpeled/stringer/cmd/stringer      (cached)
?       github.com/kfirpeled/stringer/pkg/stringer      [no test files]
```

# Resources

- https://gianarb.it/blog/golang-mockmania-cli-command-with-cobra
- https://www.thorsten-hans.com/lets-build-a-cli-in-go-with-cobra/
