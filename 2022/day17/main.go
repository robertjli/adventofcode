package main

import (
	"fmt"

	"github.com/robertjli/adventofcode/2022/day17/model"
	"github.com/robertjli/adventofcode/util"
)

const debug = false

var day = fmt.Sprintf("2022/day%d/", 17)

func main() {
	rocker := model.NewRocker()

	fmt.Println("Sample Input")
	fmt.Println("------------")
	sampleJetter := model.NewJetter(day + "sample.txt")
	fmt.Print("Part 1:\t")
	util.Assert(solve1(rocker, sampleJetter, 6), 3068)
	fmt.Print("\n")
	fmt.Print("Part 2:\t")
	util.Assert(solve2(), -1)
	fmt.Print("\n\n")

	fmt.Println("Graded Input")
	fmt.Println("------------")
	jetter := model.NewJetter(day + "input.txt")
	part1(rocker, jetter, 2022)
	//part2()
}

func part1(rocker *model.Rocker, jetter *model.Jetter, rocks int) {
	stop := util.StartTiming()
	defer stop()

	fmt.Printf("Answer for part 1:\t%d\n", solve1(rocker, jetter, rocks))
}

func part2() {
	stop := util.StartTiming()
	defer stop()

	fmt.Printf("Answer for part 2:\t%d\n", solve2())
}

func solve1(rocker *model.Rocker, jetter *model.Jetter, rocks int) int {
	chamber := model.NewChamber(4, 4000)
	chamber.Print()

	for r := 0; r < rocks; r++ {
		rock := rocker.CreateRock(chamber.Height() + 3)
		chamber.SetActive(rock)
		chamber.Print()
		for {
			// jet push

			// fall
			fell := rock.FallDown(chamber)
			if !fell {
				chamber.Solidify()
				chamber.Print()
				break
			}
			chamber.Print()
		}
	}

	return chamber.Height()
}

func solve2() int {
	return -1
}
