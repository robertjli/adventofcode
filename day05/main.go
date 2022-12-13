package main

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"

	"github.com/gammazero/deque"

	"github.com/robertjli/adventofcode2022/util"
)

const debug = false

var day = fmt.Sprintf("day%d/", 5)

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
	solve(file, 1, func(stacks []*deque.Deque[rune], size int, from int, to int) {
		for i := 0; i < size; i++ {
			elem := stacks[from].PopFront()
			stacks[to].PushFront(elem)
		}
	})
}

func part2(file string) {
	solve(file, 2, func(stacks []*deque.Deque[rune], size int, from int, to int) {
		temp := deque.New[rune](size)
		for i := 0; i < size; i++ {
			elem := stacks[from].PopFront()
			temp.PushFront(elem)
		}
		for i := 0; i < size; i++ {
			elem := temp.PopFront()
			stacks[to].PushFront(elem)
		}
	})
}

type moveFunc func([]*deque.Deque[rune], int, int, int)

func solve(file string, part int, moveFunc moveFunc) {
	stop := util.StartTiming()
	defer stop()

	scanner, closeFunc := util.NewScanner(day + file)
	defer closeFunc()
	scanner.Scan()

	stackLines := make([]string, 0, 8)
	line := scanner.Text()
	for len(line) > 0 {
		stackLines = append(stackLines, line)
		scanner.Scan()
		line = scanner.Text()
	}
	stacks := build(stackLines)
	stacks = move(stacks, scanner, moveFunc)

	var tops string
	for _, stack := range stacks {
		tops += string(stack.Front())
	}

	fmt.Printf("Answer for part %d:\t%s\n", part, tops)
}

func build(lines []string) []*deque.Deque[rune] {
	count := 1 + ((len(lines[len(lines)-1]) - 2) / 4) // this works trust me
	stacks := make([]*deque.Deque[rune], 0, count)
	for i := 0; i < count; i++ {
		stacks = append(stacks, deque.New[rune]())
	}

	for i := len(lines) - 2; i >= 0; i-- {
		line := lines[i]
		idx := 0
		for j := 1; j < len(line); j += 4 {
			if line[j] != ' ' {
				stacks[idx].PushFront(rune(line[j]))
			}
			idx++
		}
	}

	if debug {
		fmt.Println("Parsed stacks (top to bottom):")
		print(stacks)
	}

	return stacks
}

func print(stacks []*deque.Deque[rune]) {
	for i, stack := range stacks {
		fmt.Printf("Stack %d: ", i)
		for j := 0; j < stack.Len(); j++ {
			if j != 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%c", stack.At(j))
		}
		fmt.Println()
	}
}

func move(stacks []*deque.Deque[rune], scanner *bufio.Scanner, moveFunc moveFunc) []*deque.Deque[rune] {
	for scanner.Scan() {
		tokens := strings.FieldsFunc(scanner.Text(), func(r rune) bool {
			return !unicode.IsNumber(r)
		})

		size := util.ParseInt(tokens[0])
		from := util.ParseInt(tokens[1]) - 1
		to := util.ParseInt(tokens[2]) - 1

		if debug {
			fmt.Printf("Moving %d elements from stack %d to stack %d\n", size, from, to)
		}

		moveFunc(stacks, size, from, to)

		if debug {
			print(stacks)
		}
	}

	return stacks
}
