package linked_list

import "testing"

func TestMiddleNode(t *testing.T) {
	// test 1
	l := LinkedList{head: nil}

	l.InsertAtTail(0)
	l.InsertAtTail(1)
	l.InsertAtTail(2)
	l.InsertAtTail(3)
	l.InsertAtTail(4)
	l.InsertAtTail(5)
	l.InsertAtTail(6)
	l.InsertAtTail(7)

	want := 3
	got := l.MiddleElement()

	if want != got {
		t.Errorf("Want %v, got %v", want, got)
	}

	// test 2
	l2 := LinkedList{head: nil}

	l2.InsertAtTail(0)
	l2.InsertAtTail(1)
	l2.InsertAtTail(2)
	l2.InsertAtTail(3)
	l2.InsertAtTail(4)
	l2.InsertAtTail(5)
	l2.InsertAtTail(6)
	l2.InsertAtTail(7)
	l2.InsertAtTail(8)

	want = 4
	got = l2.MiddleElement()

	if want != got {
		t.Errorf("Want %v, got %v", want, got)
	}
}
