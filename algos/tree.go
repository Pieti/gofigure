package main

type BinaryNode struct {
	value int
	right *BinaryNode
	left  *BinaryNode
}

func NewBinaryTree() *BinaryNode {
	return &BinaryNode{
		value: 20,
		right: &BinaryNode{
			value: 50,
			right: &BinaryNode{
				value: 100,
			},
			left: &BinaryNode{
				value: 30,
				right: &BinaryNode{
					value: 45,
				},
				left: &BinaryNode{
					value: 29,
				},
			},
		},
		left: &BinaryNode{
			value: 10,
			right: &BinaryNode{
				value: 15,
			},
			left: &BinaryNode{
				value: 5,
				right: &BinaryNode{
					value: 7,
				},
			},
		},
	}
}
