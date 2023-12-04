#!/bin/zsh

year=$1
day=$2

dir_path="$year/day$(printf %02d $day)"

mkdir -p $dir_path

touch $dir_path/sample.txt

$(dirname $0)/fetch_input.sh $year $day

cat > $dir_path/main.go <<EOF
package main

import (
	"fmt"

	"github.com/robertjli/adventofcode/util"
)

/*
<Insert problem statement>
*/

var day = util.DayPath($year, $day)

func main() {
	fmt.Println("Sample Input")
	fmt.Println("------------")
	fmt.Print("Part 1:\t")
	util.Assert(solve1("sample.txt"), 0)
	// fmt.Print("\n")
	// fmt.Print("Part 2:\t")
	// util.Assert(solve2("sample.txt"), 0)
	fmt.Print("\n\n")

	fmt.Println("Graded Input")
	fmt.Println("------------")
	// part1()
	// part2()
}

func part1() {
	defer util.StartTiming()()
	fmt.Printf("Part 1:\t%d\n", solve1("input.txt"))
}

func solve1(file string) int {
	return 0
}

func part2() {
	defer util.StartTiming()()
	fmt.Printf("Part 2:\t%d\n", solve2("input.txt"))
}

func solve2(file string) int {
	return 0
}
EOF
