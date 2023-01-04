package model

var _ Rock = &Dash{}

type Dash struct {
	origin Point
}

func NewDash(row int) Rock {
	return &Dash{Point{row, 2}}
}

func (d *Dash) PushLeft(chamber *Chamber) bool {
	if d.origin.col == 0 || chamber.grid[d.origin.row][d.origin.col-1] != Empty {
		return false
	}

	d.origin = Point{d.origin.row, d.origin.col - 1}
	return true
}

func (d *Dash) PushRight(chamber *Chamber) bool {
	if d.origin.col == Width-4 || chamber.grid[d.origin.row][d.origin.col+4] != Empty {
		return false
	}

	d.origin = Point{d.origin.row, d.origin.col + 1}
	return true
}

func (d *Dash) FallDown(chamber *Chamber) bool {
	if d.origin.row == 0 {
		return false
	}

	for i := 0; i < 4; i++ {
		if chamber.grid[d.origin.row-1][d.origin.col+i] != Empty {
			return false
		}
	}

	d.origin = Point{d.origin.row - 1, d.origin.col}
	return true
}

func (d *Dash) Positions() []Point {
	p := make([]Point, 0, 4)
	for i := 0; i < 4; i++ {
		p = append(p, Point{d.origin.row, d.origin.col + i})
	}
	return p
}

func (d *Dash) Contains(point Point) bool {
	return point.row == d.origin.row && point.col >= d.origin.col && point.col < d.origin.col+4
}

func (d *Dash) Top() int {
	return d.origin.row
}

func (d *Dash) Bottom() int {
	return d.origin.row
}

func (d *Dash) Height() int {
	return 1
}
