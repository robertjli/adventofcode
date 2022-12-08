package main

import "fmt"

//////////////////////
//    Tree stuff    //
//////////////////////

type node interface {
	getName() string
	getSize() int
	getParent() *dir
}

var _ node = &dir{}
var _ node = &file{}

type dir struct {
	name     string
	size     int
	parent   *dir
	children []node
}

func (d *dir) getName() string {
	return d.name
}

func (d *dir) getSize() int {
	return d.size
}

func (d *dir) getParent() *dir {
	return d.parent
}

type file struct {
	name   string
	size   int
	parent *dir
}

func (f *file) getName() string {
	return f.name
}

func (f *file) getSize() int {
	return f.size
}

func (f *file) getParent() *dir {
	return f.parent
}

func typeOf(n node) string {
	switch n.(type) {
	case *dir:
		return "dir"
	case *file:
		return "file"
	default:
		return "wot"
	}
}

func printTree(root node) {
	type entry struct {
		node   node
		indent string
	}

	queue := []entry{{root, ""}}
	for len(queue) > 0 {
		last := len(queue) - 1
		curr := queue[last]
		queue = queue[:last]
		fmt.Printf("%s%s %s (%d)\n",
			curr.indent, typeOf(curr.node), curr.node.getName(), curr.node.getSize())

		switch n := curr.node.(type) {
		case *dir:
			for i := len(n.children) - 1; i >= 0; i-- {
				queue = append(queue, entry{n.children[i], curr.indent + "    "})
			}
		}
	}
}
