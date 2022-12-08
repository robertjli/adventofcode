package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/robertjli/adventofcode2022/util"
)

const debug = false

var day = fmt.Sprintf("day%d/", 7)

func main() {
	fmt.Println("Sample Input")
	fmt.Println("------------")

	fmt.Println("Building tree")
	sampleTree := build(day + "sample.txt")
	if debug {
		printTree(sampleTree)
	}

	fmt.Print("Part 1:\t")
	util.Assert(solve1(sampleTree) == 95437)
	fmt.Print("\n")
	fmt.Print("Part 2:\t")
	util.Assert(solve2(sampleTree) == 24933642)
	fmt.Print("\n\n")

	fmt.Println("Graded Input")
	fmt.Println("------------")

	fmt.Println("Building tree")
	tree := build(day + "input.txt")
	if debug {
		printTree(tree)
	}

	part1(tree)
	part2(tree)
}

func part1(tree node) {
	stop := util.StartTiming()
	defer stop()

	fmt.Printf("Answer for part 1:\t%d\n", solve1(tree))
}

func part2(tree node) {
	stop := util.StartTiming()
	defer stop()

	fmt.Printf("Answer for part 2:\t%d\n", solve2(tree))
}

func solve1(tree node) int {
	var sum int
	reduce(tree, func(node node) {
		switch n := node.(type) {
		case *dir:
			if n.size < 100000 {
				sum += n.size
			}
		}
	})

	return sum
}

func solve2(tree node) int {
	diff := 70000000 - tree.getSize()
	needed := 30000000 - diff
	min := math.MaxInt
	reduce(tree, func(node node) {
		switch n := node.(type) {
		case *dir:
			if n.size > needed && n.size < min {
				min = n.size
			}
		}
	})

	return min
}

func build(file string) node {
	stop := util.StartTiming()
	defer stop()

	root := &dir{name: "/"}
	var curr node = root

	lines := util.ReadFile(file)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if debug {
			fmt.Printf("cmd: %s, ", line)
		}

		if strings.HasPrefix(line, "$ cd ") {
			target := line[len("$ cd "):]
			switch target {
			case "/":
				if debug {
					fmt.Println("setting current to root")
				}
				curr = root
			case "..":
				if debug {
					fmt.Printf("setting current to %s\n", curr.getParent().name)
				}
				curr = curr.getParent()
			default:
				if debug {
					fmt.Printf("setting current to %s\n", target)
				}
				d := curr.(*dir)
				for _, c := range d.children {
					if c.getName() == target {
						curr = c
						break
					}
				}
			}
		} else if line == "$ ls" {
			if debug {
				fmt.Printf("listing contents of %s\n", curr.getName())
			}
			for i++; i < len(lines); i++ {
				line = lines[i]
				if strings.HasPrefix(line, "$") {
					i--
					break
				}
				d := curr.(*dir)
				d.children = append(d.children, newNode(line, d))
			}
		}
	}

	size(root)

	return root
}

func newNode(line string, parent *dir) node {
	if strings.HasPrefix(line, "dir ") {
		return &dir{
			name:   line[len("dir "):],
			parent: parent,
		}
	}

	tokens := strings.SplitN(line, " ", 2)
	return &file{
		name:   tokens[1],
		size:   util.ParseInt(tokens[0]),
		parent: parent,
	}
}

func size(root node) {
	if root.getSize() == 0 {
		d := root.(*dir) // files always have size
		for _, c := range d.children {
			size(c)
			d.size += c.getSize()
		}
	}
}

func reduce(root node, reducer func(node)) {
	reducer(root)
	switch n := root.(type) {
	case *dir:
		for _, c := range n.children {
			reduce(c, reducer)
		}
	}
}
