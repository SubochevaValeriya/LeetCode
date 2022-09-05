package findCenter

func findCenter(edges [][]int) int {
	if len(edges) == 1 {
		return edges[0][0]
	}

	for i := range edges[0] {
		for j := range edges[1] {
			if edges[0][i] == edges[1][j] {
				return edges[0][i]
			}
		}
	}

	return -1
}
