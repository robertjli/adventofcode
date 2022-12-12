package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/gammazero/deque"

	"github.com/robertjli/adventofcode2022/util"
)

const debug = false

var day = fmt.Sprintf("day%d/", 12)

func main() {
	fmt.Println("Sample Input")
	fmt.Println("------------")

	fmt.Println("Building map...")
	sampleMap, sampleStart, sampleLows := buildMap("sample.txt")

	fmt.Print("Part 1:\t")
	util.Assert(solve(sampleMap, sampleStart), 31)
	fmt.Print("\n")
	fmt.Print("Part 2:\t")
	util.Assert(solve(sampleMap, sampleLows...), 29)
	fmt.Print("\n\n")

	fmt.Println("Graded Input")
	fmt.Println("------------")

	fmt.Println("Building map...")
	m, start, lows := buildMap("input.txt")

	part1(m, start)
	part2(m, lows)
}

type coord struct {
	x, y  int
	steps int
}

func (c *coord) Bare() bareCoord {
	return bareCoord{c.x, c.y}
}

type bareCoord struct {
	x, y int
}

type heightmap struct {
	grid [][]rune
}

func (m *heightmap) Print() {
	for _, row := range m.grid {
		for _, c := range row {
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}

func (m *heightmap) CharAt(c coord) rune {
	return m.grid[c.x][c.y]
}

func (m *heightmap) ValAt(c coord) rune {
	return m.ValAtXY(c.x, c.y)
}

func (m *heightmap) ValAtXY(x, y int) rune {
	val := m.grid[x][y]
	if val == 'S' {
		return 'a'
	} else if val == 'E' {
		return 'z'
	}
	return val
}

func part1(m *heightmap, start coord) {
	stop := util.StartTiming()
	defer stop()

	fmt.Printf("Answer for part 1:\t%d\n", solve(m, start))
}

func part2(m *heightmap, lows []coord) {
	stop := util.StartTiming()
	defer stop()

	fmt.Printf("Answer for part 2:\t%d\n", solve(m, lows...))
}

func buildMap(file string) (*heightmap, coord, []coord) {
	scanner, closeFunc := util.NewScanner(day + file)
	defer closeFunc()

	m := &heightmap{}
	m.grid = make([][]rune, 0, 41)
	var start coord
	a := make([]coord, 0)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		m.grid = append(m.grid, []rune(line))

		for col, c := range line {
			if c == 'S' {
				start = coord{row, col, 0}
				a = append(a, coord{row, col, 0})
			} else if c == 'a' {
				a = append(a, coord{row, col, 0})
			}
		}

		if col := strings.IndexRune(line, 'S'); col != -1 {
			start = coord{row, col, 0}
		}
		if col := strings.IndexRune(line, 'S'); col != -1 {
			start = coord{row, col, 0}
		}
		row++
	}

	if debug {
		m.Print()
		fmt.Printf("Start: (%d, %d)/n", start.x, start.y)
		fmt.Printf("All a: ")
		for i, c := range a {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("(%d, %d),", c.x, c.y)
		}
		fmt.Println()
	}

	return m, start, a
}

func solve(m *heightmap, starts ...coord) int {
	min := math.MaxInt
	for _, start := range starts {
		length := solveMap(m, start)
		if length < min && length > 0 {
			min = length
		}
	}
	return min
}

func solveMap(m *heightmap, start coord) int {
	queue := deque.New[coord]()
	queue.PushBack(start)
	seen := make(map[bareCoord]struct{})

	if debug {
		fmt.Println()
	}

	for queue.Len() > 0 {
		elem := queue.PopFront()
		if _, ok := seen[elem.Bare()]; ok {
			continue
		}

		seen[elem.Bare()] = struct{}{}
		for _, child := range children(elem, m) {
			if m.CharAt(child) == 'E' {
				return child.steps
			}
			if _, ok := seen[child.Bare()]; !ok {
				queue.PushBack(child)
			}
		}
	}

	return -1
}

func children(c coord, m *heightmap) []coord {
	grid := m.grid
	val := m.ValAt(c)

	children := make([]coord, 0, 4)

	if c.y > 0 && m.ValAtXY(c.x, c.y-1) <= val+1 {
		children = append(children, coord{c.x, c.y - 1, c.steps + 1})
	}

	if c.y < len(grid[c.x])-1 && m.ValAtXY(c.x, c.y+1) <= val+1 {
		children = append(children, coord{c.x, c.y + 1, c.steps + 1})
	}

	if c.x > 0 && m.ValAtXY(c.x-1, c.y) <= val+1 {
		children = append(children, coord{c.x - 1, c.y, c.steps + 1})
	}

	if c.x < len(grid)-1 && m.ValAtXY(c.x+1, c.y) <= val+1 {
		children = append(children, coord{c.x + 1, c.y, c.steps + 1})
	}

	if debug {
		fmt.Printf("Children for (%d, %d, s%d): ", c.x, c.y, c.steps)
		for i, child := range children {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("(%d, %d, s%d)", child.x, child.y, child.steps)
		}
		fmt.Println()
	}

	return children
}
