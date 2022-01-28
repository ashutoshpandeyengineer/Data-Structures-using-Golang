package main

import (
	"fmt"
)

type Graph struct {
	vertices []*Vertex
}
type Vertex struct {
	key      int
	adjacent []*Vertex
}

func (g *Graph) Addvertex(k int) {
	if contains(g.vertices, k) { //this ststement checks using contains method...if it is true then vertex is existing
		err := fmt.Errorf("Vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())

	} else {

		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

//contins---> This is used to check wheather the key is in the vertex or not.
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}
func (g *Graph) Addedge(from, to int) {
	fromVertex := g.getvertex(from)
	toVertex := g.getvertex(to)

	//check if the vertices are existing or not .
	if fromVertex == nil || toVertex == nil { // if the vertices are not existing.
		err := fmt.Errorf("Invalid edge(%v %v)", from, to)
		fmt.Println(err.Error())

	} else if contains(fromVertex.adjacent, to) { //check if their is existing edge or not....if existing edge is there then itb will through error.
		err := fmt.Errorf("Existing edges(%v----> %v)", from, to)
		fmt.Println(err.Error())

	} else { //else the vertices are existing and the edges are not
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

func (g Graph) getvertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v :", v.key)
		for _, v := range v.adjacent {
			fmt.Printf("%v ", v.key)
		}
	}
	fmt.Println()
}
func main() {
	test := Graph{}
	for i := 0; i < 5; i++ {
		test.Addvertex(i)
	}
	test.Print()
	test.Addedge(1, 2)
	test.Addedge(6, 2)
	test.Addvertex(6)
	test.Addedge(6, 2)
	test.Print()
	test.Addedge(1, 2)
}
