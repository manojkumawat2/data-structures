package linked_list

import "testing"

func TestGetLength(t *testing.T) {
	l := LinkedList{head: nil}
	l.InsertAtTail(1)
	l.InsertAtTail(2)
	l.InsertAtHead(0)
	l.InsertAtTail(3)

	want := 4

	got := l.GetLength()

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
