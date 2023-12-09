package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/robertjli/adventofcode/util"
)

/*
--- Day 8: Haunted Wasteland ---
You're still riding a camel across Desert Island when you spot a sandstorm quickly approaching. When
you turn to warn the Elf, she disappears before your eyes! To be fair, she had just finished warning
you about ghosts a few minutes ago.

One of the camel's pouches is labeled "maps" - sure enough, it's full of documents (your puzzle
input) about how to navigate the desert. At least, you're pretty sure that's what they are; one of
the documents contains a list of left/right instructions, and the rest of the documents seem to
describe some kind of network of labeled nodes.

It seems like you're meant to use the left/right instructions to navigate the network. Perhaps if
you have the camel follow the same instructions, you can escape the haunted wasteland!

After examining the maps for a bit, two nodes stick out: AAA and ZZZ. You feel like AAA is where
you are now, and you have to follow the left/right instructions until you reach ZZZ.

This format defines each node of the network individually. For example:

```
RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
```

Starting with AAA, you need to look up the next element based on the next left/right instruction in
your input. In this example, start with AAA and go right (R) by choosing the right element of AAA,
CCC. Then, L means to choose the left element of CCC, ZZZ. By following the left/right instructions,
you reach ZZZ in 2 steps.

Of course, you might not find ZZZ right away. If you run out of left/right instructions, repeat the
whole sequence of instructions as necessary: RL really means RLRLRLRLRLRLRLRL... and so on. For
example, here is a situation that takes 6 steps to reach ZZZ:

```
LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
```

Starting at AAA, follow the left/right instructions. How many steps are required to reach ZZZ?

--- Part Two ---
The sandstorm is upon you and you aren't any closer to escaping the wasteland. You had the camel
follow the instructions, but you've barely left your starting position. It's going to take
significantly more steps to escape!

What if the map isn't for people - what if the map is for ghosts? Are ghosts even bound by the laws
of spacetime? Only one way to find out.

After examining the maps a bit longer, your attention is drawn to a curious fact: the number of
nodes with names ending in A is equal to the number ending in Z! If you were a ghost, you'd probably
just start at every node that ends with A and follow all of the paths at the same time until they
all simultaneously end up at nodes that end with Z.

For example:

```
LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
```

Here, there are two starting nodes, 11A and 22A (because they both end with A). As you follow each
left/right instruction, use that instruction to simultaneously navigate away from both nodes you're
currently on. Repeat this process until all of the nodes you're currently on end with Z. (If only
some of the nodes you're on end with Z, they act like any other node and you continue as normal.) In
this example, you would proceed as follows:

- Step 0: You are at 11A and 22A.
- Step 1: You choose all of the left paths, leading you to 11B and 22B.
- Step 2: You choose all of the right paths, leading you to 11Z and 22C.
- Step 3: You choose all of the left paths, leading you to 11B and 22Z.
- Step 4: You choose all of the right paths, leading you to 11Z and 22B.
- Step 5: You choose all of the left paths, leading you to 11B and 22C.
- Step 6: You choose all of the right paths, leading you to 11Z and 22Z.

So, in this example, you end up entirely on nodes that end in Z after 6 steps.

Simultaneously start on every node that ends with A. How many steps does it take before you're only
on nodes that end with Z?
*/

var day = util.DayPath(2023, 8)

func main() {
	fmt.Println("Sample Input")
	fmt.Println("------------")
	fmt.Print("Part 1:\t")
	util.Assert(solve1("sample1.txt"), 2)
	util.Assert(solve1("sample2.txt"), 6)
	fmt.Print("\n")
	fmt.Print("Part 2:\t")
	util.Assert(solve2("sample1.txt"), 2)
	util.Assert(solve2("sample2.txt"), 6)
	util.Assert(solve2("sample3.txt"), 6)
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
	sc, close := util.NewScanner(day + file)
	defer close()

	sc.Scan()
	instr := sc.Text()
	sc.Scan()

	ntwk, _ := parseGraph(sc)

	return navigate("AAA", ntwk, instr)
}

func part2() {
	defer util.StartTiming()()
	fmt.Printf("Part 2:\t%d\n", solve2("input.txt"))
}

func solve2(file string) int {
	sc, close := util.NewScanner(day + file)
	defer close()

	sc.Scan()
	instr := sc.Text()
	sc.Scan()

	ntwk, starts := parseGraph(sc)

	steps := make([]int, len(starts))
	for i, start := range starts {
		steps[i] = navigate(start, ntwk, instr)
	}

	return lcms(steps)
}

func parseGraph(sc *bufio.Scanner) (graph map[string]string, starts []string) {
	graph = make(map[string]string)
	starts = make([]string, 0, 6)
	for sc.Scan() {
		line := sc.Text()
		k, v, _ := strings.Cut(line, " = ")
		l, r, _ := strings.Cut(v, ", ")
		graph[k+"L"] = l[1:]
		graph[k+"R"] = r[:len(r)-1]

		if strings.HasSuffix(k, "A") {
			starts = append(starts, k)
		}
	}
	return graph, starts
}

func navigate(start string, ntwk map[string]string, instr string) int {
	steps := 0
	curr := start
	for {
		curr = ntwk[curr+string(instr[steps%len(instr)])]
		steps++
		if strings.HasSuffix(curr, "Z") {
			break
		}
	}
	return steps
}

func lcms(xs []int) int {
	mult := xs[0]
	for _, x := range xs {
		mult = lcm(mult, x)
	}
	return mult
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
