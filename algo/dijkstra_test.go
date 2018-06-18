package algo

import (
	"fmt"
	"testing"
	"reflect"
)

func TestDijkstra(t *testing.T) {
	g := map[string]map[string]int{
		"Start": map[string]int{"A": 6, "B": 2,},
		"A": map[string]int{"Finish": 1,},
		"B": map[string]int{"A": 3, "Finish": 5,},
	}
	fmt.Println("The graph", g)

	found, path, cost := Dijkstra(g, "Start", "Finish")
	if !found {
		t.Error("Expected", false, "got ", found)
	}
	expectedPath := []string{"Start", "B", "A", "Finish"}
	if !reflect.DeepEqual(path, expectedPath) {
		t.Error("Expected", expectedPath, "got", path)
	}
	if cost != 6 {
		t.Error("Expected", 6, ", got", cost)
	}
}
