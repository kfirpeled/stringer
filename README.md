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

# Distribute

```sh
❯ goreleaser release --snapshot --rm-dist
  • starting release...
  • loading config file                              file=.goreleaser.yaml
  • loading environment variables
  • getting and validating git state
    • ignoring errors because this is a snapshot     error=git doesn't contain any tags. Either add a tag or use --snapshot
    • building...                                    commit=1b761675d72ffe832b2655a5a4f15697c212b49f latest tag=v0.0.0
    • pipe skipped                                   reason=disabled during snapshot mode
  • parsing tag
  • setting defaults
  • running before hooks
    • running                                        hook=go mod tidy
    • running                                        hook=go generate ./...
  • snapshotting
    • building snapshot...                           version=0.0.1-next
  • checking distribution directory
  • loading go mod information
  • build prerequisites
  • writing effective config file
    • writing                                        config=dist/config.yaml
  • building binaries
    • building                                       binary=dist/stringer_darwin_arm64/stringer
    • building                                       binary=dist/stringer_linux_amd64_v1/stringer
    • building                                       binary=dist/stringer_linux_arm64/stringer
    • building                                       binary=dist/stringer_windows_386/stringer.exe
    • building                                       binary=dist/stringer_windows_amd64_v1/stringer.exe
    • building                                       binary=dist/stringer_windows_arm64/stringer.exe
    • building                                       binary=dist/stringer_linux_386/stringer
    • building                                       binary=dist/stringer_darwin_amd64_v1/stringer
    • took: 1s
  • archives
    • creating                                       archive=dist/stringer_0.0.1-next_Darwin_arm64.tar.gz
    • creating                                       archive=dist/stringer_0.0.1-next_Windows_i386.tar.gz
    • creating                                       archive=dist/stringer_0.0.1-next_Windows_arm64.tar.gz
    • creating                                       archive=dist/stringer_0.0.1-next_Linux_x86_64.tar.gz
    • creating                                       archive=dist/stringer_0.0.1-next_Linux_arm64.tar.gz
    • creating                                       archive=dist/stringer_0.0.1-next_Linux_i386.tar.gz
    • creating                                       archive=dist/stringer_0.0.1-next_Darwin_x86_64.tar.gz
    • creating                                       archive=dist/stringer_0.0.1-next_Windows_x86_64.tar.gz
  • calculating checksums
  • storing release metadata
    • writing                                        file=dist/artifacts.json
    • writing                                        file=dist/metadata.json
  • release succeeded after 1s
```

# Resources

- https://gianarb.it/blog/golang-mockmania-cli-command-with-cobra
- https://www.thorsten-hans.com/lets-build-a-cli-in-go-with-cobra/
