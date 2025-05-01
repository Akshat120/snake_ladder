package gameutils

import (
	"fmt"
	"math/rand"
)

type Player interface {
	Print()
	GetName() string
	GetPos() int
	RollDice() int
	UpdatePos(int)
}

type player struct {
	name    string
	currPos int
}

func InitPlayer(name string) Player {
	return &player{
		name:    name,
		currPos: 0,
	}
}

func (p *player) GetName() string {
	return p.name
}

func (p *player) GetPos() int {
	return p.currPos
}

func (p *player) Print() {
	fmt.Printf("%#v\n", p)
}

func (p *player) RollDice() int {
	return (rand.Intn(6) + 1)
}

func (p *player) UpdatePos(newPos int) {
	p.currPos = newPos
}
