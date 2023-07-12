package main

func BreadthFirstSearch(head *BinaryNode, needle int) bool {
	ch := make(chan *BinaryNode, 100)
	ch <- head

	for len(ch) > 0 {
		curr := <-ch

		if curr.value == needle {
			return true
		}

		if curr.left != nil {
			ch <- curr.left
		}

		if curr.right != nil {
			ch <- curr.right
		}
	}
	return false
}
