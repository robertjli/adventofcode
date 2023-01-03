package model

var _ Rock = &Box{}

type Box struct {
	origin Point // bottom left
}

func NewBox(row int) Rock {
	return &Box{Point{row, 2}}
}

func (b *Box) PushLeft(chamber *Chamber) bool {
	//TODO implement me
	panic("implement me")
}

func (b *Box) PushRight(chamber *Chamber) bool {
	//TODO implement me
	panic("implement me")
}

func (b *Box) FallDown(chamber *Chamber) bool {
	if b.origin.row == 0 ||
		chamber.grid[b.origin.row-1][b.origin.col] != Empty ||
		chamber.grid[b.origin.row-1][b.origin.col+1] != Empty {
		return false
	}

	b.origin = Point{b.origin.row - 1, b.origin.col}
	return true
}

func (b *Box) Positions() []Point {
	return []Point{
		{b.origin.row, b.origin.col},
		{b.origin.row, b.origin.col + 1},
		{b.origin.row + 1, b.origin.col},
		{b.origin.row + 1, b.origin.col + 1},
	}
}

func (b *Box) Contains(point Point) bool {
	return (point.row == b.origin.row || point.row == b.origin.row+1) &&
		(point.col == b.origin.col || point.col == b.origin.col+1)
}

func (b *Box) Top() int {
	return b.origin.row + 1
}

func (b *Box) Bottom() int {
	return b.origin.row
}

func (b *Box) Height() int {
	return 1
}
