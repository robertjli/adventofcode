package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/robertjli/adventofcode2022/util"
)

const debug = false

var day = fmt.Sprintf("day%d/", 9)

func main() {
	fmt.Println("Sample Input")
	fmt.Println("------------")
	fmt.Print("Part 1:\t")
	util.Assert(solve1("sample.txt"), 13)
	fmt.Print("\n")
	fmt.Print("Part 2:\t")
	util.Assert(solve2("sample.txt"), 1)
	util.Assert(solve2("sample2.txt"), 36)
	fmt.Print("\n\n")

	fmt.Println("Graded Input")
	fmt.Println("------------")
	part1("input.txt")
	part2("input.txt")
}

type pos struct {
	x, y int
}

func part1(file string) {
	stop := util.StartTiming()
	defer stop()

	fmt.Printf("Answer for part 1:\t%d\n", solve1(file))
}

func solve1(file string) int {
	visited := make(map[pos]struct{})

	head := pos{0, 0}
	tail := pos{0, 0}
	visited[tail] = struct{}{}

	if debug {
		fmt.Printf("\nStart: Head (%d,%d), Tail (%d,%d)\n", head.x, head.y, tail.x, tail.y)
	}

	util.ProcessLines(day+file, func(line string) {
		tokens := strings.Fields(line)
		dir := tokens[0]
		dist := util.ParseInt(tokens[1])
		for i := 0; i < dist; i++ {
			head = moveHead(head, dir)
			tail = moveTail(head, tail)

			if debug {
				fmt.Printf("Move %s: Head (%d,%d), Tail (%d,%d)\n", dir, head.x, head.y, tail.x, tail.y)
			}

			visited[tail] = struct{}{}
		}
	})

	return len(visited)
}

func part2(file string) {
	stop := util.StartTiming()
	defer stop()

	fmt.Printf("Answer for part 2:\t%d\n", solve2(file))
}

func solve2(file string) int {
	visited := make(map[pos]struct{})

	head := pos{0, 0}
	tails := make([]pos, 9)
	visited[tails[8]] = struct{}{}

	if debug {
		fmt.Printf("\nStart: Head (%d,%d), Tails", head.x, head.y)
		for i, t := range tails {
			if i != 0 {
				fmt.Print(",")
			}
			fmt.Printf(" (%d, %d)", t.x, t.y)
		}
		fmt.Println()
		printGrid(head, tails)
	}

	util.ProcessLines(day+file, func(line string) {
		tokens := strings.Fields(line)
		dir := tokens[0]
		dist := util.ParseInt(tokens[1])
		for i := 0; i < dist; i++ {
			head = moveHead(head, dir)
			for i := range tails {
				if i == 0 {
					tails[i] = moveTail(head, tails[i])
				} else {
					tails[i] = moveTail(tails[i-1], tails[i])
				}
			}

			if debug {
				fmt.Printf("Move: Head (%d,%d), Tails", head.x, head.y)
				for i, t := range tails {
					if i != 0 {
						fmt.Print(",")
					}
					fmt.Printf(" (%d, %d)", t.x, t.y)
				}
				fmt.Println()
				printGrid(head, tails)
			}

			visited[tails[8]] = struct{}{}
		}
	})

	return len(visited)
}

func moveHead(head pos, dir string) pos {
	switch dir {
	case "R":
		head.x += 1
	case "L":
		head.x -= 1
	case "U":
		head.y += 1
	case "D":
		head.y -= 1
	}

	return head
}

func moveTail(head, tail pos) pos {
	if head.x == tail.x {
		if head.y == tail.y+2 {
			tail.y += 1
		} else if head.y == tail.y-2 {
			tail.y -= 1
		}
	} else if head.y == tail.y {
		if head.x == tail.x+2 {
			tail.x += 1
		} else if head.x == tail.x-2 {
			tail.x -= 1
		}
	} else if math.Abs(float64(head.x-tail.x))+math.Abs(float64(head.y-tail.y)) > 2 {
		if head.x > tail.x {
			tail.x += 1
		} else {
			tail.x -= 1
		}
		if head.y > tail.y {
			tail.y += 1
		} else {
			tail.y -= 1
		}
	}
	return tail
}

func printGrid(head pos, tails []pos) {
	for i := 15; i >= -5; i-- {
		for j := -11; j <= 14; j++ {
			if j == head.x && i == head.y {
				fmt.Print("H")
			} else {
				var match bool
				for k, t := range tails {
					if j == t.x && i == t.y {
						fmt.Print(k + 1)
						match = true
						break
					}
				}

				if !match {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}
}
