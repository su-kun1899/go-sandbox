package queue

type StringQueue interface {
	Enqueue(value string) error
	Dequeue() (string, error)
}

type myQueue struct {
	values []string
	size   int
}

func (q *myQueue) Enqueue(value string) error {
	q.values = append(q.values, value)
	return nil
}

func (q *myQueue) Dequeue() (string, error) {
	value := q.values[0]
	q.values = q.values[1:]
	return value, nil
}

func NewQueue(size int) StringQueue {
	values := make([]string, 0)
	return &myQueue{values: values, size: size}
}
