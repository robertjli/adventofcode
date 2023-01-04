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
	util.Assert(solve(rocker, sampleJetter, 2022), 3068)
	fmt.Print("\n")
	fmt.Print("Part 2:\t")
	util.Assert(solve(rocker, sampleJetter, 1000000000000), 1514285714288)
	fmt.Print("\n\n")

	fmt.Println("Graded Input")
	fmt.Println("------------")
	jetter := model.NewJetter(day + "input.txt")
	part1(rocker, jetter, 2022)
	part2(rocker, jetter, 1000000000000)
}

func part1(rocker *model.Rocker, jetter *model.Jetter, rocks int) {
	stop := util.StartTiming()
	defer stop()

	fmt.Printf("Answer for part 1:\t%d\n", solve(rocker, jetter, rocks))
}

func part2(rocker *model.Rocker, jetter *model.Jetter, rocks int) {
	stop := util.StartTiming()
	defer stop()

	fmt.Printf("Answer for part 2:\t%d\n", solve(rocker, jetter, rocks))
}

func solve(rocker *model.Rocker, jetter *model.Jetter, rocks int) int {
	fmt.Println(rocks)
	rocker.Reset()
	jetter.Reset()
	chamber := model.NewChamber(4, 4000)

	if debug {
		chamber.Print()
	}

	for r := 0; r < rocks; r++ {
		rock := rocker.CreateRock(chamber.Height() + 3)
		chamber.SetActive(rock)

		if debug {
			chamber.Print("spawned rock")
		}

		for {
			// jet push
			dir := jetter.PushRock(rock, chamber)

			if debug {
				chamber.Print(fmt.Sprintf("pushed rock %c", dir))
			}

			// fall
			fell := rock.FallDown(chamber)
			if !fell {
				chamber.Solidify()

				if debug {
					chamber.Print("solidified")
				}

				break
			}

			if debug {
				chamber.Print("fell one")
			}
		}
	}

	return chamber.Height()
}
