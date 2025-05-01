package gameutils

import "fmt"

type Ladder interface {
	Print()
	GetBottom() int
	GetTop() int
}

type ladder struct {
	bottom int
	top    int
}

func InitLadder(bottom, top int) Ladder {
	return &ladder{
		bottom: bottom,
		top:    top,
	}
}

func (l *ladder) GetBottom() int {
	return l.bottom
}

func (l *ladder) GetTop() int {
	return l.top
}

func (l *ladder) Print() {
	fmt.Printf("bottom: %v; top: %v\n", l.bottom, l.top)
}
