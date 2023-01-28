package common

import (
	"fmt"
	"io"
	"reflect"
	"runtime"
	"sort"
	"strings"
)

type SolutionFunc func(io.Reader, io.Writer)

var registry = make(map[string]map[string]SolutionFunc)
var DefaultPrefix = ""

func NormalizeName(name string) string {
	name = strings.ToLower(name)
	for _, c := range []string{" ", ".", "_"} {
		name = strings.ReplaceAll(name, c, "-")
	}
	return name
}

func SplitSolutionPrefix(name string) (string, string, bool) {
	parts := strings.Split(name, "/")
	if len(parts) > 2 {
		return "", "", false
	}
	if len(parts) == 1 {
		return DefaultPrefix, parts[0], true
	}
	return parts[0], parts[1], true
}

func GetFunctionName(i interface{}) string {
	// https://stackoverflow.com/a/7053871/5958455
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func RegisterSolution(prefix, name string, f SolutionFunc) {
	name = NormalizeName(name)
	if f, ok := registry[prefix][name]; ok {
		panic(fmt.Sprintf("Solution %s/%s already registered as %s", prefix, name, GetFunctionName(f)))
	}
	registry[prefix][name] = f
}

func GetPrefixRegistrar(prefix string) func(string, SolutionFunc) {
	if prefix > DefaultPrefix {
		DefaultPrefix = prefix
	}
	if _, ok := registry[prefix]; !ok {
		registry[prefix] = make(map[string]SolutionFunc)
	}
	return func(name string, f SolutionFunc) {
		RegisterSolution(prefix, name, f)
	}
}

func GetSolution(prefix, name string) (SolutionFunc, bool) {
	name = NormalizeName(name)
	f, ok := registry[prefix][name]
	return f, ok
}

func ListPrefixes() []string {
	names := make([]string, 0, len(registry))
	for name := range registry {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func ListSolutions(prefix string) []string {
	if _, ok := registry[prefix]; !ok {
		return nil
	}
	names := make([]string, 0, len(registry[prefix]))
	for name := range registry[prefix] {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
