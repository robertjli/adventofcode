package model

import (
	"fmt"

	"github.com/robertjli/adventofcode/util"
)

const (
	colorReset   = "\033[0m"
	colorWall    = "\033[31m"
	colorFalling = "\033[32m"
	colorStopped = "\033[33m"
	colorEmpty   = "\033[37m"
)

type Unit rune

const Empty Unit = '.'
const Stopped Unit = '#'
const Falling Unit = '@'

type Point struct {
	row, col int
}

const Width = 7

type Chamber struct {
	grid   [][7]Unit // first element is bottom row, last element is top
	height int       // height of rock formation
	active Rock      // currently falling rock
}

func NewChamber(rows, cap int) *Chamber {
	c := &Chamber{grid: make([][Width]Unit, 0, cap)}
	for i := 0; i < rows; i++ {
		c.AddRow()
	}

	return c
}

func (c *Chamber) AddRow() {
	row := [Width]Unit{Empty, Empty, Empty, Empty, Empty, Empty, Empty}
	c.grid = append(c.grid, row)
}

func (c *Chamber) SetActive(rock Rock) {
	if c.active != nil {
		panic(fmt.Sprintf("overwriting active rock %s", c.active))
	}
	c.active = rock

	for len(c.grid) < rock.Top()+1 {
		c.AddRow()
	}
}

func (c *Chamber) Height() int {
	return c.height
}

func (c *Chamber) Top() [Width]Unit {
	if c.height == 0 {
		return c.grid[0]
	}
	return c.grid[c.height-1]
}

func (c *Chamber) Solidify() {
	for _, p := range c.active.Positions() {
		c.grid[p.row][p.col] = Stopped

		if p.row+1 > c.height {
			c.height = p.row + 1
		}
	}

	c.active = nil
}

func (c *Chamber) Print() {
	fmt.Println()
	for i := len(c.grid) - 1; i >= 0; i-- {
		row := c.grid[i]
		fmt.Print(colorWall, "|")
		for j, cell := range row {
			if c.active != nil && c.active.Contains(Point{i, j}) {
				fmt.Print(colorFalling, string(Falling))
			} else if cell == Stopped {
				fmt.Print(colorStopped, string(cell))
			} else {
				fmt.Print(colorEmpty, string(cell))
			}
		}
		fmt.Print(colorWall, "|\n")
	}
	fmt.Print(colorWall, "+-------+\n", colorReset)
}

type Dir rune

const Left Dir = '<'
const Right Dir = '>'

type Jetter struct {
	order []Dir
	index int
}

func NewJetter(file string) *Jetter {
	line := util.ReadFile(file)[0]
	dirs := make([]Dir, 0, len(line))
	for _, char := range line {
		dirs = append(dirs, Dir(char))
	}

	return &Jetter{
		order: dirs,
		index: 0,
	}
}
