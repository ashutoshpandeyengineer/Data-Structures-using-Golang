package main

import "fmt"

type Queue struct {
	data []int
}

func (q *Queue) Enqueue(i int) {
	q.data = append(q.data, i)
}
func (q *Queue) Dequeue() int {
	toremove := q.data[0]
	q.data = q.data[1:]
	return toremove
}
func main() {
	myqueu := Queue{}
	myqueu.Enqueue(12)
	myqueu.Dequeue()
	fmt.Println("Done")
}
