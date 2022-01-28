package main

import "fmt"

//Defined the array size constant for throughout the programme.
const Arraysize = 7

type HashTable struct {
	array [Arraysize]*bucket
}
type bucket struct {
	head *bucketNode
}
type bucketNode struct {
	key  string
	next *bucketNode
}

//Insert will take a key and add it to the hash table array .
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)

}

//Search will take the key  arnd search it throughout the whole array and returns True or false (bool).
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)

}

//Delete take the key string and deletes it from the Table.
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)

}

//insert
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newnode := &bucketNode{key: k}
		newnode.next = b.head
		b.head = newnode
	} else {
		fmt.Println("Already  Exists")
	}

}

//search
func (b *bucket) search(k string) bool {
	currentnode := b.head
	for currentnode != nil {
		if currentnode.key == k {
			return true
		}
		currentnode = currentnode.next
	}
	return false
}

//delete
func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	prevnode := b.head
	for prevnode.next != nil {
		if prevnode.next.key == k {
			prevnode.next = prevnode.next.next
		}
		prevnode = prevnode.next
	}
}

//hashfunction
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % Arraysize
}

//Initializes the bucket for every index of the array...ie; creates a linked list for every instance.
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result

}

func main() {
	test := Init()
	test.Insert("Ram")
	test.Insert("Shyam")

}
