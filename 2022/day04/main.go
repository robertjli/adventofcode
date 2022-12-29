package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/robertjli/adventofcode/util"
)

const debug = false

var day = fmt.Sprintf("2022/day%d/", 4)

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

	var count int
	util.ProcessLines(day+file, func(line string) {
		a1, a2, b1, b2 := parse(line)
		if (a1 >= b1 && a2 <= b2) || (b1 >= a1 && b2 <= a2) {
			if debug {
				fmt.Println("found overlap!")
			}
			count++
		}
	})

	fmt.Printf("Answer for part 1:\t%d\n", count)
}

func part2(file string) {
	stop := util.StartTiming()
	defer stop()

	var count int
	util.ProcessLines(day+file, func(line string) {
		a1, a2, b1, b2 := parse(line)
		if (a1 >= b1 && a1 <= b2) || (b1 >= a1 && b1 <= a2) {
			if debug {
				fmt.Println("found overlap!")
			}
			count++
		}
	})

	fmt.Printf("Answer for part 2:\t%d\n", count)
}

func parse(line string) (a1, a2, b1, b2 int) {
	s := strings.FieldsFunc(line, func(r rune) bool {
		return r == ',' || r == '-'
	})

	var err error
	a1, err = strconv.Atoi(s[0])
	util.FailIf(err, "failed to parse first elf start section")
	a2, err = strconv.Atoi(s[1])
	util.FailIf(err, "failed to parse first elf end section")
	b1, err = strconv.Atoi(s[2])
	util.FailIf(err, "failed to parse second elf start section")
	b2, err = strconv.Atoi(s[3])
	util.FailIf(err, "failed to parse second elf end section")
	if debug {
		fmt.Printf("first elf: %d-%d, second elf: %d-%d\n", a1, a2, b1, b2)
	}

	return
}
