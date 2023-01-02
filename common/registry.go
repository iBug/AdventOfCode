package common

import (
	"fmt"
	"io"
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

func GetSolution(name string) (SolutionFunc, bool) {
	name = NormalizeName(name)
	f, ok := registry[name]
	return f, ok
}

func ListSolutions() []string {
	names := make([]string, 0, len(registry))
	for name := range registry {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
