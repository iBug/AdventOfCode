package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"time"
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

func GetFunctionName(i interface{}) string {
	// https://stackoverflow.com/a/7053871/5958455
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func RegisterSolution(name string, f SolutionFunc) {
	name = NormalizeName(name)
	if f, ok := registry[name]; ok {
		panic(fmt.Sprintf("Solution %s already registered as %s", name, GetFunctionName(f)))
	}
	registry[name] = f
}

func Usage() {
	solutions := make([]string, 0, len(registry))
	for name := range registry {
		solutions = append(solutions, name)
	}
	sort.Strings(solutions)

	w := flag.CommandLine.Output()
	fmt.Fprintf(w, "Usage: %s [option...] <solution> [input...]\n", os.Args[0])
	fmt.Fprintf(w, "\nAvailable options:\n")
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Fprintf(w, "  -%s\t%s\n", f.Name, f.Usage) // f.Name, f.Value
	})

	const TARGET = 80

	fmt.Fprintf(w, "\nAvailable solutions:\n")
	printed, _ := fmt.Fprintf(w, "  %s", solutions[0])
	for i := 1; i < len(solutions); i++ {
		name := solutions[i]
		if printed+len(name)+1 > TARGET {
			printed, _ = fmt.Fprintf(w, ",\n  %s", name)
			printed -= 2
		} else {
			n, _ := fmt.Fprintf(w, ", %s", name)
			printed += n
		}
	}
	fmt.Fprintln(w)
}

var fShowPerformance bool

func main() {
	flag.Usage = Usage
	flag.BoolVar(&fShowPerformance, "p", false, "show performance information")
	flag.Parse()
	if flag.NArg() < 1 {
		Usage()
		os.Exit(1)
	}

	solution := NormalizeName(flag.Arg(0))
	fn, ok := registry[solution]
	if !ok {
		Usage()
		fmt.Fprintf(os.Stderr, "Unknown solution: %s\n", solution)
		os.Exit(1)
	}
	rs := []io.Reader{}
	for _, path := range flag.Args()[1:] {
		f, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot open %s: %v\n", path, err)
			os.Exit(1)
		}
		defer f.Close()
		rs = append(rs, f)
	}
	if len(rs) == 0 {
		rs = append(rs, os.Stdin)
	}

	startTime := time.Now()
	fn(io.MultiReader(rs...))
	duration := time.Since(startTime)

	if fShowPerformance {
		fmt.Fprintf(os.Stderr, "\nDuration: %s\n", duration)
	}
}
