package main

import (
	"fmt"
)

func main() {
	// [6,0],[1,0],[5,1],[2,5],[3,1],[4,3]
	edges := [][]int{{6, 0}, {1, 0}, {5, 1}, {2, 5}, {3, 1}, {4, 3}}
	res := countGoodNodes(edges)
	fmt.Println(res)
}

func countGoodNodes(edges [][]int) int {
	adjs := make([][]int, len(edges)+1)
	for _, edge := range edges {
		adjs[edge[0]] = append(adjs[edge[0]], edge[1])
		adjs[edge[1]] = append(adjs[edge[1]], edge[0])
	} 
}
