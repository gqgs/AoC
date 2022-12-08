package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Dir struct {
	children map[string]*Dir
	parent   *Dir
	size     int
}

func NewDir(parent *Dir) *Dir {
	return &Dir{
		children: make(map[string]*Dir),
		parent:   parent,
	}
}

func (d Dir) Sizes() (int, []int) {
	size := d.size
	var sizes []int
	for _, child := range d.children {
		csize, csizes := child.Sizes()
		size += csize
		sizes = append(sizes, csizes...)
	}
	sizes = append(sizes, size)
	return size, sizes
}

func (d *Dir) ProcessCmd(cmd string) *Dir {
	cmd = strings.TrimSpace(cmd)
	switch {
	case strings.HasPrefix(cmd, "$ ls"):
		// pass
	case strings.HasPrefix(cmd, "$ cd "):
		dstDir := strings.TrimPrefix(cmd, "$ cd ")
		if dstDir == ".." {
			return d.parent
		}
		if dstDir == "/" {
			return d
		}
		for name, c := range d.children {
			if name == dstDir {
				return c
			}
		}
		panic(fmt.Errorf("dir not found: %s", dstDir))
	case strings.HasPrefix(cmd, "dir "):
		childDir := strings.TrimPrefix(cmd, "dir ")
		d.children[childDir] = NewDir(d)
	default:
		var fileSize int
		var fileName int
		fmt.Sscanf(cmd, "%d %s", &fileSize, &fileName)
		d.size += fileSize
	}
	return d
}

func silver(cmds []string) int {
	root := NewDir(nil)
	cwd := root
	for _, cmd := range cmds {
		cwd = cwd.ProcessCmd(cmd)
	}

	var total int
	_, sizes := root.Sizes()
	for _, size := range sizes {
		if size < 1e5 {
			total += size
		}
	}

	return total
}

func gold(cmds []string) int {
	root := NewDir(nil)
	cwd := root
	for _, cmd := range cmds {
		cwd = cwd.ProcessCmd(cmd)
	}

	rootSize, sizes := root.Sizes()
	availableSpace := 70_000_000 - rootSize
	sort.Ints(sizes)
	for i, size := range sizes {
		if availableSpace+size > 30_000_000 {
			return sizes[i]
		}
	}

	return -1
}

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	var cmds []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		next := scanner.Text()
		cmds = append(cmds, next)
	}

	println("silver:", silver(cmds))
	println("gold:", gold(cmds))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
