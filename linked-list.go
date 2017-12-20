package main

import "fmt"

type Node struct {
	value interface{}
	next  *Node
	prev  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func main() {
	l := LinkedList{}

	l.Insert(1)
	l.Insert(2)
	l.Insert(3)

	l.Print()
	l.Reverse()
	l.Print()
}

func (L *LinkedList) Insert(v interface{}) {
	//create new node
	list := &Node{
		value: v,
		next:  L.head,
	}

	//if head non empty, assign new list to previous address
	if L.head != nil {
		L.head.prev = list
	}

	L.head = list

	l := L.head

	for l.next != nil {
		l = l.next
	}

	L.tail = l
}

func (L *LinkedList) Print() {
	list := L.head
	for list != nil {
		fmt.Printf("%+v -> ", list.value)
		list = list.next
	}
	fmt.Println()
}

func (L *LinkedList) Reverse() {
	tmp := L.head
	var prev *Node

	L.tail = L.head

	for tmp != nil {
		next := tmp.next
		tmp.next = prev
		prev = tmp
		tmp = next
	}

	L.head = prev
}
