package main

import (
	"fmt"
	"strings"

	"github.com/robertjli/adventofcode/util"
)

const debug = false

var day = fmt.Sprintf("2022/day%d/", 10)

func main() {
	fmt.Println("Sample Input")
	fmt.Println("------------")
	fmt.Print("Part 1:\t")
	util.Assert(solve1("sample.txt"), 13140)
	fmt.Print("\n")
	fmt.Print("Part 2:\t")
	util.AssertSlice(solve2("sample.txt"), []string{
		"##..##..##..##..##..##..##..##..##..##..",
		"###...###...###...###...###...###...###.",
		"####....####....####....####....####....",
		"#####.....#####.....#####.....#####.....",
		"######......######......######......####",
		"#######.......#######.......#######.....",
	})
	fmt.Print("\n\n")

	fmt.Println("Graded Input")
	fmt.Println("------------")
	part1("input.txt")
	part2("input.txt")
}

func part1(file string) {
	stop := util.StartTiming()
	defer stop()

	fmt.Printf("Answer for part 1:\t%d\n", solve1(file))
}

func solve1(file string) int {
	cycle := 1
	rec := 20
	val := 1
	sum := 0

	util.ProcessLines(day+file, func(line string) {
		if debug {
			fmt.Printf("cycle: %d, val: %d, read command %s\n", cycle, val, line)
		}
		if cycle == rec {
			sum += val * rec
			rec += 40
			if debug {
				fmt.Printf("this was a significant cycle, added %d, running sum %d\n", val*rec, sum)
			}
		}
		cycle++

		if strings.HasPrefix(line, "addx") {
			if debug {
				fmt.Printf("cycle: %d, val: %d, processing %s\n", cycle, val, line)
			}

			if cycle == rec {
				sum += val * rec
				rec += 40
				if debug {
					fmt.Printf("this was a significant cycle, added %d, running sum %d\n", val*rec, sum)
				}
			}
			cycle++
			val += util.ParseInt(strings.Fields(line)[1])
		}
	})

	return sum
}

func part2(file string) {
	stop := util.StartTiming()
	defer stop()

	fmt.Println("Answer for part 2:")
	render(solve2(file))
}

func solve2(file string) []string {
	height := 6
	width := 40
	crt := make([]string, 0, height)
	val := 1
	var curr string
	var builder strings.Builder

	scanner, closeFunc := util.NewScanner(day + file)
	defer closeFunc()

	for row := 0; row < height; row++ {
		builder.Grow(width)

		for col := 0; col < width; col++ {
			if debug {
				fmt.Printf("row: %d, cycle: %d, val: %d", row, col+1, val)
			}

			if col == val || col == val+1 || col == val-1 {
				builder.WriteRune('#')
				if debug {
					fmt.Printf(", drew in col %d", col)
				}
			} else {
				builder.WriteRune('.')
			}

			if debug {
				fmt.Println()
			}

			if len(curr) > 0 {
				val += util.ParseInt(strings.Fields(curr)[1])
				curr = ""
			} else if scanner.Scan() {
				line := scanner.Text()
				if strings.HasPrefix(line, "addx") {
					curr = line
				}
			}
		}

		crt = append(crt, builder.String())
		builder.Reset()
	}

	return crt
}

func render(crt []string) {
	for _, line := range crt {
		fmt.Println(line)
	}
}
