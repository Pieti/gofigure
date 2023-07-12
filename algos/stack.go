package main

type SNode struct {
	value any
	prev  *SNode
}

type Stack struct {
	size int
	head *SNode
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(value any) {
	node := &SNode{value: value}
	if s.head == nil {
		s.head = node
	} else {
		node.prev = s.head
		s.head = node
	}
	s.size++
}

func (s *Stack) Pop() any {
	if s.head == nil {
		return nil
	}
	node := s.head
	s.head = node.prev
	s.size--
	return node.value
}

func (s *Stack) Peek() any {
	if s.head == nil {
		return nil
	}
	return s.head.value
}
