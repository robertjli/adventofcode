package main

import (
	"fmt"

	"github.com/robertjli/adventofcode/util"
)

const debug = false

var day = fmt.Sprintf("2022/day%d/", 2)

var scoresPart1 = map[string]int{
	"A X": 4, // rock     (1) + draw (3)
	"A Y": 8, // paper    (2) + win  (6)
	"A Z": 3, // scissors (3) + lose (0)
	"B X": 1, // rock     (1) + lose (0)
	"B Y": 5, // paper    (2) + draw (3)
	"B Z": 9, // scissors (3) + win  (6)
	"C X": 7, // rock     (1) + win  (6)
	"C Y": 2, // paper    (2) + lose (0)
	"C Z": 6, // scissors (3) + draw (3)
}

var scoresPart2 = map[string]int{
	"A X": 3, // lose (0) + scissors (3)
	"A Y": 4, // draw (3) + rock     (1)
	"A Z": 8, // win  (6) + paper    (2)
	"B X": 1, // lose (0) + rock     (1)
	"B Y": 5, // draw (3) + paper    (2)
	"B Z": 9, // win  (6) + scissors (3)
	"C X": 2, // lose (0) + paper    (2)
	"C Y": 6, // draw (3) + scissors (3)
	"C Z": 7, // win  (6) + rock     (1)
}

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

func solve(file string, part int) {
	stop := util.StartTiming()
	defer stop()

	var sum int
	scores := scoresPart1
	if part == 2 {
		scores = scoresPart2
	}
	util.ProcessLines(day+file, func(line string) {
		sum += scores[line]
		if debug {
			fmt.Printf("round: %s, score: %d, running sum: %d\n", line, scores[line], sum)
		}
	})

	fmt.Printf("Answer for part %d:\t%d\n", part, sum)
}

func part1(file string) {
	solve(file, 1)
}

func part2(file string) {
	solve(file, 2)
}
