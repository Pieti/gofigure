package trees

func depthFirstSearch(head *BinaryNode, needle int) bool {
	if head == nil {
		return false
	}

	if head.value == needle {
		return true
	}

	return depthFirstSearch(head.left, needle) || depthFirstSearch(head.right, needle)
}

func DepthFirstSearch(head *BinaryNode, needle int) bool {
	return depthFirstSearch(head, needle)
}
