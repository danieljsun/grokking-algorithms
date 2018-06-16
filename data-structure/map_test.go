package algo

import (
	//"fmt"
	"testing"
)

/*
Map types are reference types, like pointers or slices, and so the value of m above is nil; it doesn't point to an initialized map. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic; don't do that.
*/
func TestMapUninitializedMap(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected ", "not nil recover", ", got ", nil)
		}
	}()
	var m map[string]int
	//panic: assignment to entry in nil map
	m["a"] = 1
}

func TestMap(t *testing.T) {
	var m0 map[string]string
	s, ok := m0["a"]
	if ok {
		t.Error("Expected ", false, ", got ", ok)
	}
	if s != "" {
		t.Error("Expected ", "", ", got ", s)
	}
	var m map[string]int
	n, ok := m["a"]
	if ok {
		t.Error("Expected ", false, ", got ", ok)
	}
	if n != 0 {
		t.Error("Expected ", 0, ", got ", n)
	}
	m = make(map[string]int)
	m["a"] = 1
	if m["a"] != 1 {
		t.Error("Expected ", 1, ", got ", m["a"])
	}
	//To initialize a map with some data, use a map literal:
	m = map[string]int{
		"a": 97,
		"b": 98,
	}
	//To iterate over the contents of a map, use the range keyword:
	for k, v := range m {
		if (k != "a") && (k != "b") {
			t.Error("Expected a or b, got ", k)
		}
		if (v != 97) && (v != 98) {
			t.Error("Expected 97 or 98, got ", v)
		}
	}
}

