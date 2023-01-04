package model

type Rock interface {
	PushLeft(*Chamber) bool
	PushRight(*Chamber) bool
	FallDown(*Chamber) bool
	Positions() []Point
	Contains(Point) bool
	Top() int
	Bottom() int
	Height() int
}

type Rocker struct {
	order []func(row int) Rock
	index int
}

func NewRocker() *Rocker {
	return &Rocker{
		order: []func(row int) Rock{NewDash, NewPlus, NewAngle, NewPipe, NewBox},
		index: 0,
	}
}

func (r *Rocker) CreateRock(row int) Rock {
	rock := r.order[r.index](row)
	r.index = (r.index + 1) % len(r.order)
	return rock
}

func (r *Rocker) Reset() {
	r.index = 0
}
