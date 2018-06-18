package algo

import "fmt"

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func prep(g map[string]map[string]int, from string) (costs map[string]int, parents map[string]string) {
	costs = make(map[string]int)
	parents = make(map[string]string)
	for k, v := range g {
		for j, u := range v {
			if k == from {
				costs[j] = u
				parents[j] = k
			}/* else {
				//costs[j] = MaxInt
				//parents[j] = ""
			}*/
		}
	}
	return
}

func findCheapest(costs map[string]int, processed map[string]int) (node string) {
	cost := MaxInt
	for k,v := range costs {
		if _, found := processed[k]; !found && v < cost {
			node = k
			cost = v
		}
	}
	return
}

func Dijkstra(g map[string]map[string]int, from string, to string) (found bool, path []string, cost int) {
	costs, parents := prep(g, from)
	processed := make(map[string]int)
	fmt.Println("costs", costs, "parents", parents)

	node := findCheapest(costs, processed)
	for node != "" {
		cost := costs[node]
		fmt.Println("cheapest node", node, cost)
		for k,v := range g[node] {
			newCost := cost + v
			if oldCost, ok := costs[k]; !ok || oldCost > newCost {
				costs[k] = newCost
				parents[k] = node
			}
		}
		processed[node] = 1
		node = findCheapest(costs, processed)
	}

	_, found = parents[to]
	if found {
		path = []string{to}
		for path[0] != from {
			path = append([]string{parents[path[0]]}, path...)
		}
	}
	return found, path, costs[to]
}
