package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

func init() {
	RegisterSolution("7", Solution7)
	RegisterSolution("7-1", Solution7)
	RegisterSolution("7-2", Solution7)
}

type Inode struct {
	mode     byte
	size     int
	parent   *Inode
	children map[string]*Inode
}

const (
	M_FILE = iota
	M_DIR
)

func DirSizeRecurse_7(dir *Inode) int {
	if dir.size > 0 {
		return dir.size
	}

	total := 0
	for _, inode := range dir.children {
		if inode.mode == M_FILE {
			total += inode.size
		} else {
			total += DirSizeRecurse_7(inode)
		}
	}
	dir.size = total
	return total
}

func RecurseTotalSize_7(dir *Inode, thresh int) int {
	total := 0
	for _, inode := range dir.children {
		if inode.mode == M_FILE {
			continue
		}
		size := DirSizeRecurse_7(inode)
		if size <= thresh {
			total += size
		}
		total += RecurseTotalSize_7(inode, thresh)
	}
	return total
}

func RecurseMinimumSize_7(dir *Inode, thresh int) int {
	total := math.MaxInt
	if dir.size >= thresh && dir.size < total {
		total = dir.size
	}

	for _, inode := range dir.children {
		if inode.mode == M_FILE {
			continue
		}
		size := RecurseMinimumSize_7(inode, thresh)
		if size >= thresh && size < total {
			total = size
		}
	}
	return total
}

func Solution7(r io.Reader) {
	root := Inode{mode: M_DIR, children: make(map[string]*Inode, 0)}
	cwd := &root
	root.parent = &root

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f := strings.Fields(scanner.Text())
		switch f[0] {
		case "$":
			if f[1] == "cd" {
				if f[2] == "/" {
					cwd = &root
				} else if f[2] == ".." {
					cwd = cwd.parent
				} else {
					cwd = cwd.children[f[2]]
				}
			}
		case "dir":
			cwd.children[f[1]] = &Inode{mode: M_DIR, parent: cwd, children: make(map[string]*Inode, 0)}
		default:
			size, _ := strconv.Atoi(f[0])
			cwd.children[f[1]] = &Inode{mode: M_FILE, parent: cwd, size: size}
		}
	}

	thresh := 100000
	fmt.Println(RecurseTotalSize_7(&root, thresh))

	fs_total := 70000000
	fs_required := 30000000
	fs_needs := fs_required - (fs_total - DirSizeRecurse_7(&root))
	fmt.Println(RecurseMinimumSize_7(&root, fs_needs))
}
