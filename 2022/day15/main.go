package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"

	"github.com/robertjli/adventofcode/util"
)

const debug = false

var day = fmt.Sprintf("2022/day%d/", 15)

func main() {
	fmt.Println("Sample Input")
	fmt.Println("------------")
	fmt.Print("Part 1:\t")
	util.Assert(solve1("sample.txt", 10), 26)
	fmt.Print("\n")
	fmt.Print("Part 2:\t")
	util.Assert(solve2b("sample.txt", 20), 56000011)
	fmt.Print("\n\n")

	fmt.Println("Graded Input")
	fmt.Println("------------")
	part1("input.txt")
	part2("input.txt")
}

func part1(file string) {
	stop := util.StartTiming()
	defer stop()

	fmt.Printf("Answer for part 1:\t%d\n", solve1(file, 2000000))
}

func part2(file string) {
	stop := util.StartTiming()
	defer stop()

	fmt.Printf("Answer for part 2:\t%d\n", solve2b(file, 4000000))
}

type pos rune

const sensor pos = 'S'
const beacon pos = 'B'
const empty pos = '#'

func solve1(file string, row int) int {
	coverage := make(map[int]pos)

	util.ProcessLines(day+file, func(line string) {
		parts := strings.FieldsFunc(line, func(r rune) bool {
			return !unicode.IsDigit(r) && r != '-'
		})
		sx, sy := util.ParseInt(parts[0]), util.ParseInt(parts[1])
		bx, by := util.ParseInt(parts[2]), util.ParseInt(parts[3])
		if sy == row {
			coverage[sx] = sensor
		}
		if by == row {
			coverage[bx] = beacon
		}

		// how far the sensor is from the beacon
		dist := int(math.Abs(float64(sx-bx)) + math.Abs(float64(sy-by)))

		// how much of the target row is covered by this sensor
		spread := dist - int(math.Abs(float64(sy-row)))
		if spread < 0 {
			// the sensor is too far from the target row, we're done
			return
		}

		for i := sx - spread; i <= sx+spread; i++ {
			if _, ok := coverage[i]; !ok {
				coverage[i] = empty
			}
		}
	})

	nonBeacon := 0
	for _, r := range coverage {
		if r == sensor || r == empty {
			nonBeacon++
		}
	}
	return nonBeacon
}

type point struct {
	x, y int
}

func solve2(file string, cap int) int {
	if debug {
		fmt.Println()
	}

	coverage := make(map[point]pos)

	util.ProcessLines(day+file, func(line string) {
		parts := strings.FieldsFunc(line, func(r rune) bool {
			return !unicode.IsDigit(r) && r != '-'
		})
		sx, sy := util.ParseInt(parts[0]), util.ParseInt(parts[1])
		bx, by := util.ParseInt(parts[2]), util.ParseInt(parts[3])
		coverage[point{sx, sy}] = sensor
		coverage[point{bx, by}] = beacon

		// how far the sensor is from the beacon
		dist := int(math.Abs(float64(sx-bx)) + math.Abs(float64(sy-by)))

		if debug {
			fmt.Printf("sensor: (%d, %d), beacon: (%d, %d), distance: %d\n", sx, sy, bx, by, dist)
		}

		x := int(math.Max(0, float64(sx-dist)))
		dy := int(math.Max(0, float64(dist-sx)))
		if sx == 2 && sy == 0 {
			fmt.Printf("x: %d, dy: %d\n", x, dy)
		}
		for x <= sx+dist && x <= cap {
			y := int(math.Max(0, float64(sy-dy)))
			for ; y <= sy+dy && y <= cap; y++ {

				if sx == 2 && sy == 0 {
					fmt.Printf("considering (%d, %d), which is %c\n", x, y, coverage[point{x, y}])
				}

				if _, ok := coverage[point{x, y}]; !ok {
					coverage[point{x, y}] = empty
				}
			}
			x++
			if x <= sx {
				dy++
			} else {
				dy--
			}
		}
	})

	for x := 0; x <= cap; x++ {
		for y := 0; y <= cap; y++ {
			if _, ok := coverage[point{x, y}]; !ok {
				if debug {
					fmt.Printf("found beacon! (%d, %d)\n", x, y)
				}

				return x*4000000 + y
			}
		}
	}
	return -1
}

// for each possible location, iterate the list of sensors until it's too close to one
func solve2b(file string, cap int) int {
	if debug {
		fmt.Println()
	}

	sensors := make(map[point]int)      // k: sensor location v: distance ruled out
	beacons := make(map[point]struct{}) // skip any locations that have beacons

	util.ProcessLines(day+file, func(line string) {
		parts := strings.FieldsFunc(line, func(r rune) bool {
			return !unicode.IsDigit(r) && r != '-'
		})
		sx, sy := util.ParseInt(parts[0]), util.ParseInt(parts[1])
		bx, by := util.ParseInt(parts[2]), util.ParseInt(parts[3])

		// how far the sensor is from the beacon
		dist := int(math.Abs(float64(sx-bx)) + math.Abs(float64(sy-by)))

		sensors[point{sx, sy}] = dist
		beacons[point{bx, by}] = struct{}{}

		if debug {
			fmt.Printf("sensor: (%d, %d), beacon: (%d, %d), distance: %d\n", sx, sy, bx, by, dist)
		}
	})

	for row := 0; row <= cap; row++ {
		if debug {
			fmt.Printf("processing row %d\n", row)
		}

		for col := 0; col <= cap; col++ {
			if _, ok := beacons[point{row, col}]; ok {
				continue
			}
			found := false
			skip := 0
			for sns, dist := range sensors {
				found, skip = withinRange(sns, dist, point{col, row})
				if found {
					if debug {
						fmt.Printf(
							"\tpoint (%d, %d) within range of sensor (%d, %d), skipping %d\n",
							col, row, sns.x, sns.y, skip)
					}
					break
				}
			}
			if !found {
				if debug {
					fmt.Printf("no sensor covered (%d, %d)\n", col, row)
				}
				return col*4000000 + row
			}
			col += skip
		}
	}

	return -1
}

// withinRange checks if p is within range of the given sensor s.
// If so, it returns true and the number of additional spaces in the same row that are covered by
// that sensor. Otherwise, it returns false.
func withinRange(s point, srange int, p point) (bool, int) {
	dx := int(math.Abs(float64(s.x - p.x)))
	dy := int(math.Abs(float64(s.y - p.y)))
	dist := dx + dy
	return dist <= srange, dx + srange - dy
}
