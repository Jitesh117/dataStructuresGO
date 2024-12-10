package queue

import "errors"

type Queue struct {
	Items []int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) IsEmpty() bool {
	return len(q.Items) == 0
}

func (q *Queue) Push(data int) {
	q.Items = append(q.Items, data)
}

func (q *Queue) Size() int {
	return len(q.Items)
}

func (q *Queue) Front() (int, error) {
	if q.Size() == 0 {
		return 0, errors.New("queue is empty")
	}

	return q.Items[0], nil
}

func (q *Queue) Pop() (int, error) {
	if q.Size() == 0 {
		return 0, errors.New("queue is empty")
	}
	popped := q.Items[0]
	q.Items = q.Items[1:]
	return popped, nil
}

func (q *Queue) Reset() {
	q.Items = []int{}
}
