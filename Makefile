adventofcode: $(wildcard *.go) $(wildcard */*.go) go.mod
	go build -ldflags='-s -w'
