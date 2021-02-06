package sql

type Queue struct {
	items []string
}

func NewQueue() *Queue {
	q := Queue{}
	q.items = make([]string, 0)
	return &q
}

func (q *Queue) Enqueue(item string) *Queue {
	q.items = append(q.items, item)
	return q
}

func (q *Queue) Dequeue() string {
	item := q.items[0]
	q.items = q.items[1:len(q.items)]
	return item
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}
