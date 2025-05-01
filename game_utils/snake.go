package gameutils

import "fmt"

type Snake interface {
	Print()
	GetHead() int
	GetTail() int
}

type snake struct {
	head int
	tail int
}

func InitSnake(head, tail int) Snake {
	return &snake{
		head: head,
		tail: tail,
	}
}

func (s *snake) GetHead() int {
	return s.head
}

func (s *snake) GetTail() int {
	return s.tail
}

func (s *snake) Print() {
	fmt.Printf("head: %v; tail: %v\n", s.head, s.tail)
}
