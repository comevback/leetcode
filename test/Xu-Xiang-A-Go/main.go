package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var H, W, N int
	fmt.Scanf("%d %d %d", &H, &W, &N)
	scanner := bufio.NewScanner(os.Stdin)
	graph := make([][]int, H)
	for i := range graph {
		graph[i] = make([]int, W)
	}

	for scanner.Scan() {
		text := scanner.Text()
		arr := strings.Fields(text)
		hi, _ := strconv.Atoi(arr[0])
		wi, _ := strconv.Atoi(arr[1])
		xi, _ := strconv.Atoi(arr[2])

		width := []int{xi, xi + wi - 1}

		for i := 0; i < H; i++ {
			layerWidth := getWidth(graph[i])
			fmt.Println(layerWidth)
			if i != H-1 && (layerWidth[0] == -1 || width[1] < layerWidth[0] || width[0] > layerWidth[1]) {
				continue
			} else if i == H-1 && (layerWidth[0] == -1 || width[1] < layerWidth[0] || width[0] > layerWidth[1]) {
				for j := i; j > i-hi; j-- {
					setWidth(graph[j], width)
				}
				break
			} else {
				for j := i - 1; j > i-1-hi; j-- {
					setWidth(graph[j], width)
				}
				break
			}
		}

		N -= 1
		if N < 1 {
			break
		}
	}

	stringGraph := convertGraph(graph)
	printGraph(stringGraph)

}

func getWidth(arr []int) []int {
	var head, tail int = 0, len(arr) - 1
	for head <= tail {
		if arr[head] == 0 {
			head += 1
		}
		if arr[tail] == 0 {
			tail -= 1
		} else {
			break
		}
	}
	if head > tail {
		return []int{-1}
	}
	return []int{head, tail}
}

func setWidth(arr1 []int, arr2 []int) {
	for i := arr2[0]; i <= arr2[1]; i++ {
		arr1[i] = 1
	}
}

func printGraph(graph [][]string) {
	for _, v := range graph {
		fmt.Println(v)
	}
}

func convertGraph(arr [][]int) [][]string {
	var stringGraph [][]string = make([][]string, len(arr))
	for i := range stringGraph {
		stringGraph[i] = make([]string, len(arr[0]))
		for j := range stringGraph[i] {
			if arr[i][j] == 1 {
				stringGraph[i][j] = "#"
			} else {
				stringGraph[i][j] = "."
			}
		}
	}

	return stringGraph
}
