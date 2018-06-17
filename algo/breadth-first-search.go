package algo

import "fmt"

func BreadthFirstSearch(g map[string][]string, from string, to string) (found bool, distance int) {
	q := g[from][:]
	for len(q) > 0 {
		var x string
		x, q = dequeue(q)
		if x != to {
			q = append(q, g[x]...)
		} else {
			return true, 2
		}
		fmt.Println("x", x, "q", q)
	}
	return
}

func dequeue(q []string) (string, []string) {
	top := q[0]
	q2 := q[1:]
	//fmt.Println("top", top, "q2", q2)
	return top, q2
}
