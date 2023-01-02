package main

import (
	"adventofcode/common"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"runtime/pprof"
	"strings"

	_ "adventofcode/2021"
	_ "adventofcode/2022"
)

func PrintListAutoWrap(w io.Writer, width int, list []string) {
	printed, _ := fmt.Fprintf(w, "  %s", list[0])
	for i := 1; i < len(list); i++ {
		name := list[i]
		if printed+len(name)+1 > width {
			printed, _ = fmt.Fprintf(w, ",\n  %s", name)
			printed -= 2
		} else {
			n, _ := fmt.Fprintf(w, ", %s", name)
			printed += n
		}
	}
	fmt.Fprintln(w)
}

const TERM_WIDTH = 80

func PrintYears(w io.Writer) {
	fmt.Fprint(w, "Available years:\n")
	PrintListAutoWrap(w, TERM_WIDTH, common.ListPrefixes())
}

func PrintSolutions(w io.Writer, prefix string) {
	solutions := common.ListSolutions(prefix)
	if len(solutions) == 0 {
		fmt.Fprintf(w, "No solutions for %s.\n", prefix)
		return
	}
	fmt.Fprintf(w, "Available solutions for %s:\n", prefix)
	PrintListAutoWrap(w, TERM_WIDTH, solutions)
}

func Usage() {
	w := flag.CommandLine.Output()
	fmt.Fprintf(w, "Usage: %[1]s [option...] [year/]<solution> [input]\n\n"+
		"Run the specified solution from the specified year.\n"+
		"If no year is given, defaults to %[2]s.\n"+
		"To list available solutions for a given year, omit the solution part.\n"+
		"  Example: %[1]s %[2]s/\n"+
		"\n"+
		"If no input file is given, attempts to search for an appropriate one.\n"+
		"If input file is -, reads from stdin.\n", os.Args[0], common.DefaultPrefix)
	fmt.Fprintf(w, "\nAvailable options:\n")
	flag.PrintDefaults()
	fmt.Fprintln(w)

	PrintYears(w)
	fmt.Fprintln(w)
	PrintSolutions(w, common.DefaultPrefix)
}

func IsTerminal(f *os.File) bool {
	fi, err := f.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice != 0
}

func FindInputFile(prefix, name string) string {
	name = common.NormalizeName(name)
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
	var (
		fShowPerformance bool
		fPprofFile       string
	)
	flag.Usage = Usage
	flag.BoolVar(&fShowPerformance, "p", false, "show performance information")
	flag.StringVar(&fPprofFile, "P", "", "output CPU profiling information")
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	prefix, solution, ok := common.SplitSolutionPrefix(flag.Arg(0))
	if !ok {
		flag.Usage()
		fmt.Fprintf(os.Stderr, "\nInvalid solution: %s\n", flag.Arg(0))
		os.Exit(1)
	}

	if solution == "" {
		PrintSolutions(os.Stdout, prefix)
		os.Exit(0)
	}

	fn, ok := common.GetSolution(prefix, solution)
	if !ok {
		flag.Usage()
		fmt.Fprintf(os.Stderr, "\nUnknown solution: %s\n", flag.Arg(0))
		os.Exit(1)
	}

	var r io.Reader
	if flag.NArg() == 1 {
		path := FindInputFile(prefix, solution)
		if path == "" {
			fmt.Fprintf(os.Stderr, "Cannot find input file for %s\n", flag.Arg(0))
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

	if fPprofFile != "" {
		f, err := os.Create(fPprofFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot create %s: %v\n", fPprofFile, err)
			os.Exit(1)
		}
		defer f.Close()
		pprof.StartCPUProfile(f)
	}
	pState := GetPerformanceState()
	fn(r)
	pInfo := DiffPerformanceState(pState)
	if fPprofFile != "" {
		pprof.StopCPUProfile()
	}

	if fShowPerformance {
		w := os.Stderr
		if IsTerminal(os.Stdout) && IsTerminal(os.Stderr) {
			fmt.Fprintln(w)
		}
		PrintPerformanceInfo(w, pInfo)
	}
}
