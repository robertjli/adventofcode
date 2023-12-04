package main

import (
	"fmt"
	"math"
	"unicode"

	"github.com/robertjli/adventofcode/util"
)

/*
--- Day 3: Gear Ratios ---
You and the Elf eventually reach a gondola lift station; he says the gondola lift will take you up
to the water source, but this is as far as he can bring you. You go inside.

It doesn't take long to find the gondolas, but there seems to be a problem: they're not moving.

"Aaah!"

You turn around to see a slightly-greasy Elf with a wrench and a look of surprise. "Sorry, I wasn't
expecting anyone! The gondola lift isn't working right now; it'll still be a while before I can fix
it." You offer to help.

The engineer explains that an engine part seems to be missing from the engine, but nobody can figure
out which one. If you can add up all the part numbers in the engine schematic, it should be easy to
work out which part is missing.

The engine schematic (your puzzle input) consists of a visual representation of the engine. There
are lots of numbers and symbols you don't really understand, but apparently any number adjacent to a
symbol, even diagonally, is a "part number" and should be included in your sum. (Periods (.) do not
count as a symbol.)

Here is an example engine schematic:

```
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
```

In this schematic, two numbers are not part numbers because they are not adjacent to a symbol: 114
(top right) and 58 (middle right). Every other number is adjacent to a symbol and so is a part
number; their sum is 4361.

Of course, the actual engine schematic is much larger. What is the sum of all of the part numbers in
the engine schematic?

--- Part Two ---
The engineer finds the missing part and installs it in the engine! As the engine springs to life,
you jump in the closest gondola, finally ready to ascend to the water source.

You don't seem to be going very fast, though. Maybe something is still wrong? Fortunately, the
gondola has a phone labeled "help", so you pick it up and the engineer answers.

Before you can explain the situation, she suggests that you look out the window. There stands the
engineer, holding a phone in one hand and waving with the other. You're going so slowly that you
haven't even left the station. You exit the gondola.

The missing part wasn't the only issue - one of the gears in the engine is wrong. A gear is any *
symbol that is adjacent to exactly two part numbers. Its gear ratio is the result of multiplying
those two numbers together.

This time, you need to find the gear ratio of every gear and add them all up so that the engineer
can figure out which gear needs to be replaced.

Consider the same engine schematic again:

```
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
```

In this schematic, there are two gears. The first is in the top left; it has part numbers 467 and
35, so its gear ratio is 16345. The second gear is in the lower right; its gear ratio is 451490.
(The * adjacent to 617 is not a gear because it is only adjacent to one part number.) Adding up all
of the gear ratios produces 467835.

What is the sum of all of the gear ratios in your engine schematic?
*/

var day = util.DayPath(2023, 3)

func main() {
	fmt.Println("Sample Input")
	fmt.Println("------------")
	fmt.Print("Part 1:\t")
	util.Assert(solve1("sample.txt"), 4361)
	fmt.Print("\n")
	fmt.Print("Part 2:\t")
	util.Assert(solve2("sample.txt"), 467835)
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
	m := buildMap(day + file)

	return filterSumParts(m)
}

func part2() {
	defer util.StartTiming()()
	fmt.Printf("Part 2:\t%d\n", solve2("input.txt"))
}

func solve2(file string) int {
	m := buildMap(day + file)

	return filterSumGears(m)
}

type Entry struct {
	num int
	sym rune
}

type Coord struct {
	row, col int
}

func buildMap(file string) map[Coord]Entry {
	m := make(map[Coord]Entry)

	row := 0
	util.ProcessLines(file, func(line string) {
		for col := 0; col < len(line); col++ {
			if unicode.IsDigit(rune(line[col])) {
				end := col + 1
				for ; end < len(line) && unicode.IsDigit(rune(line[end])); end++ {
				}
				m[Coord{row, col}] = Entry{num: util.ParseInt(line[col:end])}
				col = end - 1
			} else if line[col] != '.' {
				m[Coord{row, col}] = Entry{sym: rune(line[col])}
			}
		}

		row++
	})

	return m
}

func filterSumParts(m map[Coord]Entry) int {
	var sum int

	for coord, entry := range m {
		if entry.num == 0 { // assuming zero is invalid part number
			continue
		}

		for _, c := range partNeighbors(coord, entry.num) {
			if e, ok := m[c]; ok && e.sym != 0 {
				sum += entry.num
				break
			}
		}
	}

	return sum
}

func partNeighbors(c Coord, num int) []Coord {
	neighbors := make([]Coord, 0, 12)

	x := c.row
	y := c.col
	l := int(math.Log10(float64(num))) + 1

	// top & bottom
	for i := 0; i < l; i++ {
		neighbors = append(neighbors, Coord{x - 1, y + i}, Coord{x + 1, y + i})
	}

	// left & right
	neighbors = append(neighbors,
		Coord{x - 1, y - 1}, Coord{x, y - 1}, Coord{x + 1, y - 1},
		Coord{x - 1, y + l}, Coord{x, y + l}, Coord{x + 1, y + l},
	)

	return neighbors
}

func filterSumGears(m map[Coord]Entry) int {
	potentialGears := make(map[Coord][]Entry)
	for coord, entry := range m {
		if entry.num == 0 {
			continue
		}

		for _, c := range partNeighbors(coord, entry.num) {
			if e, ok := m[c]; ok && e.sym == '*' {
				if pg, ok := potentialGears[c]; ok {
					pg = append(pg, entry)
					potentialGears[c] = pg
				} else {
					potentialGears[c] = []Entry{entry}
				}
			}
		}
	}

	var sum int
	for _, entries := range potentialGears {
		if len(entries) == 2 {
			sum += entries[0].num * entries[1].num
		}
	}
	return sum
}

func gearNeighbors(c Coord) []Coord {
	return []Coord{
		{c.row - 1, c.col - 1}, {c.row - 1, c.col}, {c.row - 1, c.col + 1},
		{c.row, c.col - 1}, {c.row, c.col + 1},
		{c.row + 1, c.col - 1}, {c.row + 1, c.col}, {c.row + 1, c.col + 1},
	}
}
