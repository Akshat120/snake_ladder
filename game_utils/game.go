package gameutils

import (
	"fmt"
	"time"
)

const MAX_POS = 100

type Game interface {
	Run()
}

type game struct {
	snakes  []Snake
	ladders []Ladder
	players []Player
}

var gameInstance *game

func InitGame(snakes []Snake, ladders []Ladder, players []Player) Game {
	if gameInstance != nil {
		return gameInstance
	}

	gameInstance = &game{
		snakes:  snakes,
		ladders: ladders,
		players: players,
	}

	return gameInstance
}

func (g *game) Run() {

	for {

		for _, p := range g.players {
			time.Sleep(500 * time.Millisecond)
			r := p.RollDice()
			newPos := p.GetPos() + r

			if newPos <= MAX_POS {

				fmt.Printf("%v rolled a %v and moved from %v to %v\n", p.GetName(), r, p.GetPos(), newPos)

				if afterBitePos, ok := g.checkSnakeBite(newPos); ok {
					fmt.Printf("%v got snake bitten and drops to %v\n", p.GetName(), afterBitePos)
					newPos = afterBitePos
				}

				if afterJumpPos, ok := g.checkLadderJump(newPos); ok {
					fmt.Printf("%v got ladder jump and tops to %v\n", p.GetName(), afterJumpPos)
					newPos = afterJumpPos
				}

				p.UpdatePos(newPos)

				if p.GetPos() == MAX_POS {
					fmt.Printf("%v Wins the game.\n", p.GetName())
					return
				}
			} else {
				fmt.Printf("%v rolled a %v and can't go outside board from %v\n", p.GetName(), r, p.GetPos())
			}

		}
	}
}

func (g *game) checkSnakeBite(pos int) (int, bool) {

	for _, snake := range g.snakes {
		if snake.GetHead() == pos {
			return snake.GetTail(), true
		}
	}

	return -1, false
}

func (g *game) checkLadderJump(pos int) (int, bool) {

	for _, ladder := range g.ladders {
		if ladder.GetBottom() == pos {
			return ladder.GetTop(), true
		}
	}

	return -1, false
}
