package graphs

type WeightedAdjacencyMatrix [][]int

func BreadthFirstSearch(graph WeightedAdjacencyMatrix, source, needle int) []int {
	seen := make([]bool, len(graph))
	prev := make([]int, len(graph))
	q := make(chan int, len(graph))

	for i := range prev {
		prev[i] = -1
	}

	seen[source] = true
	q <- source

	for len(q) > 0 {
		curr := <-q
		if curr == needle {
			break
		}

		adjs := graph[curr]
		for i := 0; i < len(graph); i++ {
			if adjs[i] == 0 {
				continue
			}

			if seen[i] {
				continue
			}

			seen[i] = true
			prev[i] = curr

			q <- i
		}
	}

	curr := needle
	out := []int{}

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	if len(out) == 0 {
		return nil
	}

	out = append(out, source)

	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}

	return out
}
