# DUP

An utility written in go to find duplicate lines in a file or the stdin.

Lines must have at *most* `65536` characters.

# Dev

Build with `go build`

Install with `go install`

# Usage

Install it with `go get github.com/abilioesteves/dup`.

You can find usage details by typing `./dup --help`.

## File

To read from file, type:

```go
./dup -p <file path>
```

## STDIN

To read from the `STDIN`, type:

```go
./dup
```
And then type or paste your lines at the terminal.

The results will be printed after terminating the program with `Ctrl+C`.



