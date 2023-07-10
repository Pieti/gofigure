package datastructures

type QNode struct {
	value any
	next  *QNode
}

type Queue struct {
	size int
	head *QNode
	tail *QNode
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(value any) {
	node := &QNode{value: value}
	if q.head == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.size++
}

func (q *Queue) Dequeue() any {
	if q.head == nil {
		return nil
	}
	node := q.head
	q.head = node.next
	q.size--
	return node.value
}

func (q *Queue) Peek() any {
	if q.head == nil {
		return nil
	}
	return q.head.value
}
