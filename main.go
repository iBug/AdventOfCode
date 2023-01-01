package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"reflect"
	"runtime"
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
	fmt.Fprintf(w, "Usage: %s [option...] <solution> [input]\n\n", os.Args[0])
	fmt.Fprint(w, "Run the specified solution.\n"+
		"If no input file is given, attempts to search for an appropriate one.\n"+
		"If input file is -, reads from stdin.\n")
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

func IsTerminal(f *os.File) bool {
	fi, err := f.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice != 0
}

func FindInputFile(name string) string {
	name = strings.SplitN(name, "-", 2)[0]
	candidates := make([]string, 0, 12)
	for _, dir := range []string{"inputs", "input", ""} {
		for _, ext := range []string{".txt", ".in"} {
			candidates = append(candidates, path.Join(dir, "input-"+name+ext))
			candidates = append(candidates, path.Join(dir, name+ext))
		}
		candidates = append(candidates, path.Join(dir, "day"+name, "input.txt"))
		candidates = append(candidates, path.Join(dir, "input-"+name))
	}

	for _, c := range candidates {
		if _, err := os.Stat(c); err == nil {
			return c
		}
	}
	return ""
}

func main() {
	var fShowPerformance bool
	flag.Usage = Usage
	flag.BoolVar(&fShowPerformance, "p", false, "show performance information")
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	solution := NormalizeName(flag.Arg(0))
	fn, ok := registry[solution]
	if !ok {
		flag.Usage()
		fmt.Fprintf(os.Stderr, "\nUnknown solution: %s\n", solution)
		os.Exit(1)
	}

	var r io.Reader
	if flag.NArg() == 1 {
		path := FindInputFile(solution)
		if path == "" {
			fmt.Fprintf(os.Stderr, "Cannot find input file for %s\n", solution)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "Found input file: %s\n", path)
		f, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot open %s: %v\n", path, err)
			os.Exit(1)
		}
		defer f.Close()
		r = f
	} else {
		if flag.Arg(1) == "-" {
			r = os.Stdin
		} else {
			path := flag.Arg(1)
			f, err := os.Open(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Cannot open %s: %v\n", path, err)
				os.Exit(1)
			}
			defer f.Close()
			r = f
		}
	}

	pState := GetPerformanceState()
	fn(r)
	pInfo := DiffPerformanceState(pState)

	if fShowPerformance {
		w := os.Stderr
		if IsTerminal(os.Stdout) && IsTerminal(os.Stderr) {
			fmt.Fprintln(w)
		}
		PrintPerformanceInfo(w, pInfo)
	}
}
