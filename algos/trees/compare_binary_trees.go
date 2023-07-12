package trees

func CompareBinaryTrees(a *BinaryNode, b *BinaryNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.value != b.value {
		return false
	}

	return CompareBinaryTrees(a.left, b.left) && CompareBinaryTrees(a.right, b.right)
}
