package main

import "fmt"

type node struct {
	data string
	next *node
}
type linkedlist struct {
	head *node
	len  int
}

func (l *linkedlist) Insertatfirst(val string) {
	n := node{}
	n.data = val
	if l.head == nil {
		l.head = &n
		return
	}
	n.next = l.head
	l.head = &n
	l.len++
}
func (l linkedlist) Print() {
	if l.len == 0 {
		fmt.Println("No nodes in list")
	}
	ptr := l.head
	for i := 0; i <= l.len; i++ {
		fmt.Println("Node:->", ptr.data)
		ptr = ptr.next
	}
}
func main() {
	fmt.Println("Hello")
	n := linkedlist{}
	n.Insertatfirst("Hy")
	n.Insertatfirst("Hi")
	n.Insertatfirst("H")

	n.Print()

}
