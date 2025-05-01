package main

import (
	"fmt"
	"math/rand"
	gu "snake_ladder_1/game_utils"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var snakes []gu.Snake
	var ladders []gu.Ladder
	var players []gu.Player

	fmt.Println("Welcome to Snake Ladder I")
	fmt.Println("-------------------------")
	var noOfSnake int
	fmt.Scan(&noOfSnake)

	for range noOfSnake {
		var head, tail int
		fmt.Scan(&head)
		fmt.Scan(&tail)
		snake := gu.InitSnake(head, tail)
		snakes = append(snakes, snake)
		snake.Print()
	}

	var noOfLadder int
	fmt.Scan(&noOfLadder)

	for range noOfLadder {
		var bottom, top int
		fmt.Scan(&bottom)
		fmt.Scan(&top)
		ladder := gu.InitLadder(bottom, top)
		ladders = append(ladders, ladder)
		ladder.Print()
	}

	var noOfPlayer int
	fmt.Scan(&noOfPlayer)

	for range noOfPlayer {
		var name string
		fmt.Scan(&name)
		player := gu.InitPlayer(name)
		players = append(players, player)
		player.Print()
	}

	game := gu.InitGame(snakes, ladders, players)

	game.Run()

}
