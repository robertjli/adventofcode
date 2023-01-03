package model

var _ Rock = &Plus{}

type Plus struct {
	center Point
}

func NewPlus(row int) Rock {
	return &Plus{Point{row + 1, 3}}
}

func (p *Plus) PushLeft(chamber *Chamber) bool {
	//TODO implement me
	panic("implement me")
}

func (p *Plus) PushRight(chamber *Chamber) bool {
	//TODO implement me
	panic("implement me")
}

func (p *Plus) FallDown(chamber *Chamber) bool {
	if p.center.row == 1 {
		return false
	}

	if chamber.grid[p.center.row-2][p.center.col] != Empty ||
		chamber.grid[p.center.row-1][p.center.col-1] != Empty ||
		chamber.grid[p.center.row-1][p.center.col+1] != Empty {
		return false
	}

	p.center = Point{p.center.row - 1, p.center.col}
	return true
}

func (p *Plus) Positions() []Point {
	return []Point{
		{p.center.row, p.center.col},
		{p.center.row + 1, p.center.col},
		{p.center.row - 1, p.center.col},
		{p.center.row, p.center.col + 1},
		{p.center.row, p.center.col - 1},
	}
}

func (p *Plus) Contains(point Point) bool {
	if point.row == p.center.row &&
		(point.col >= p.center.col-1 && point.col <= p.center.col+1) {
		return true
	}

	if point.col == p.center.col &&
		(point.row >= p.center.row-1 && point.row <= p.center.row+1) {
		return true
	}

	return false
}

func (p *Plus) Top() int {
	return p.center.row + 1
}

func (p *Plus) Bottom() int {
	return p.center.row - 1
}

func (p *Plus) Height() int {
	return 3
}
