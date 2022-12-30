adventofcode: $(wildcard *.go) go.mod
	go build -ldflags='-s -w'
