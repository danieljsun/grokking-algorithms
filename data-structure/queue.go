package datastructure

type Queue struct {
	S []string
}

func (q Queue) Empty() bool {
	return len(q.S) == 0
}

func (q Queue) Top() string {
	return q.S[0]
}

func (q Queue) Enqueue(x string) Queue {
	q.S = append(q.S, x)
	return q
}

func (q Queue) Dequeue() (string, Queue) {
	x := q.S[0]
	q.S = q.S[1:]
	return x, q
}
