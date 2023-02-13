package linked_list

/**
* Reverse linked list
 */
func (l *LinkedList) ReverseList() {
	if l.head == nil {
		return
	}

	var prev *Node
	curr := l.head

	for curr != nil {
		temp := curr.next
		curr.next = prev
		prev = curr
		curr = temp
	}

	l.head = prev
}
