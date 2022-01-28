package main

import "fmt"

type Node struct {
	key   int
	left  *Node
	right *Node
}

func (n *Node) Insert(k int) {
	if n.key < k {
		//moves right
		if n.right == nil {
			n.right = &Node{key: k}
		} else {
			n.right.Insert(k)
		}

	} else {
		//moves Left
		if n.left == nil {
			n.left = &Node{key: k}
		} else {
			n.left.Insert(k)
		}
	}

}
func (n *Node) Print() {

	//root

	fmt.Print(n.key)
	fmt.Print(",")

	//left
	if n.left != nil {
		n.left.Print()
	}
	// right
	if n.right != nil {
		n.right.Print()
	}
}

func (n *Node) Search(k int) bool {
	//for the elements that are not available in the tree
	if n == nil {
		return false
	}

	if n.key < k {
		//moves right
		return n.right.Search(k)

	} else if n.key > k {
		//moves left
		return n.left.Search(k)

	}

	return true

}
func main() {
	fmt.Println("Hello")
	n := &Node{key: 20}
	n.Insert(12)
	n.Insert(15)
	n.Insert(21)
	n.Insert(24)
	n.Insert(25)
	n.Print()
	var a int
	fmt.Println("Enter the element to search")
	fmt.Scanf("%d", a)
	n.Search(a)

}
