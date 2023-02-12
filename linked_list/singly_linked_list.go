package main

import "fmt"

// Linked List Node
type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) InsertAtTail(d int) {
	newNode := &Node{data: d, next: nil}

	if l.head == nil {
		l.head = newNode
	} else {
		temp := l.head

		for temp.next != nil {
			temp = temp.next
		}

		temp.next = newNode
	}
}

func (l *LinkedList) InsertAtHead(d int) {
	newNode := &Node{data: d, next: nil}

	if l.head == nil {
		l.head = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}
}

func (l *LinkedList) Delete(d int) {
	if l.head == nil {
		return
	}

	var prev *Node
	curr := l.head

	for curr != nil && curr.data != d {
		prev = curr
		curr = curr.next
	}

	if curr == nil {
		return
	}

	if prev == nil {
		l.head = l.head.next
	} else {
		prev.next = curr.next
	}
}

func (l *LinkedList) Search(d int) bool {
	curr := l.head

	for curr != nil {
		if curr.data == d {
			return true
		}
		curr = curr.next
	}

	return false
}

func (l *LinkedList) GetHead() *Node {
	return l.head
}

func (l *LinkedList) Show() {
	temp := l.head

	for temp != nil {
		fmt.Print(temp.data, " ")
		temp = temp.next
	}
	fmt.Println()
}

func main() {
	list := LinkedList{head: nil}

	list.InsertAtTail(1)
	list.InsertAtTail(2)
	list.InsertAtTail(3)
	list.InsertAtTail(4)

	list.InsertAtHead(0)

	list.Show()

	list.Delete(1)

	list.Show()

	fmt.Println(list.Search(6))
}
