package linked_list

func (l *LinkedList) DetectLoop() bool {
	// Using Floyd’s Cycle Finding Algorithm
	if l.head == nil {
		return false
	}

	slow_ptr := l.head
	fast_ptr := l.head.next

	for fast_ptr != nil && fast_ptr.next != nil && slow_ptr != fast_ptr {
		slow_ptr = slow_ptr.next
		fast_ptr = fast_ptr.next.next
	}

	if slow_ptr == fast_ptr {
		return true
	}

	return false
}
