package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gammazero/deque"

	"github.com/robertjli/adventofcode/util"
)

const debug = false

var day = fmt.Sprintf("2022/day%d/", 16)

func main() {
	fmt.Println("Sample Input")
	fmt.Println("------------")
	fmt.Print("Building graph...")
	sampleGraph := buildGraph("sample.txt")
	fmt.Println("Done")
	fmt.Print("Part 1:\t")
	util.Assert(solve1(sampleGraph), 1651)
	fmt.Print("\n")
	fmt.Print("Part 2:\t")
	util.Assert(solve2(sampleGraph), 1707)
	fmt.Print("\n\n")

	fmt.Println("Graded Input")
	fmt.Println("------------")
	fmt.Print("Building graph...")
	graph := buildGraph("input.txt")
	fmt.Println("Done")
	part1(graph)
	part2(graph)
}

func part1(graph map[string]*Node) {
	stop := util.StartTiming()
	defer stop()

	fmt.Printf("Answer for part 1:\t%d\n", solve1(graph))
}

func part2(graph map[string]*Node) {
	stop := util.StartTiming()
	defer stop()

	fmt.Printf("Answer for part 2:\t%d\n", solve2(graph))
}

func solve1(graph map[string]*Node) int {
	return findMax(graph, 30)
}

func solve2(graph map[string]*Node) int {
	return findMaxWithHelper(graph, 26)
}

var re = regexp.MustCompile(
	`Valve ([A-Z]+) has flow rate=(\d+); tunnels? leads? to valves? ([A-Z, ]+)`,
)

type Node struct {
	name   string
	rate   int
	paths  []string
	wpaths map[string]int // path to valve with nonzero flow
}

func (n *Node) String() string {
	return fmt.Sprintf("name: %s, flow rate: %2d, paths: %v", n.name, n.rate, n.paths)
}

func buildGraph(file string) map[string]*Node {
	if debug {
		fmt.Println("\nBuilding graph...")
	}

	graph := make(map[string]*Node, 61)

	util.ProcessLines(day+file, func(line string) {
		matches := re.FindStringSubmatch(line)
		name := matches[1]
		graph[name] = &Node{
			name:  name,
			rate:  util.ParseInt(matches[2]),
			paths: strings.Split(matches[3], ", "),
		}

		if debug {
			fmt.Printf("Parsed node %v\n", graph[name])
		}
	})

	graph = buildWeightedGraph(graph)
	if debug {
		fmt.Println("\nWeighted graph:")
		for name, node := range graph {
			fmt.Printf("node: %s, paths: %v\n", name, node.wpaths)
		}
	}

	return graph
}

func buildWeightedGraph(graph map[string]*Node) map[string]*Node {
	for name, node := range graph {
		node = buildWeightedGraphFrom(graph, node)
		graph[name] = node
	}

	return graph
}

func buildWeightedGraphFrom(graph map[string]*Node, node *Node) *Node {
	type State struct {
		name  string
		level int
	}

	wpaths := make(map[string]int, len(graph))
	queue := deque.New[*State]()
	queue.PushBack(&State{node.name, 0})

	for queue.Len() > 0 {
		curr := queue.PopFront()
		if _, ok := wpaths[curr.name]; !ok && curr.name != node.name {
			wpaths[curr.name] = curr.level
		}

		currNode := graph[curr.name]
		for _, neighbor := range currNode.paths {
			if _, ok := wpaths[neighbor]; !ok && neighbor != node.name {
				queue.PushBack(&State{neighbor, curr.level + 1})
			}
		}
	}

	node.wpaths = make(map[string]int, len(graph))
	for name, dist := range wpaths {
		if graph[name].rate > 0 {
			node.wpaths[name] = dist
		}
	}

	return node
}

func findMax(graph map[string]*Node, time int) int {
	if debug {
		fmt.Println("\nFinding max...")
	}

	type State struct {
		name string
		flow int
		time int
		open map[string]int // valve name -> time opened
	}

	max := 0
	start := graph["AA"]
	stack := deque.New[*State]()
	stack.PushFront(&State{
		name: start.name,
		flow: 0,
		time: 0,
		open: make(map[string]int),
	})

	for stack.Len() > 0 {
		curr := stack.PopFront()
		if curr.time >= time {
			continue
		}

		// if this valve is already open, something messed up
		if _, exists := curr.open[curr.name]; exists {
			panic(fmt.Sprintf("we arrived at valve %s but it was already open: %v",
				curr.name, curr.open))
		}

		currNode := graph[curr.name]
		if currNode.rate > 0 {
			curr.time += 1
			curr.open[curr.name] = curr.time
			curr.flow += currNode.rate * (time - curr.time)

			if curr.flow > max {
				max = curr.flow
			}

			if debug && rightPath(curr.open) {
				fmt.Printf(
					"we arrived at valve %s and opened it during minute %2d, current flow is %4d, "+
						"open valves: %v\n",
					curr.name, curr.time, curr.flow, curr.open)
			}
		}

		for neighbor, dist := range currNode.wpaths {
			if curr.time+dist >= time {
				continue
			}
			if _, exists := curr.open[neighbor]; exists {
				continue
			}
			stack.PushFront(&State{
				name: neighbor,
				flow: curr.flow,
				time: curr.time + dist,
				open: copyOpen(curr.open),
			})
		}
	}

	return max
}

