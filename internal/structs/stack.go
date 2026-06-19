package structs

import (
	"fmt"
)

type Stack[T any] struct {
	st []T
}

func BuildStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(el T) {
	s.st = append(s.st, el)
}

func (s *Stack[T]) Pop() T {
	el := s.st[len(s.st)-1]
	s.st = s.st[:len(s.st)-1]
	return el
}

func (s *Stack[T]) Count() int {
	return len(s.st)
}

func (s *Stack[T]) Print() {
	fmt.Println("----------")
	for i := len(s.st) - 1; i >= 0; i-- {
		fmt.Printf("%v\n", s.st[i])
	}
	fmt.Println("----------")
}
