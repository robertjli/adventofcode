package model

var _ Rock = &Pipe{}

type Pipe struct {
	bottom Point
}

func NewPipe(row int) Rock {
	return &Pipe{Point{row, 2}}
}

func (p *Pipe) PushLeft(chamber *Chamber) bool {
	if p.bottom.col == 0 {
		return false
	}

	for i := 0; i < 4; i++ {
		if chamber.grid[p.bottom.row+i][p.bottom.col-1] != Empty {
			return false
		}
	}

	p.bottom = Point{p.bottom.row, p.bottom.col - 1}
	return true
}

func (p *Pipe) PushRight(chamber *Chamber) bool {
	if p.bottom.col == Width-1 {
		return false
	}

	for i := 0; i < 4; i++ {
		if chamber.grid[p.bottom.row+i][p.bottom.col+1] != Empty {
			return false
		}
	}

	p.bottom = Point{p.bottom.row, p.bottom.col + 1}
	return true
}

func (p *Pipe) FallDown(chamber *Chamber) bool {
	if p.bottom.row == 0 || chamber.grid[p.bottom.row-1][p.bottom.col] != Empty {
		return false
	}

	p.bottom = Point{p.bottom.row - 1, p.bottom.col}
	return true
}

func (p *Pipe) Positions() []Point {
	return []Point{
		{p.bottom.row, p.bottom.col},
		{p.bottom.row + 1, p.bottom.col},
		{p.bottom.row + 2, p.bottom.col},
		{p.bottom.row + 3, p.bottom.col},
	}
}

func (p *Pipe) Contains(point Point) bool {
	return point.col == p.bottom.col && point.row >= p.bottom.row && point.row <= p.bottom.row+3
}

func (p *Pipe) Top() int {
	return p.bottom.row + 3
}

func (p *Pipe) Bottom() int {
	return p.bottom.row
}

func (p *Pipe) Height() int {
	return 4
}
