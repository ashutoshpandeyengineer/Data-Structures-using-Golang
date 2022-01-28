package main

import (
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}
func (s *Stack) Pop() int {
	l := len(s.items)
	toremove := s.items[l]
	s.items = s.items[:l]
	return toremove

}
func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(12)
	myStack.Pop()

}
