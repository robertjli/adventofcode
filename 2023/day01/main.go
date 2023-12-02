package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/robertjli/adventofcode/util"
)

/*
--- Day 1: Trebuchet?! ---
Something is wrong with global snow production, and you've been selected to take a look. The Elves
have even given you a map; on it, they've used stars to mark the top fifty locations that are likely
to be having problems.

You've been doing this long enough to know that to restore snow operations, you need to check all
fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent
calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star.
Good luck!

You try to ask why they can't just use a weather machine ("not powerful enough") and where they're
even sending you ("the sky") and why your map looks mostly blank ("you sure ask a lot of questions")
and hang on did you just say the sky ("of course, where do you think snow comes from") when you
realize that the Elves are already loading you into a trebuchet ("please hold still, we need to
strap you in").

As they're making the final adjustments, they discover that their calibration document (your puzzle
input) has been amended by a very young Elf who was apparently just excited to show off her art
skills. Consequently, the Elves are having trouble reading the values on the document.

The newly-improved calibration document consists of lines of text; each line originally contained a
specific calibration value that the Elves now need to recover. On each line, the calibration value
can be found by combining the first digit and the last digit (in that order) to form a single
two-digit number.

For example:

```
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
```

In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these
together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?

--- Part Two ---
Your calculation isn't quite right. It looks like some of the digits are actually spelled out with
letters: one, two, three, four, five, six, seven, eight, and nine also count as valid "digits".

Equipped with this new information, you now need to find the real first and last digit on each line.
For example:

```
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
```

In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76. Adding these together
produces 281.

What is the sum of all of the calibration values?
*/

var day = util.DayPath(2023, 1)

func main() {
	fmt.Println("Sample Input")
	fmt.Println("------------")
	fmt.Print("Part 1:\t")
	util.Assert(solve1("sample.txt"), 142)
	fmt.Print("\n")
	fmt.Print("Part 2:\t")
	util.Assert(solve2("sample2.txt"), 281)
	fmt.Print("\n\n")

	fmt.Println("Graded Input")
	fmt.Println("------------")
	part1()
	part2()
}

func part1() {
	defer util.StartTiming()()
	fmt.Printf("Part 1:\t%d\n", solve1("input.txt"))
}

func solve1(file string) int {
	var sum int
	util.ProcessLines(day+file, func(line string) {
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				sum += 10 * toInt(line[i])
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				sum += toInt(line[i])
				break
			}
		}
	})
	return sum
}

func part2() {
	defer util.StartTiming()()
	fmt.Printf("Part 2:\t%d\n", solve2("input.txt"))
}

func solve2(file string) int {
	var sum int
	util.ProcessLines(day+file, func(line string) {
		for {
			if unicode.IsDigit(rune(line[0])) {
				sum += 10 * toInt(line[0])
				break
			} else if val, ok := startsWithNumber(line); ok {
				sum += 10 * val
				break
			}
			line = line[1:]
		}

		for {
			if unicode.IsDigit(rune(line[len(line)-1])) {
				sum += toInt(line[len(line)-1])
				break
			} else if val, ok := endsWithNumber(line); ok {
				sum += val
				break
			}
			line = line[:len(line)-1]
		}
	})
	return sum
}

func toInt(b byte) int {
	return int(b - '0')
}

var numbers = [10]string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

func startsWithNumber(line string) (int, bool) {
	for val, text := range numbers {
		if strings.HasPrefix(line, text) {
			return val, true
		}
	}
	return 0, false
}

func endsWithNumber(line string) (int, bool) {
	for val, text := range numbers {
		if strings.HasSuffix(line, text) {
			return val, true
		}
	}
	return 0, false
}
