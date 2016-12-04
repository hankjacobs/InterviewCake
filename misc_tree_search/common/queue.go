package common

type Queue struct {
	c []*Node
}

func NewQueue() *Queue {
	return &Queue{[]*Node{}}
}

func (q *Queue) Enqueue(n *Node) {
	q.c = append(q.c, n)
}

func (q *Queue) Dequeue() *Node {
	n := q.c[0]
	q.c = q.c[1:]
	return n
}

func (q *Queue) Size() int {
	return len(q.c)
}
