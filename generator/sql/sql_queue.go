package sql

type Queue struct {
	items []string
}

func NewQueue() *Queue {
	q := Queue{}
	q.items = make([]string, 0)
	return &q
}

func (from *Queue) Enqueue(item string) *Queue {
	from.items = append(from.items, item)
	return from
}

func (from *Queue) Dequeue() string {
	item := from.items[0]
	from.items = from.items[1:len(from.items)]
	return item
}

func (from *Queue) IsEmpty() bool {
	return len(from.items) == 0
}

func (from *Queue) Size() int {
	return len(from.items)
}
