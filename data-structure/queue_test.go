package datastructure

//import "fmt"
import "testing"

func TestQueue(t *testing.T) {
	q := Queue{make([]string, 0)}
	//q := Queue()
	q = q.Enqueue("foo")
	if len(q.S) != 1 {
		t.Error("Expected", 1, "got", len(q.S))
	}
	if q.S[0] != "foo" {
		t.Error("Expected", "foo", "got", q.S[0])
	}

	foo, empty := q.Dequeue()
	if foo != "foo" {
		t.Error("Expected", "foo", "got", foo)
	}
	if len(empty.S) != 0 {
		t.Error("Expected", 0, "got", len(empty.S))
	}
}

