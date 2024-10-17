package main

import (
	"fmt"
)

func main() {
	deck := &Deck{}
	deck.AddFront(1)
	deck.AddBack(2)
	deck.AddFront(3)

	fmt.Println(deck.IsExist(2))
	fmt.Println(deck.IsExist(4))

	value, ok := deck.PopFront()
	if ok {
		fmt.Println("PopFront:", value)
	}
	value, ok = deck.PopBack()
	if ok {
		fmt.Println("PopBack:", value)
	}
}

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

type Deck struct {
	Tail *Node
	Head *Node
}

func (d *Deck) AddFront(value int) {
	newNode := &Node{Value: value, Next: d.Head, Prev: nil}
	if d.Head != nil {
		d.Head.Prev = newNode
	}
	d.Head = newNode
	if d.Tail == nil {
		d.Tail = newNode
	}
}

func (d *Deck) AddBack(value int) {
	newNode := &Node{Value: value, Next: nil, Prev: d.Tail}
	if d.Tail != nil {
		d.Tail = newNode
	}
	d.Tail = newNode
	if d.Head == nil {
		d.Head = newNode
	}
}

func (d *Deck) PopFront() (int, bool) {
	if d.Head == nil {
		return 0, false
	}
	value := d.Head.Value
	d.Head = d.Head.Next
	if d.Head == nil {
		d.Tail = nil
	} else {
		d.Head.Prev = nil
	}
	return value, true
}

func (d *Deck) PopBack() (int, bool) {
	if d.Tail == nil {
		return 0, false
	}
	value := d.Tail.Value
	d.Tail = d.Tail.Prev
	if d.Tail == nil {
		d.Head = nil
	} else {
		d.Tail.Prev = nil
	}
	return value, true
}

func (d *Deck) IsExist(value int) bool {
	current := d.Head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}
