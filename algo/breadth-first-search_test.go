package algo

import "fmt"
import "testing"

func TestBreadthFirstSearch(t *testing.T) {
	g := map[string][]string{
		"Ronaldo": []string{"Valencia", "Bale"},
		"Lingard": []string{"Martial", "Rashford"},
		"Valencia": []string{"Lingard", "Rashford"},
	}
	fmt.Println("The graph", g)

	_, n := BreadthFirstSearch(g, "Ronaldo", "Rashford")
	if n != 2 {
		t.Error("Expected ", 2, ", got ", n)
	}

}
