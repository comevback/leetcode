package main

import (
	"errors"
	"fmt"
)

func main() {
	graph := NewGraph_Map()
	graph.addVertex(1)
	graph.addVertex(2)
	graph.addVertex(4)
	graph.addVertex(5)
	graph.addVertex(3)
	graph.addEdge(1, 5)
	graph.addEdge(2, 5)
	graph.addEdge(1, 4)
	graph.addEdge(3, 4)
	graph.addEdge(2, 4)
	graph.addEdge(1, 5)
	fmt.Println(graph.NumOfEdge)
	fmt.Println(graph.NumOfNode)
	graph.showConnections()
}

// =========================================  map implementation  ==========================================
// 如果要给两个不同的node赋同样的值，这里的map[int][]int可以改成map[string][]string，用字符串代表node，而不是整数
type Graph_Map struct {
	NumOfNode int
	NumOfEdge int
	adjList   map[int][]int
}

func NewGraph_Map() *Graph_Map {
	return &Graph_Map{
		NumOfNode: 0,
		NumOfEdge: 0,
		adjList:   make(map[int][]int),
	}
}

func (graph *Graph_Map) addVertex(node int) error {
	if _, exist := graph.adjList[node]; !exist {
		graph.adjList[node] = []int{}
		graph.NumOfNode += 1
	} else {
		return errors.New("node already exist")
	}
	return nil
}

func (graph *Graph_Map) addEdge(node1 int, node2 int) error {
	_, exist1 := graph.adjList[node1]
	_, exist2 := graph.adjList[node2]
	if !exist1 || !exist2 {
		return errors.New("node not exist")
	}
	graph.adjList[node1] = append(graph.adjList[node1], node2)
	graph.adjList[node2] = append(graph.adjList[node2], node1)
	graph.NumOfEdge += 1
	return nil
}

func (graph *Graph_Map) showConnections() {
	for key, value := range graph.adjList {
		fmt.Println(key, value)
	}
}

// =========================================  array implementation  ==========================================
// type Graph_Node struct {
// 	Val      int
// 	Adjacent []*int
// }

type Graph_Array struct {
	NumOfNode int
	NumOfEdge int
	adjList   [][]int
}

func NewGraph_Array() *Graph_Array {
	return &Graph_Array{
		NumOfNode: 0,
		NumOfEdge: 0,
	}
}

func (graph *Graph_Array) addVertex(value int) {
	graph.adjList = append(graph.adjList, []int{})
}

func (graph *Graph_Array) addEdge(node1 int, node2 int) error {
	if node1 >= 0 && node1 < len(graph.adjList) && node2 > 0 && node2 < len(graph.adjList) {
		graph.adjList[node1] = append(graph.adjList[node1], node2)
		graph.adjList[node2] = append(graph.adjList[node2], node1)
	}
}

func (graph *Graph_Array) showConnections() {

}
