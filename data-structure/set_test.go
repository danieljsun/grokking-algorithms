package datastructure

//import "fmt"
import "testing"

func TestSet(t *testing.T) {
	s := Set{make(map[string]bool)}
	if len(s.S) != 0 {
		t.Error("Expected", 0, "got", len(s.S))
	}
	s.Add("foo")
	if len(s.S) != 1 || !s.Contains("foo") {
		t.Error("Expected", 1, "got", len(s.S))
	}
	s.Add("foo")
	if len(s.S) != 1 {
		t.Error("Expected", 1, "got", len(s.S))
	}
	s.Remove("foo")
	if len(s.S) != 0 {
		t.Error("Expected", 0, "got", len(s.S))
	}
	s.Add("foo")
	s_union := s.Union(s)
	if len(s_union.S) != 1 {
		t.Error("Expected", 1, "got", len(s_union.S))
	}
	if !s_union.Contains("foo") {
		t.Error("Expected", "foo", "got", s_union.S)
	}
	s_inter := s.Intersect(s)
	if len(s_inter.S) != 1 {
		t.Error("Expected", 1, "got", len(s_inter.S))
	}
	if !s_inter.Contains("foo") {
		t.Error("Expected", "foo", "got", s_inter.S)
	}
	s_substract := s.Substract(s)
	if len(s_substract.S) != 0 {
		t.Error("Expected", 0, "got", len(s_substract.S))
	}
}
