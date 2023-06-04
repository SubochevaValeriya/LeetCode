package main

// https://leetcode.com/problems/minimum-number-of-vertices-to-reach-all-nodes/

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	canReach := make([]bool, n)

	for _, edge := range edges {
		canReach[edge[1]] = true
	}

	var result []int
	for vertex, reachable := range canReach {
		if !reachable {
			result = append(result, vertex)
		}
	}

	return result
}

//
//func main() {
//	//n := 6
//	//edges := [][]int{{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}}
//
//	n := 4
//	edges := [][]int{{2, 0}, {0, 3}, {3, 1}}
//	fmt.Println(findSmallestSetOfVertices(n, edges))
//}
//
//func findSmallestSetOfVertices(n int, edges [][]int) []int {
//
//	relations := map[int][]int{}
//
//	for _, edge := range edges {
//		relations[edge[0]] = append(relations[edge[0]], edge[1])
//	}
//
//	result := []int{}
//	reachable := map[int][]int{}
//	for from := 0; from < n; from++ {
//		for _, vertex := range relations[from] {
//			relations[from] = append(relations[from], relations[vertex]...)
//			reachable[vertex] = append(reachable[vertex], from)
//			fmt.Println(vertex, relations[from], from, "G")
//			//fmt.Println(reachable[vertex])
//		}
//	}
//
//	canReach := make([]bool, n)
//	for i := 0; i < n; i++ {
//		fmt.Println(result)
//		if canReach[i] == true {
//			continue
//		}
//		from, ok := reachable[i]
//		if !ok {
//			result = append(result, i)
//			canReach[i] = true
//			for _, to := range relations[i] {
//				canReach[to] = true
//			}
//			continue
//		}
//		maxLenPatch := 0
//		fromVertex := -1
//
//		for _, vertex := range from {
//			if len(relations[vertex]) > maxLenPatch {
//				maxLenPatch = len(relations[vertex])
//				fromVertex = vertex
//			}
//		}
//		result = append(result, fromVertex)
//		canReach[i] = true
//		canReach[fromVertex] = true
//		for _, to := range relations[fromVertex] {
//			canReach[to] = true
//		}
//	}
//
//	fmt.Println(relations, reachable)
//	return result
//}
//
////
////func bfs(relations map[int][]int, start int) map[int][]int {
////	for _, vertex := range relations[start] {
////		relations[start] = append(relations[start], relations[vertex]...)
////	}
////}
