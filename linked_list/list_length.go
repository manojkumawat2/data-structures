package linked_list

/*
* Find the length of linked list
 */
func (l *LinkedList) GetLength() int {
	count := 0
	curr := l.head

	for curr != nil {
		count++
		curr = curr.next
	}

	return count
}
