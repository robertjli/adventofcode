package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	"github.com/robertjli/adventofcode2022/util"
)

const debug = false

var day = fmt.Sprintf("day%d/", 11)

var (
	p1Rounds       = 20
	p1RelaxDivisor = 3
	p2Rounds       = 10000
	p2RelaxDivisor = 1
)

func main() {
	fmt.Println("Sample Input")
	fmt.Println("------------")
	fmt.Print("Part 1:\t")
	util.Assert(solve("sample.txt", p1Rounds, p1RelaxDivisor), 10605)
	fmt.Print("\n")
	fmt.Print("Part 2:\t")
	util.Assert(solve("sample.txt", p2Rounds, p2RelaxDivisor), 2713310158)
	fmt.Print("\n\n")

	fmt.Println("Graded Input")
	fmt.Println("------------")
	part1()
	part2()
}

type monkey struct {
	items []int
	op    func(old int) (new int)
	throw func(item int) (target int)
}

func part1() {
	stop := util.StartTiming()
	defer stop()

	fmt.Printf("Answer for part 1:\t%d\n", solve("input.txt", p1Rounds, p1RelaxDivisor))
}

func part2() {
	stop := util.StartTiming()
	defer stop()

	fmt.Printf("Answer for part 2:\t%d\n", solve("input.txt", p2Rounds, p2RelaxDivisor))
}

func solve(file string, rounds int, relaxDivisor int) int {
	if debug {
		fmt.Println("\nBuilding monkeys...")
	}
	monkeys, groupDivisor := buildMonkeys(file)

	inspections := run(monkeys, rounds, relaxDivisor, groupDivisor)
	sort.Slice(inspections, func(i, j int) bool {
		return inspections[i] > inspections[j]
	})
	total := inspections[0] * inspections[1]

	if debug {
		fmt.Printf("busiest monkeys inspected %d * %d = %d\n",
			inspections[0], inspections[1], total)
	}

	return total
}

func buildMonkeys(file string) ([]*monkey, int) {
	scanner, closeFunc := util.NewScanner(day + file)
	defer closeFunc()

	monkeys := make([]*monkey, 0, 8)
	groupDivisor := 1

	for scanner.Scan() {
		monkey := &monkey{}
		idx := util.ParseInt(strings.TrimRight(strings.Fields(scanner.Text())[1], ":"))

		scanner.Scan()
		monkey.items = getItems(scanner.Text())

		scanner.Scan()
		monkey.op = getOp(scanner.Text())

		scanner.Scan()
		test := scanner.Text()
		scanner.Scan()
		ifTrue := scanner.Text()
		scanner.Scan()
		ifFalse := scanner.Text()
		var divisor int
		monkey.throw, divisor = getThrow(test, ifTrue, ifFalse)
		groupDivisor *= divisor

		scanner.Scan()

		monkeys = append(monkeys, monkey)

		if debug {
			fmt.Printf("Monkey %d, items: %v\n", idx, monkey.items)
		}
	}

	if debug {
		fmt.Printf("Group divisor: %d\n", groupDivisor)
	}

	return monkeys, groupDivisor
}

func getItems(line string) []int {
	tokens := strings.FieldsFunc(line, func(r rune) bool {
		return !unicode.IsNumber(r)
	})
	items := make([]int, 0, len(tokens))
	for _, token := range tokens {
		items = append(items, util.ParseInt(token))
	}

	return items
}

func getOp(line string) func(int) int {
	op := line[len("  Operation: new = "):]
	if op == "old * old" {
		return func(old int) int { return old * old }
	} else if strings.HasPrefix(op, "old *") {
		x := util.ParseInt(op[len("old * "):])
		return func(old int) int { return old * x }
	} else if strings.HasPrefix(op, "old +") {
		x := util.ParseInt(op[len("old + "):])
		return func(old int) int { return old + x }
	}
	panic(fmt.Sprintf("unknown op %s", op))
}

func getThrow(test, ifTrue, ifFalse string) (func(int) int, int) {
	divisor := util.ParseInt(test[len("  Test: divisible by "):])
	trueTarget := util.ParseInt(ifTrue[len("    If true: throw to monkey "):])
	falseTarget := util.ParseInt(ifFalse[len("    If false: throw to monkey "):])

	return func(val int) int {
		if val%divisor == 0 {
			return trueTarget
		}
		return falseTarget
	}, divisor
}

func run(monkeys []*monkey, rounds int, relaxDivisor int, groupDivisor int) []int {
	inspections := make([]int, len(monkeys))

	if debug {
		fmt.Println()
	}

	for round := 1; round <= rounds; round++ {
		if debug {
			fmt.Printf("Round %d\n", round)
		}

		for m, mnk := range monkeys {
			if debug {
				fmt.Printf("  Monkey %d\n", m)
			}

			for _, item := range mnk.items {
				inspections[m] += 1

				if debug {
					fmt.Printf("    Monkey %d inspecting item %d → ", m, item)
				}

				item = mnk.op(item) / relaxDivisor

				if debug {
					fmt.Printf("value is now %d → ", item)
				}

				target := mnk.throw(item)
				monkeys[target].items = append(monkeys[target].items, item%groupDivisor)

				if debug {
					fmt.Printf("thrown to monkey %d\n", target)
				}
			}
			mnk.items = nil
		}

		if debug {
			for m, mnk := range monkeys {
				fmt.Printf("  Monkey %d now has items: %v\n", m, mnk.items)
			}
			for i, ins := range inspections {
				fmt.Printf("  Monkey %d has inspected %d times\n", i, ins)
			}
		}
	}

	return inspections
}
