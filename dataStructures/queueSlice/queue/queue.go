package queue

import "errors"

type queue struct {
	slice []int
}

func NewQueue() *queue {
	return &queue{slice: make([]int, 0)}
}

func (q *queue) Enqueue(elems ...int) {
	q.slice = append(q.slice, elems...)
}

func (q *queue) Dequeue() (int, error) {
	l := len(q.slice) - 1
	if l == -1 {
		return 0, errors.New("Empty Queue")
	}
	elem := q.slice[0]
	if l == 0 {
		q.slice = q.slice[:0]
	} else {
		q.slice = q.slice[1:]
	}
	return elem, nil
}

func (q *queue) Peek() (int, error) {
	l := len(q.slice) - 1
	if l == -1 {
		return 0, errors.New("Empty Queue")
	}
	return q.slice[0], nil
}
