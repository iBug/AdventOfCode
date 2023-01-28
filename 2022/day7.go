package year

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

type Inode7 struct {
	mode     byte
	size     int
	parent   *Inode7
	children map[string]*Inode7
}

const (
	M_FILE = iota
	M_DIR
)

func DirSizeRecurse7(dir *Inode7) int {
	if dir.size > 0 {
		return dir.size
	}

	total := 0
	for _, inode := range dir.children {
		if inode.mode == M_FILE {
			total += inode.size
		} else {
			total += DirSizeRecurse7(inode)
		}
	}
	dir.size = total
	return total
}

func RecurseTotalSize7(dir *Inode7, thresh int) int {
	total := 0
	for _, inode := range dir.children {
		if inode.mode == M_FILE {
			continue
		}
		size := DirSizeRecurse7(inode)
		if size <= thresh {
			total += size
		}
		total += RecurseTotalSize7(inode, thresh)
	}
	return total
}

func RecurseMinimumSize7(dir *Inode7, thresh int) int {
	total := math.MaxInt
	if dir.size >= thresh && dir.size < total {
		total = dir.size
	}

	for _, inode := range dir.children {
		if inode.mode == M_FILE {
			continue
		}
		size := RecurseMinimumSize7(inode, thresh)
		if size >= thresh && size < total {
			total = size
		}
	}
	return total
}

func Solution7(r io.Reader, w io.Writer) {
	root := Inode7{mode: M_DIR, children: make(map[string]*Inode7, 0)}
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
			cwd.children[f[1]] = &Inode7{mode: M_DIR, parent: cwd, children: make(map[string]*Inode7, 0)}
		default:
			size, _ := strconv.Atoi(f[0])
			cwd.children[f[1]] = &Inode7{mode: M_FILE, parent: cwd, size: size}
		}
	}

	thresh := 100000
	fmt.Fprintln(w, RecurseTotalSize7(&root, thresh))

	fs_total := 70000000
	fs_required := 30000000
	fs_needs := fs_required - (fs_total - DirSizeRecurse7(&root))
	fmt.Fprintln(w, RecurseMinimumSize7(&root, fs_needs))
}