func copyOpen(src map[string]int) map[string]int {
	dst := make(map[string]int, len(src)+1)
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func rightPath(open map[string]int) bool {
	for k, v := range open {
		switch k {
		case "BB":
			if v != 5 {
				return false
			}
		case "CC":
			if v != 24 {
				return false
			}
		case "DD":
			if v != 2 {
				return false
			}
		case "EE":
			if v != 21 {
				return false
			}
		case "HH":
			if v != 17 {
				return false
			}
		case "JJ":
			if v != 9 {
				return false
			}
		}
	}
	return true
}

type Open struct {
	by   string
	time int
}

func findMaxWithHelper(graph map[string]*Node, time int) int {
	if debug {
		fmt.Println("\nFinding max...")
	}

	type State struct {
		me        string
		meTime    int
		ellie     string
		ellieTime int
		flow      int
		open      map[string]Open // valve name -> time opened
	}

	max := 0
	start := graph["AA"]
	stack := deque.New[*State]()
	stack.PushFront(&State{
		me:        start.name,
		meTime:    0,
		ellie:     start.name,
		ellieTime: 0,
		flow:      0,
		open:      make(map[string]Open),
	})

	for stack.Len() > 0 {
		curr := stack.PopFront()
		if curr.meTime >= time && curr.ellieTime >= time {
			continue
		}

		// open the valve that we just arrived at
		newV := 0

		if _, exists := curr.open[curr.me]; !exists {
			newV++
			meNode := graph[curr.me]
			if meNode.rate > 0 {
				curr.meTime += 1
				curr.open[curr.me] = Open{"me", curr.meTime}
				curr.flow += meNode.rate * (time - curr.meTime)

				if debug && rightPathWithHelper(curr.open) {
					fmt.Printf(
						"I arrived at valve %s and opened it during minute %2d,"+
							" current flow is %4d, open valves: %v\n",
						curr.me, curr.meTime, curr.flow, curr.open)
				}
			}
		}

		if _, exists := curr.open[curr.ellie]; !exists {
			newV++
			ellieNode := graph[curr.ellie]
			if ellieNode.rate > 0 {
				curr.ellieTime += 1
				curr.open[curr.ellie] = Open{"ellie", curr.ellieTime}
				curr.flow += ellieNode.rate * (time - curr.ellieTime)

				if debug && rightPathWithHelper(curr.open) {
					fmt.Printf(
						"Ellie arrived at valve %s and opened it during minute %2d,"+
							" current flow is %4d, open valves: %v\n",
						curr.me, curr.meTime, curr.flow, curr.open)
				}
			}
		}

		if newV == 0 {
			panic(fmt.Sprintf("no valves were opened :(, me %s, ellie %s, open %v",
				curr.me, curr.ellie, curr.open))
		}

		if curr.flow > max {
			max = curr.flow
		}

		// we move one person at a time, so whoever has more time should move now
		iMoved := false
		if curr.meTime < curr.ellieTime {
			meNode := graph[curr.me]
			for neighbor, dist := range meNode.wpaths {
				if curr.meTime+dist >= time {
					continue
				}
				if _, exists := curr.open[neighbor]; exists {
					continue
				}
				stack.PushFront(&State{
					me:        neighbor,
					meTime:    curr.meTime + dist,
					ellie:     curr.ellie,
					ellieTime: curr.ellieTime,
					flow:      curr.flow,
					open:      copyOpenWithHelper(curr.open),
				})
				iMoved = true
			}
		}

		// handle the case where I have more time but nowhere to go
		if !iMoved {
			ellieNode := graph[curr.ellie]
			for neighbor, dist := range ellieNode.wpaths {
				if curr.ellieTime+dist >= time {
					continue
				}
				if _, exists := curr.open[neighbor]; exists {
					continue
				}
				stack.PushFront(&State{
					me:        curr.me,
					meTime:    curr.meTime,
					ellie:     neighbor,
					ellieTime: curr.ellieTime + dist,
					flow:      curr.flow,
					open:      copyOpenWithHelper(curr.open),
				})
			}
		}
	}

	return max
}

func copyOpenWithHelper(src map[string]Open) map[string]Open {
	dst := make(map[string]Open, len(src)+1)
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func rightPathWithHelper(open map[string]Open) bool {
	for k, v := range open {
		var o Open
		switch k {
		case "BB":
			o = Open{"me", 7}
		case "CC":
			o = Open{"me", 9}
		case "DD":
			o = Open{"ellie", 2}
		case "EE":
			o = Open{"ellie", 11}
		case "HH":
			o = Open{"ellie", 7}
		case "JJ":
			o = Open{"me", 3}
		}
		if v != o {
			return false
		}
	}
	return true
}
