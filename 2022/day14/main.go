package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/robertjli/adventofcode/util"
)

const debug = false

var day = fmt.Sprintf("2022/day%d/", 14)

func main() {
	fmt.Println("Sample Input")
	fmt.Println("------------")

	fmt.Print("Building map...")
	sampleMap := buildMap("sample.txt")
	fmt.Println("Done")

	fmt.Print("Solving...")
	sampleCaught, sampleSands := solve(sampleMap)
	fmt.Println("Done")

	fmt.Print("Part 1:\t")
	util.Assert(sampleCaught, 24)
	fmt.Print("\n")
	fmt.Print("Part 2:\t")
	util.Assert(sampleSands, 93)
	fmt.Print("\n\n")

	fmt.Println("Graded Input")
	fmt.Println("------------")

	fmt.Print("Building map...")
	m := buildMap("input.txt")
	fmt.Println("Done")

	fmt.Print("Solving...")
	answer(m)
	fmt.Println("Done")
}

type unit int

const empty unit = 0
const rock unit = 1
const sand unit = 2

type point struct {
	x, y int
}

type rockMap struct {
	m    map[point]unit
	minX int
	maxX int
	maxY int
}

func answer(m *rockMap) {
	stop := util.StartTiming()
	caught, sands := solve(m)
	stop()

	fmt.Printf("Answer for part 1:\t%d\n", caught)
	fmt.Printf("Answer for part 2:\t%d\n", sands)
}

func buildMap(file string) *rockMap {
	m := &rockMap{
		m:    make(map[point]unit),
		minX: math.MaxInt,
		maxX: math.MinInt,
		maxY: math.MinInt,
	}
	util.ProcessLines(day+file, func(line string) {
		parts := strings.Split(line, " -> ")
		var curr point
		for i, part := range parts {
			tokens := strings.Split(part, ",")
			p := point{util.ParseInt(tokens[0]), util.ParseInt(tokens[1])}
			m.m[p] = rock
			if p.x < m.minX {
				m.minX = p.x
			}
			if p.x > m.maxX {
				m.maxX = p.x
			}
			if p.y > m.maxY {
				m.maxY = p.y
			}

			if i != 0 {
				if curr.x == p.x {
					if curr.y > p.y {
						for i := curr.y - 1; i > p.y; i-- {
							m.m[point{curr.x, i}] = rock
						}
					} else if curr.y < p.y {
						for i := curr.y + 1; i < p.y; i++ {
							m.m[point{curr.x, i}] = rock
						}
					}
				} else if curr.y == p.y {
					if curr.x > p.x {
						for i := curr.x - 1; i > p.x; i-- {
							m.m[point{i, curr.y}] = rock
						}
					} else if curr.x < p.x {
						for i := curr.x + 1; i < p.x; i++ {
							m.m[point{i, curr.y}] = rock
						}
					}
				}
			}
			curr = p
		}
	})

	if debug {
		fmt.Println()
		m.print()
	}

	return m
}

func (m *rockMap) print() {
	for i := 0; i <= m.maxY+1; i++ {
		for j := m.minX; j <= m.maxX; j++ {
			var u unit
			var ok bool
			if u, ok = m.m[point{j, i}]; !ok {
				u = empty
			}

			switch u {
			case empty:
				if i == 0 && j == 500 {
					fmt.Print("V")
				} else {
					fmt.Print(".")
				}
			case rock:
				fmt.Print("#")
			case sand:
				fmt.Print("o")
			}
		}
		fmt.Println()
	}

	for j := m.minX; j <= m.maxX; j++ {
		fmt.Print("#")
	}
	fmt.Println()
}

func solve(m *rockMap) (caught int, sands int) {
	if debug {
		fmt.Println()
	}

	abyss := false
	for ; m.m[point{500, 0}] == empty; sands++ {
		p := dropOne(m)

		if debug {
			fmt.Printf("dropped sand %d, landed at (%d, %d)\n", sands+1, p.x, p.y)
			m.print()
		}

		if !abyss {
			if p.y < m.maxY {
				caught++
			} else {
				abyss = true
			}
		}
	}
	return caught, sands
}

func dropOne(m *rockMap) point {
	s := point{500, 0}
	next := getNext(m.m, s)
	for s != next { // still moving
		if s.y == m.maxY+1 {
			if s.x < m.minX {
				m.minX = s.x
			}
			if s.x > m.maxX {
				m.maxX = s.x
			}
			// landed on the floor
			break
		}

		s = next
		next = getNext(m.m, s)
	}

	// now at rest, record the final spot
	m.m[s] = sand
	return s
}

func getNext(m map[point]unit, s point) point {
	if u, ok := m[point{s.x, s.y + 1}]; !ok || u == empty {
		return point{s.x, s.y + 1}
	}

	if u, ok := m[point{s.x - 1, s.y + 1}]; !ok || u == empty {
		return point{s.x - 1, s.y + 1}
	}

	if u, ok := m[point{s.x + 1, s.y + 1}]; !ok || u == empty {
		return point{s.x + 1, s.y + 1}
	}

	return s
}
