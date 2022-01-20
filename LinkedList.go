package main

import "fmt"

//structure of a node in a Linked list.
type node struct {
	data int //refers to data and dataype
	next *node
} // next refers to the pointer of node.

//structure of a Linked list.
type LinkedList struct {
	head *node
	len  int
}

//function for Inserting value At last.
func (l *LinkedList) Insertatlast(value int) {
	n := node{}
	n.data = value
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.len++
			return
		}
		ptr = ptr.next
	}
}
func (l *LinkedList) Insertatfirst(value int) {
	n := node{}
	n.data = value
	if l.head == nil {
		l.head = &n
		l.len++
		return
	}
	n.next = l.head
	l.head = &n
	l.len++
}
func (l *LinkedList) GetAt(pos int) *node {
	ptr := l.head
	if pos < 0 {
		return ptr
	}
	if pos > (l.len - 1) {
		return nil
	}
	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	return ptr
}

func (l *LinkedList) Insertatparticularposition(pos int, value int) {
	n := node{}
	n.data = value
	if pos < 0 {
		fmt.Println("Please Enter a Non Negetive value")
	}
	if pos == 0 {
		l.head = &n
		l.len++
		return
	}
	if pos > l.len {
		fmt.Println("The position is out of range ")
		return
	}
	m := l.GetAt(pos)
	n.next = m
	o := l.GetAt(pos - 1)
	o.next = &n
	l.len++
}
func (l *LinkedList) Deleteatparticularposition(pos int) {
	//	n:=node{}
	if pos < 0 {
		fmt.Println("position can not be negative")
	}
	p := l.GetAt(pos - 1)
	if p == nil {
		fmt.Println("Node not found")
	}
	p.next = l.GetAt(pos).next
	l.len--
	return
}

func (l *LinkedList) RemoveFront() {
	if l.head == nil {
		fmt.Println("Error no elements in list")
	}
	l.head = l.head.next
	l.len--
	return
}

func (l LinkedList) Print() {
	if l.len == 0 {
		fmt.Println("No nodes in list")
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		fmt.Println("node:", i, ":->", ptr.data)
		ptr = ptr.next
	}
}

func main() {
	fmt.Println("hello")
	l := LinkedList{}

	l.Insertatfirst(2)
	l.Insertatfirst(12)
	l.Insertatlast(15)
	l.Insertatlast(17)
	l.Print()
	//	l.Deleteatparticularposition(3)   //working fine
	//l.RemoveFront()                       //working fine
	l.Insertatparticularposition(2, 40)
	l.Print()

}
