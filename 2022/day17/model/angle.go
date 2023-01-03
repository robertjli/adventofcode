package model

var _ Rock = &Angle{}

type Angle struct {
	joint Point
}

func NewAngle(row int) Rock {
	return &Angle{Point{row, 4}}
}

func (a *Angle) PushLeft(chamber *Chamber) bool {
	//TODO implement me
	panic("implement me")
}

func (a *Angle) PushRight(chamber *Chamber) bool {
	//TODO implement me
	panic("implement me")
}

func (a *Angle) FallDown(chamber *Chamber) bool {
	if a.joint.row == 0 {
		return false
	}

	if chamber.grid[a.joint.row-1][a.joint.col] != Empty ||
		chamber.grid[a.joint.row-1][a.joint.col-1] != Empty ||
		chamber.grid[a.joint.row-1][a.joint.col-2] != Empty {
		return false
	}

	a.joint = Point{a.joint.row - 1, a.joint.col}
	return true
}

func (a *Angle) Positions() []Point {
	return []Point{
		{a.joint.row, a.joint.col},
		{a.joint.row, a.joint.col - 1},
		{a.joint.row, a.joint.col - 2},
		{a.joint.row + 1, a.joint.col},
		{a.joint.row + 2, a.joint.col},
	}
}

func (a *Angle) Contains(point Point) bool {
	if point.row == a.joint.row &&
		(point.col >= a.joint.col-2 && point.col <= a.joint.col) {
		return true
	}

	if point.col == a.joint.col &&
		(point.row >= a.joint.row && point.row <= a.joint.row+2) {
		return true
	}

	return false
}

func (a *Angle) Top() int {
	return a.joint.row + 2
}

func (a *Angle) Bottom() int {
	return a.joint.row
}

func (a *Angle) Height() int {
	return 3
}
