package main

func PreOrderSearch(head *BinaryNode) *[]int {
	var path []int
	pre_order_walk(head, &path)
	return &path
}

func InOrderSearch(head *BinaryNode) *[]int {
	var path []int
	in_order_walk(head, &path)
	return &path
}

func PostOrderSearch(head *BinaryNode) *[]int {
	var path []int
	post_order_walk(head, &path)
	return &path
}

func pre_order_walk(curr *BinaryNode, path *[]int) {
	if curr == nil {
		return
	}
	*path = append(*path, curr.value)
	pre_order_walk(curr.left, path)
	pre_order_walk(curr.right, path)
}

func in_order_walk(curr *BinaryNode, path *[]int) {
	if curr == nil {
		return
	}
	in_order_walk(curr.left, path)
	*path = append(*path, curr.value)
	in_order_walk(curr.right, path)
}

func post_order_walk(curr *BinaryNode, path *[]int) {
	if curr == nil {
		return
	}
	post_order_walk(curr.left, path)
	post_order_walk(curr.right, path)
	*path = append(*path, curr.value)
}
