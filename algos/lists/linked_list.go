package lists

import "fmt"

type List interface {
	Length() int
	Append(item any)
	Prepend(item any)
	InsertAt(item any, index int)
	Remove(item any) (any, bool)
	RemoveAt(index int) (any, bool)
	Get(index int) (any, bool)
}

type Node struct {
	item any
	next *Node
	prev *Node
}

func (n Node) Print() {
	fmt.Printf("%+v\n", &n)
}

type LinkedList struct {
	size int
	head *Node
	tail *Node
}

func (l *LinkedList) Print() {
	fmt.Println()
	for n := l.head; n != nil; n = n.next {
		n.Print()
	}
}

func (l *LinkedList) Length() int {
	return l.size
}

func (l *LinkedList) Append(item any) {
	l.InsertAt(item, l.size)
}

func (l *LinkedList) Prepend(item any) {
	l.InsertAt(item, 0)
}

func (l *LinkedList) InsertAt(item any, index int) {
	if index < 0 || index > l.size {
		return
	}

	if index == 0 {
		if l.head == nil {
			l.head = &Node{item: item}
			l.tail = l.head
		} else {
			l.head.prev = &Node{item: item, next: l.head}
			l.head = l.head.prev
		}
	} else if index == l.size {
		if l.head == nil {
			l.head = &Node{item: item}
			l.tail = l.head
		} else {
			l.tail.next = &Node{item: item, prev: l.tail}
			l.tail = l.tail.next
		}
	} else {
		n := l.head
		for i := 0; i < index-1; i++ {
			n = n.next
		}

		n.next = &Node{item: item, prev: n, next: n.next}
		if n.next.next != nil {
			n.next.next.prev = n.next
		}
	}

	l.size++
}

func (l *LinkedList) Remove(item any) (any, bool) {
	if l.size == 0 {
		return nil, false
	}

	n := l.head
	for n != nil {
		if n.item == item {
			if n.prev != nil {
				n.prev.next = n.next
			} else {
				l.head = n.next
			}
			if n.next != nil {
				n.next.prev = n.prev
			} else {
				l.tail = n.prev
			}
			l.size--
			return n.item, true
		}
		n = n.next
	}
	return nil, false
}

func (l *LinkedList) RemoveAt(index int) (any, bool) {
	n, ok := l.getAt(index)
	if ok {
		if n.prev != nil {
			n.prev.next = n.next
		} else {
			l.head = n.next
		}
		if n.next != nil {
			n.next.prev = n.prev
		} else {
			l.tail = n.prev
		}
		l.size--
		return n.item, true
	}
	return nil, false
}

func (l *LinkedList) Get(index int) (any, bool) {
	n, ok := l.getAt(index)
	if ok {
		return n.item, true
	}
	return nil, false
}

func (l *LinkedList) getAt(index int) (*Node, bool) {
	n := l.head
	if n == nil {
		return nil, false
	}

	for i := 0; i < index; i++ {
		n = n.next
		if n == nil {
			return nil, false
		}
	}
	return n, true
}
