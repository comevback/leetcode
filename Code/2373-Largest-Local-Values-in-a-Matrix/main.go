package main

import (
	"fmt"
	"math"
)

func main() {
	var grid = [][]int{
		{9, 9, 8, 1},
		{5, 6, 2, 6},
		{8, 2, 6, 4},
		{6, 2, 2, 2},
	}

	res := largestLocal(grid)
	fmt.Println(res)
}

// 循环遍历3*3的二维数组，求出每个3*3的二维数组中的最大值，存入新的二维数组中
func largestLocal(grid [][]int) [][]int {
	length := len(grid)
	// 生成一个新的二维数组，长度为原二维数组的长度减2
	var res [][]int = make([][]int, length-2)
	for i := range res {
		res[i] = make([]int, length-2)
	}

	for i := 0; i <= length-3; i++ {
		for j := 0; j <= length-3; j++ {
			max := threeMax(buildThree(grid, i, j))
			res[i][j] = max
		}
	}

	return res
}

// 从一个更大的二维数组中提取一个以X，Y为左上角的3*3的二维数组
func buildThree(grid [][]int, X int, Y int) [][]int {
	// 生成一个3*3的二维数组
	threeGrid := make([][]int, 3)
	for i := range threeGrid {
		threeGrid[i] = make([]int, 3)
	}

	// 新生成的二维数组的值为原二维数组中以X，Y为左上角的3*3的二维数组的值
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			threeGrid[i][j] = grid[i+X][j+Y]
		}
	}

	return threeGrid
}

// 求一个二维数组（矩阵）中的最大值
func threeMax(grid [][]int) int {
	max := math.MinInt
	for _, v := range grid {
		for _, w := range v {
			if w > max {
				max = w
			}
		}
	}

	return max
}

// ==========================================  简化版 from leetcode  ===============================================
func largestLocal1(grid [][]int) [][]int {
	length := len(grid)
	maxLocal := make([][]int, length-2)

	for i := 0; i < length-2; i++ {
		maxLocal[i] = make([]int, length-2)
		for j := 0; j < length-2; j++ {
			max := grid[i][j]
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					if grid[i+k][j+l] > max {
						max = grid[i+k][j+l]
					}
				}
			}
			maxLocal[i][j] = max
		}
	}

	return maxLocal
}
