package main

import (
	"fmt"
	"strconv"

	"github.com/robertjli/adventofcode2022/util"
)

const debug = false

var day = fmt.Sprintf("day%d/", 1)

func main() {
	fmt.Println("Part 1")
	fmt.Println("------")
	part1("sample.txt")
	part1("input.txt")
	fmt.Println()

	fmt.Println("Part 2")
	fmt.Println("------")
	part2("sample.txt")
	part2("input.txt")
}

func part1(file string) {
	stop := util.StartTiming()
	defer stop()

	var cur, max int
	util.ProcessLines(day+file, func(line string) {
		if len(line) > 0 {
			val, err := strconv.Atoi(line)
			util.FailIf(err, "could not parse", line)
			cur += val
			if debug {
				fmt.Printf("parsed %d, running total %d, max %d\n", val, cur, max)
			}
		} else {
			if cur > max {
				max = cur
			}
			if debug {
				fmt.Printf("end of elf, total %d, max %d\n\n", cur, max)
			}
			cur = 0
		}
	})

	// last elf
	if cur > max {
		max = cur
	}

	fmt.Printf("Answer for %s:\t%d\n", file, max)
}

func part2(file string) {
	stop := util.StartTiming()
	defer stop()

	var cur int
	max := make([]int, 3)
	util.ProcessLines(day+file, func(line string) {
		if len(line) > 0 {
			val, err := strconv.Atoi(line)
			util.FailIf(err, "could not parse", line)
			cur += val
			if debug {
				fmt.Printf("parsed %d, running total %d, max %v\n", val, cur, max)
			}
		} else {
			insert(max, cur)
			if debug {
				fmt.Printf("end of elf, total %d, max %v\n\n", cur, max)
			}
			cur = 0
		}
	})

	// last elf
	insert(max, cur)

	var sum int
	for _, m := range max {
		sum += m
	}
	fmt.Printf("Answer for %s:\t%d\n", file, sum)
}

// Assumes s is length 3
func insert(s []int, x int) {
	if x > s[0] {
		s[0] = x
		if s[0] > s[1] {
			s[0] = s[1]
			s[1] = x
			if s[1] > s[2] {
				s[1] = s[2]
				s[2] = x
			}
		}
	}
}
