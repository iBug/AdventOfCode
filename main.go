package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type SolutionFunc func(io.Reader)

var registry = make(map[string]SolutionFunc)

func NormalizeName(name string) string {
	name = strings.ToLower(name)
	for _, c := range []string{" ", ".", "_"} {
		name = strings.ReplaceAll(name, c, "-")
	}
	return name
}

func RegisterSolution(name string, f SolutionFunc) {
	registry[NormalizeName(name)] = f
}

func Usage() {
	solutions := make([]string, 0, len(registry))
	for name := range registry {
		solutions = append(solutions, name)
	}
	sort.Strings(solutions)

	w := flag.CommandLine.Output()
	fmt.Fprintf(w, "Usage of %s: [options] solution [input...]\n", os.Args[0])
	fmt.Fprintf(w, "Available solutions:\n")
	for _, name := range solutions {
		fmt.Fprintf(w, "\t%s\n", name)
	}
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		Usage()
		os.Exit(1)
	}

	solution := NormalizeName(flag.Arg(0))
	fn, ok := registry[solution]
	if !ok {
		Usage()
		os.Exit(1)
	}
	rs := []io.Reader{}
	for _, path := range flag.Args()[1:] {
		f, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot open %s: %v", path, err)
			os.Exit(1)
		}
		defer f.Close()
		rs = append(rs, f)
	}
	if len(rs) == 0 {
		rs = append(rs, os.Stdin)
	}

	fn(io.MultiReader(rs...))
}
