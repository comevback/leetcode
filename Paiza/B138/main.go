package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type mapNode struct {
	nodeMap [][]int
}

func NewMapNode(w int, h int) *mapNode {
	newNode := make([][]int, w)

	for i := range newNode {
		newNode[i] = make([]int, h)
	}

	return &mapNode{
		nodeMap: newNode,
	}
}

func (m *mapNode) add(w int, h int) {
	m.nodeMap[w][h] = 1
}

func main() {
	var height, width int
	fmt.Scanf("%d %d", &height, &width)
	Nmap := NewMapNode(width, height)
	h := height - 1
	var dountCount int = 0

	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	str := scanner.Text()
	// 	for i := 0; i < width; i++ {
	// 		if str[i] == '.' {
	// 			Nmap.add(i, h)
	// 		}
	// 	}

	// 	h -= 1
	// 	if h < 0 {
	// 		break
	// 	}
	// }

	file, err := os.OpenFile("./test-case_1.txt", os.O_RDONLY, 0644)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for reader {
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		if line



		if err == io.EOF{
			break
		}
	}

	for i := 1; i < width-1; i++ {
		for j := 1; j < height-1; j++ {
			if Nmap.nodeMap[i][j] == 1 && Nmap.nodeMap[i-1][j] == 0 && Nmap.nodeMap[i+1][j] == 0 && Nmap.nodeMap[i][j+1] == 0 && Nmap.nodeMap[i][j-1] == 0 {
				dountCount += 1
			}
		}
	}

	fmt.Println(Nmap.nodeMap)
	fmt.Println(dountCount)
}

// func main() {
// 	var height, width int
// 	fmt.Scanf("%d %d", &height, &width)
// 	var dountCount int

// 	var strArr []string = make([]string, height)
// 	scanner := bufio.NewScanner(os.Stdin)
// 	hCount := 0
// 	for scanner.Scan() {
// 		strArr[hCount] = scanner.Text()
// 		hCount += 1
// 		if hCount == height {
// 			break
// 		}
// 	}

// 	for i := 1; i < len(strArr)-1; i++ {
// 		for j := 1; j < len(strArr[i])-1; j++ {
// 			if strArr[i][j] == '.' && strArr[i-1][j] == '#' && strArr[i+1][j] == '#' && strArr[i][j-1] == '#' && strArr[i][j+1] == '#' {
// 				dountCount += 1
// 			}
// 		}
// 	}

// 	fmt.Println(dountCount)
// }
