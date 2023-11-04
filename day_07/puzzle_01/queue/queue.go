package queue

type Queue[T any] struct {
	e    []T
	size int
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		e:    make([]T, 0),
		size: 0,
	}
}

func (q *Queue[T]) Enqueue(e T) {
	q.size++
	q.e = append(q.e, e)
}

func (q *Queue[T]) Dequeue() T {
	r := q.e[0]
	q.size--
	q.e = q.e[1:]

	return r
}

func (q *Queue[T]) Size() int {
	return q.size
}
