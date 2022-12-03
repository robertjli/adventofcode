package main

import (
	"fmt"

	"github.com/robertjli/adventofcode2022/util"
)

const debug = false

var day = fmt.Sprintf("day%d/", 3)

func main() {
	fmt.Println("Sample Input")
	fmt.Println("------------")
	part1("sample.txt")
	part2("sample.txt")
	fmt.Println()

	fmt.Println("Graded Input")
	fmt.Println("------------")
	part1("input.txt")
	part2("input.txt")
}

func part1(file string) {
	stop := util.StartTiming()
	defer stop()

	var sum int
	util.ProcessLines(day+file, func(line string) {
		first := make(map[rune]struct{}, len(line)/2)
		for _, c := range line[:len(line)/2] {
			if debug {
				fmt.Printf("adding %c to set\n", c)
			}
			first[c] = struct{}{}
		}

		for _, c := range line[len(line)/2:] {
			if debug {
				fmt.Printf("checking if %c is in the set\n", c)
			}
			if _, ok := first[c]; ok {
				p := priorityFor(c)
				sum += p
				if debug {
					fmt.Printf("found %c in the set, adding %d, running sum %d\n\n", c, p, sum)
				}
				break
			}
		}
	})

	fmt.Printf("Answer for part 1:\t%d\n", sum)
}

func part2(file string) {
	stop := util.StartTiming()
	defer stop()

	var sum, counter int
	set := make(map[rune]int)
	util.ProcessLines(day+file, func(line string) {
		if counter == 0 {
			for _, c := range line {
				if debug {
					fmt.Printf("adding %c to set\n", c)
				}
				set[c] = 1
			}
		} else if counter == 1 {
			for _, c := range line {
				if _, ok := set[c]; ok {
					if debug {
						fmt.Printf("incrementing %c in set\n", c)
					}
					set[c] = 2
				}
			}
		} else {
			for _, c := range line {
				if v, ok := set[c]; ok && v == 2 {
					p := priorityFor(c)
					sum += p
					if debug {
						fmt.Printf("found %c in the set, adding %d, running sum %d\n\n", c, p, sum)
					}
					set = make(map[rune]int)
					break
				}
			}
		}
		counter = (counter + 1) % 3
	})

	fmt.Printf("Answer for part 2:\t%d\n", sum)
}
func priorityFor(c rune) int {
	if c >= 97 {
		return int(c - 96)
	}
	return int(c - 38)
}
