package main

import (
	"fmt"
	_ "net/http/pprof"
)

func main() {
	fmt.Println(countSubTrees(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, "abaedcd"))
	//fmt.Println(countSubTrees(4, [][]int{{0, 1}, {1, 2}, {0, 3}}, "bbbb"))

	//fmt.Println(countSubTrees(4, [][]int{{0, 2}, {0, 3}, {1, 2}}, "aeed"))
	//	fmt.Println(countSubTrees(5, [][]int{{0, 1}, {0, 2}, {1, 3}, {0, 4}}, "aabab"))
	//fmt.Println(countSubTrees(25, [][]int{{4, 0}, {5, 4}, {12, 5}, {3, 12}, {18, 3}, {10, 18}, {8, 5}, {16, 8}, {14, 16}, {13, 16}, {9, 13}, {22, 9}, {2, 5}, {6, 2}, {1, 6}, {11, 1}, {15, 11}, {20, 11}, {7, 20}, {19, 1}, {17, 19}, {23, 19}, {24, 2}, {21, 24}}, "hcheiavadwjctaortvpsflssg"))
	//http.ListenAndServe("localhost:8080", nil)
}

type k struct {
	parent        int
	parentParents *k
}

func countSubTrees(n int, edges [][]int, labels string) []int {
	answer := make([]int, n)
	//edgesMap := map[int][]int{}
	edgesMap := map[int]*k{}
	vertexMap := map[int]map[string]int{}
	var node, parent int

	for _, edge := range edges {
		if edge[0] == 0 || edge[1] == 0 {
			switch {
			case edge[0] == 0:
				node = edge[1]
				parent = edge[0]
			case edge[1] == 0:
				node = edge[0]
				parent = edge[1]
			}

			edgesMap[0] = &k{
				parent:        0,
				parentParents: nil,
			}

			vertexMap[parent] = map[string]int{}
			vertexMap[parent][string(labels[parent])]++

		} else {
			if _, ok := edgesMap[edge[1]]; ok {
				node = edge[0]
				parent = edge[1]
			} else {
				node = edge[1]
				parent = edge[0]
			}
		}

		edgesMap[node] = &k{
			parent:        parent,
			parentParents: edgesMap[parent],
		}

		//edgesMap[node] = make([]int, len(edgesMap[parent])+1)
		//copy(edgesMap[node], edgesMap[parent])
		//
		//if len(edgesMap[node]) < 1 {
		//	edgesMap[node] = append(edgesMap[node], parent)
		//} else {
		//	edgesMap[node][len(edgesMap[node])-1] = parent
		//}

		//fmt.Println(edgesMap)
		//edgesMap[node] = append(edgesMap[node], edgesMap[parent]...)
		//
		//vertexMap[node] = 1

		vertexMap[node] = map[string]int{}
		vertexMap[node][string(labels[node])]++

		parentNode := node
		for {
			fmt.Println(node)
			parentNode = edgesMap[parentNode].parent
			fmt.Println(parentNode)

			if _, ok := vertexMap[parentNode]; ok {
				vertexMap[node] = vertexMap[parentNode]
				vertexMap[node][string(labels[node])]++
				break
			}

			if vertexMap[node] == nil {
				vertexMap[node] = map[string]int{}
				vertexMap[node][string(labels[node])]++
			}

			//if labels[parentNode] == labels[node] {
			//	vertexMap[parentNode]++
			//}
			if edgesMap[parentNode].parentParents == nil {
				break
			}

			//node = parentNode
		}

		//for _, parentNode := range edgesMap[node] {
		//	if labels[parentNode] == labels[node] {
		//		vertexMap[parentNode]++
		//	}
		//}
	}

	fmt.Println(edgesMap)
	for i, label := range labels {
		if i == 0 {
			answer[i] = vertexMap[i][string(label)]
		} else {
			answer[i] = vertexMap[i][string(label)]
		}

	}

	return answer
}

