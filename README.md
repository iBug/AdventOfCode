# AdventOfCode

My solutions for [Advent Of Code 2022](https://adventofcode.com/2022/) challenges.

## Features

- Written in Go
- Standard library only
- All solutions are compiled into a single binary
- All solutions take reasonable time and memory usage ([performance record](records.txt))
- Built-in execution time measurement, though memory tracking is only supported on Linux

## Usage

Compilation is straightforward: `make` (or just `go build`)

Built-in program help:

```text
Usage: ./adventofcode [option...] <solution> [input...]

Available options:
  -p    show performance information

Available solutions:
  1-1, 1-2, 10-1, 10-2, 11-1, 11-2, 12-1, 12-2, 13-1, 13-2, 14-1, 14-2, 15-1,
  15-2, 16-1, 16-2, 17-1, 17-2, 18-1, 18-2, 19-1, 19-2, 2-1, 2-2, 20-1, 20-2,
  21-1, 21-2, 22-1, 22-2, 22-2e, 23-1, 23-2, 24, 24-1, 24-2, 25, 25-1, 3-1, 3-2,
  4-1, 4-2, 5-1, 5-2, 6-1, 6-2, 7, 7-1, 7-2, 8-1, 8-2, 9-1, 9-1a, 9-2
```

If no input file is specified, it uses standard input.
