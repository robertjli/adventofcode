package main

import (
	"fmt"

	"github.com/robertjli/adventofcode2022/util"
)

const debug = false

var day = fmt.Sprintf("day%d/", 6)

func main() {
	fmt.Println("Sample Input")
	fmt.Println("------------")
	fmt.Print("part 1:\t")
	util.Assert(solve(4, "mjqjpqmgbljsphdztnvjfqwrcgsmlb") == 7)
	util.Assert(solve(4, "bvwbjplbgvbhsrlpgdmjqwftvncz") == 5)
	util.Assert(solve(4, "nppdvjthqldpwncqszvftbrmjlhg") == 6)
	util.Assert(solve(4, "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg") == 10)
	util.Assert(solve(4, "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw") == 11)
	fmt.Print("\n")
	fmt.Print("part 2:\t")
	util.Assert(solve(14, "mjqjpqmgbljsphdztnvjfqwrcgsmlb") == 19)
	util.Assert(solve(14, "bvwbjplbgvbhsrlpgdmjqwftvncz") == 23)
	util.Assert(solve(14, "nppdvjthqldpwncqszvftbrmjlhg") == 23)
	util.Assert(solve(14, "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg") == 29)
	util.Assert(solve(14, "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw") == 26)
	fmt.Print("\n\n")

	fmt.Println("Graded Input")
	fmt.Println("------------")
	part1("input.txt")
	part2("input.txt")
}

func part1(file string) {
	exec(file, 1, 4)
}

func part2(file string) {
	exec(file, 2, 14)
}

func exec(file string, part int, len int) {
	stop := util.StartTiming()
	defer stop()

	line := util.ReadFile(day + file)[0]
	fmt.Printf("Answer for part %d:\t%d\n", part, solve(len, line))
}

func solve(size int, line string) int {
	idx := 0
	buf := make([]uint8, size)
	for i := 0; i < size; i++ {
		buf[i] = line[i]
	}

	counts := make([]int, 26)
	for _, c := range buf {
		counts[index(c)]++
	}

	if debug {
		fmt.Println("Initial state:")
		print(buf, counts)
	}

	for i := 4; i < len(line); i++ {
		pop := buf[idx]
		counts[index(pop)]--

		c := line[i]
		buf[idx] = c
		counts[index(c)]++
		idx = (idx + 1) % size
		if debug {
			print(buf, counts)
		}

		if satisfied(counts) {
			if debug {
				fmt.Printf("All %d chars unique, distance %d\n", size, i+1)
			}
			return i + 1
		}
	}

	return -1
}

func index(c uint8) int {
	return int(rune(c) - 'a')
}

func satisfied(counts []int) bool {
	for _, c := range counts {
		if c > 1 {
			return false
		}
	}
	return true
}

func print(buf []uint8, counts []int) {
	fmt.Print("Buffer:\t")
	for _, c := range buf {
		fmt.Print(string(rune(c)))
	}
	fmt.Printf("\nCounts:\t%v\n", counts)
	fmt.Println("       \t a b c d e f g h i j k l m n o p q r s t u v w x y z")
}
