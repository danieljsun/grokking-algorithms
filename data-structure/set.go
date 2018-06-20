package datastructure

//import "fmt"

type Set struct {
	S map[string]bool
}

func (s Set) Add(x string) {
	s.S[x] = true;
}

func (s Set) Contains(x string) bool {
	return s.S[x]
}

func (s Set) Remove(x string) {
	delete(s.S, x)
}

func (s Set) Substract(other Set) Set {
	for k, _ := range other.S {
		delete(s.S, k)
	}
	return s
}

func (s Set) Union(other Set) Set {
	result := Set{make(map[string]bool)}
	for m, _ := range s.S {
		result.Add(m)
	}
	for n, _ := range other.S {
		result.Add(n)
	}
	return result
}

func (s Set) Intersect(other Set) Set {
	result := Set{make(map[string]bool)}
	for k, _ := range s.S {
		if other.Contains(k) {
			result.Add(k)
		}
	}
	return result
}
