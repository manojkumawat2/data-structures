package linked_list

import "testing"

func TestReverseList(t *testing.T) {
	l := LinkedList{head: nil}
	l.InsertAtTail(1)
	l.InsertAtTail(2)
	l.InsertAtTail(3)
	l.InsertAtTail(4)
	l.InsertAtTail(5)

	l.ReverseList()

	want := "5, 4, 3, 2, 1"

	got := l.GetListData()

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
