package linked_list

import "strconv"

// Linked List Node
type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

// Insert a new element at tail
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

// Insert a new element at head
func (l *LinkedList) InsertAtHead(d int) {
	newNode := &Node{data: d, next: nil}

	if l.head == nil {
		l.head = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}
}

// Delete a matching value node
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

// Search an element in the list and return true or false
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

// Get the head node of linked list
func (l *LinkedList) GetHead() *Node {
	return l.head
}

// Check if list is empty or not
func (l *LinkedList) IsEmpty() bool {
	if l.head == nil {
		return true
	}

	return false
}

// Get the list items in a comma separated string
func (l *LinkedList) GetListData() string {
	data := ""
	temp := l.head

	for temp != nil {
		if len(data) > 0 {
			data += ", "
		}
		data += strconv.Itoa(temp.data)
		temp = temp.next
	}

	return data
}

// Add an already created Node to the end of the linked list
func (l *LinkedList) InsertNodeAtTail(node *Node) {
	if l.head == nil {
		l.head = node
		return
	}

	temp := l.head

	for temp.next != nil {
		temp = temp.next
	}

	temp.next = node
}
