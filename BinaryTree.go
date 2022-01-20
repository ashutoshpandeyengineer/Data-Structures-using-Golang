package main

import "fmt"

type Node struct {
	Right *Node //define the types of structs.
	left  *Node
	Value int
}

func (n *Node) Insert(value int) {
	fmt.Println(n.Value)
	if value < n.Value {
		if n.left == nil {
			n.left = &Node{Value: value}
		} else {
			n.left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

func (n *Node) Print() {
	// left first
	if n.left != nil {
		n.left.Print()
	}
	// root
	fmt.Print(n.Value)
	fmt.Print(",")
	// right
	if n.Right != nil {
		n.Right.Print()
	}
}

func main() {
	tree := &Node{Value: 20}

	tree.Insert(15)
	tree.Insert(8)
	tree.Insert(17)
	tree.Insert(25)
	tree.Insert(24)
	fmt.Println("Hello")
	tree.Print()
}
