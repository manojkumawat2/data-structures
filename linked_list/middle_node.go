package linked_list

// Find the middle element of the linked list
func (l *LinkedList) MiddleElement() int {
	if l.IsEmpty() {
		return -1
	}

	if l.head.next == nil {
		return l.head.data
	}

	middleNode := l.head
	fastNode := l.head.next.next

	// Move middleNode by one node
	// Move fastNode by two nodes
	for fastNode != nil {
		middleNode = middleNode.next

		fastNode = fastNode.next

		if fastNode != nil {
			fastNode = fastNode.next
		}
	}

	if middleNode != nil {
		return middleNode.data
	}

	return 0
}
