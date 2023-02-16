package linked_list

import "testing"

func TestDetectLoop(t *testing.T) {

	// Test Case 1
	l := LinkedList{head: nil}

	l.InsertAtTail(1)
	l.InsertAtTail(2)
	loopNode := &Node{data: 0, next: nil}
	l.InsertNodeAtTail(loopNode)
	l.InsertAtTail(4)
	l.InsertAtTail(5)
	l.InsertNodeAtTail(loopNode)

	want := true
	got := l.DetectLoop()

	if want != got {
		t.Errorf("Want %v, got %v", want, got)
	}

	// Test Case 2
	l2 := LinkedList{head: nil}

	l2.InsertAtHead(1)
	l2.InsertAtHead(0)
	l2.InsertAtTail(2)
	l2.InsertAtTail(3)
	l2.InsertAtTail(4)

	want = false
	got = l2.DetectLoop()

	if want != got {
		t.Errorf("Want %v, got %v", want, got)
	}
}
