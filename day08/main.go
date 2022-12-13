package main

import (
	"fmt"

	"github.com/robertjli/adventofcode2022/util"
)

const debug = false

var day = fmt.Sprintf("day%d/", 8)

func main() {
	fmt.Println("Sample Input")
	fmt.Println("------------")

	if debug {
		fmt.Println("Growing trees")
	}
	sampleGrid, sampleVis := build(day + "sample.txt")

	fmt.Print("Part 1:\t")
	util.Assert(solve1(sampleVis), 21)
	fmt.Print("\n")
	fmt.Print("Part 2:\t")
	util.Assert(solve2(sampleGrid), 8)
	fmt.Print("\n\n")

	fmt.Println("Graded Input")
	fmt.Println("------------")

	if debug {
		fmt.Println("Growing trees")
	}
	grid, vis := build(day + "input.txt")

	part1(vis)
	part2(grid)
}

func part1(vis [][]bool) {
	stop := util.StartTiming()
	defer stop()

	fmt.Printf("Answer for part 1:\t%d\n", solve1(vis))
}

func solve1(vis [][]bool) int {
	var sum int
	for _, row := range vis {
		for _, cell := range row {
			if cell {
				sum += 1
			}
		}
	}

	return sum
}

func part2(grid [][]int) {
	stop := util.StartTiming()
	defer stop()

	fmt.Printf("Answer for part 2:\t%d\n", solve2(grid))
}

func solve2(grid [][]int) int {
	var max int
	for i, row := range grid {
		for j := range row {
			score := calculateVisFromTree(grid, i, j)
			if debug {
				fmt.Printf("vis score from (%d,%d): %d\n", i, j, score)
			}
			if score > max {
				max = score
			}
		}
	}

	return max
}

func build(file string) ([][]int, [][]bool) {
	stop := util.StartTiming()
	defer stop()

	grid := make([][]int, 0, 99)
	util.ProcessLines(file, func(line string) {
		row := make([]int, 0, len(line))
		for _, c := range line {
			row = append(row, util.ParseInt(string(c)))
		}
		grid = append(grid, row)
	})

	return grid, vis(grid)
}

func printGrid(grid [][]int) {
	for _, row := range grid {
		for i, cell := range row {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Printf("%d", cell)
		}
		fmt.Println()
	}
}

func vis(grid [][]int) [][]bool {
	vis := make([][]bool, len(grid))
	for i := range vis {
		vis[i] = make([]bool, len(grid[i]))
	}
	for i := 0; i < len(grid[0]); i++ {
		vis[0][i] = true
		vis[len(grid)-1][i] = true
	}
	for i := 1; i < len(grid)-1; i++ {
		vis[i][0] = true
		vis[i][len(grid[i])-1] = true
	}

	for i := 1; i < len(grid)-1; i++ {
		// left-to-right
		max := grid[i][0]
		for j := 1; j < len(grid[i])-1; j++ {
			if grid[i][j] > max {
				vis[i][j] = true
				max = grid[i][j]
			}
		}

		// right-to-left
		max = grid[i][len(grid[i])-1]
		for j := len(grid[i]) - 2; j > 0; j-- {
			if grid[i][j] > max {
				vis[i][j] = true
				max = grid[i][j]
			}
		}
	}

	for j := 1; j < len(grid[0])-1; j++ {
		// top-to-bottom
		max := grid[0][j]
		for i := 1; i < len(grid)-1; i++ {
			if grid[i][j] > max {
				vis[i][j] = true
				max = grid[i][j]
			}
		}

		// bottom-to-top
		max = grid[len(grid)-1][j]
		for i := len(grid) - 2; i > 0; i-- {
			if grid[i][j] > max {
				vis[i][j] = true
				max = grid[i][j]
			}
		}
	}

	if debug {
		fmt.Println("\nGrid with vis")
		printGridWithVis(grid, vis)
	}

	return vis
}

func printGridWithVis(grid [][]int, vis [][]bool) {
	for i, row := range grid {
		for j, cell := range row {
			if j > 0 {
				fmt.Print(" ")
			}

			v := 'F'
			if vis[i][j] {
				v = 'T'
			}

			fmt.Printf("%d%c", cell, v)
		}
		fmt.Println()
	}
}

func calculateVisFromTree(grid [][]int, x, y int) int {
	max := grid[x][y]
	score := 1
	curr := 0

	for i := x - 1; i >= 0; i-- {
		curr++
		if grid[i][y] >= max {
			break
		}
	}
	score *= curr
	curr = 0

	for i := x + 1; i < len(grid); i++ {
		curr++
		if grid[i][y] >= max {
			break
		}
	}
	score *= curr
	curr = 0

	for j := y - 1; j >= 0; j-- {
		curr++
		if grid[x][j] >= max {
			break
		}
	}
	score *= curr
	curr = 0

	for j := y + 1; j < len(grid[x]); j++ {
		curr++
		if grid[x][j] >= max {
			break
		}
	}
	score *= curr
	return score
}