//func countSubTrees(n int, edges [][]int, labels string) []int {
//	answer := make([]int, n)
//
//	edgesMap := map[int][]int{}
//	vertexMap := map[int]map[rune]int{}
//	//edgesMap2 := map[int][]int{}
//	//node := edges[0][0]
//	for _, edge := range edges {
//		if edge[0] == 0 || edge[1] == 0 {
//			switch {
//			case edge[0] == 0:
//				edgesMap[edge[0]] = append(edgesMap[edge[0]], edge[1])
//			case edge[1] == 0:
//				edgesMap[edge[1]] = append(edgesMap[edge[1]], edge[0])
//			}
//			vertexMap[edge[0]] = 0
//			vertexMap[edge[1]] = 0
//			continue
//		}
//
//		_, lastVertexRight := vertexMap[edge[1]]
//		_, lastVertexLeft := vertexMap[edge[0]]
//
//		switch {
//		case lastVertexRight:
//			edgesMap[edge[1]] = append(edgesMap[edge[1]], edge[0])
//			vertexMap[edge[0]] = 0
//			if labels[edge[1]]  ==
//		default:
//			edgesMap[edge[0]] = append(edgesMap[edge[0]], edge[1])
//			vertexMap[edge[0]] = 0
//		}
//
//		if lastVertexLeft {
//			vertexMap[edge[1]] = 0
//		}
//
//		//if edge[0] == 0 {
//		//	edgesMap[edge[0]] = append(edgesMap[edge[0]], edge[1])
//		//	vertexMap[edge[0]] = true
//		//	vertexMap[edge[1]] = true
//		//	continue
//		//}
//		//if edge[1] == 0 {
//		//	edgesMap[edge[1]] = append(edgesMap[edge[1]], edge[0])
//		//	vertexMap[edge[0]] = true
//		//	vertexMap[edge[1]] = true
//		//	continue
//		//}
//
//		//if _, ok := vertexMap[edge[1]]; ok {
//		//	edgesMap[edge[1]] = append(edgesMap[edge[1]], edge[0])
//		//	vertexMap[edge[0]] = true
//		//} else if _, ok := vertexMap[edge[0]]; ok {
//		//	edgesMap[edge[0]] = append(edgesMap[edge[0]], edge[1])
//		//	vertexMap[edge[1]] = true
//		//} else {
//		//	edgesMap[edge[0]] = append(edgesMap[edge[0]], edge[1])
//		//	vertexMap[edge[0]] = true
//		//}
//		//_, ok := vertexMap[edge[0]]
//		//_, ok2 := vertexMap[edge[0]]
//		//if !ok && !ok2 {
//		//	vertexMap[edge[0]] = true
//		//}
//		//
//		//
//		//if i > 0 {
//		//
//		//
//		//	if edge[0] == edges[i-1][0] || edge[0] == edges[i-1][1] {
//		//		edgesMap[edge[0]] = append(edgesMap[edge[0]], edge[1])
//		//	} else if edge[1] == edges[i-1][0] || edge[0] == edges[i-1][1] {
//		//		edgesMap[edge[1]] = append(edgesMap[edge[1]], edge[0])
//		//	} else {
//		//		edgesMap[edge[0]] = append(edgesMap[edge[0]], edge[1])
//		//	}
//		//} else {
//		//	edgesMap[edge[0]] = append(edgesMap[edge[0]], edge[1])
//		//}
//	}
//
//	//if _, ok := vertexMap[edge[1]]; !ok {
//	//	edgesMap[edge[0]] = append(edgesMap[edge[0]], edge[1])
//	//	vertexMap[edge[1]] = edge[0]
//	//} else {
//	//	edgesMap[edge[1]] = append(edgesMap[edge[1]], edge[0])
//	//}
//
//	//if edge[0] != node {
//	//	for _, num := range edgesMap[node] {
//	//		fmt.Println(edgesMap)
//	//		if num == edge[1] {
//	//			edgesMap[edge[1]] = append(edgesMap[edge[1]], edge[0])
//	//		} else {
//	//			edgesMap[edge[0]] = append(edgesMap[edge[0]], edge[1])
//	//		}
//	//		break
//	//	}
//	//} else {
//	//	edgesMap[edge[0]] = append(edgesMap[edge[0]], edge[1])
//	//}
//	//}
//	//
//	fmt.Println(edgesMap)
//
//	for i := 0; i < n; i++ {
//		label := labels[i]
//		//if DFS(label, edgesMap, i, labels) > DFS(label, edgesMap2, i, labels) {
//		answer[i] = DFS(label, edgesMap, i, labels) + 1
//		//	} else {
//		//		answer[i] = DFS(label, edgesMap2, i, labels) + 1
//		//	}
//
//	}
//
//	return answer
//}
//
//func DFS(label byte, edgesMap map[int][]int, firstNode int, labels string) int {
//	numberOfNodes := 0
//	for _, node := range edgesMap[firstNode] {
//		if labels[node] == label {
//			numberOfNodes++
//		}
//		numberOfNodes += DFS(label, edgesMap, node, labels)
//	}
//	return numberOfNodes
//}
