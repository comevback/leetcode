package main

import (
	"fmt"
)

func main() {
	grid := [][]int{{1}, {1}, {1}, {0}}
	res := minFlips(grid)
	fmt.Println(res)
}

// minFlips 计算将所有行和列变成回文所需的最小翻转次数
func minFlips(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	flips := 0 // 用于记录翻转的次数
	ones := 0  // 用于记录矩阵中 '1' 的数量

	// 遍历每个 4 元组
	for i := 0; i < m/2; i++ {
		for j := 0; j < n/2; j++ {
			// 获取四个对称位置的值
			tl := grid[i][j]
			tr := grid[i][n-1-j]
			bl := grid[m-1-i][j]
			br := grid[m-1-i][n-1-j]

			// 计算翻转次数，使四个位置的值相同
			if tl+tr+bl+br >= 2 {
				flips += 4 - (tl + tr + bl + br)
				ones += 4
			} else {
				flips += tl + tr + bl + br
			}
		}
	}

	needs := 0  // 用于记录需要翻转的次数
	oneone := 0 // 用于记录中间行/列中相同为 '1' 的数量

	// 处理奇数行的中间行
	if m%2 == 1 {
		for i := 0; i < n/2; i++ {
			if grid[m/2][i] != grid[m/2][n-1-i] {
				ones += 1
				needs += 1
			} else {
				if grid[m/2][i] == 1 {
					oneone += 1
					ones += 2
				}
			}
		}
	}

	// 处理奇数列的中间列
	if n%2 == 1 {
		for i := 0; i < m/2; i++ {
			if grid[i][n/2] != grid[m-1-i][n/2] {
				ones += 1
				needs += 1
			} else {
				if grid[i][n/2] == 1 {
					oneone += 1
					ones += 2
				}
			}
		}
	}

	flips += needs

	// 如果 oneone 是奇数且 needs 为零，则需要额外的翻转
	if oneone%2 == 1 && needs == 0 {
		flips += 2
	}

	// 处理中间元素（行数和列数均为奇数的情况）
	if m%2 == 1 && n%2 == 1 {
		if grid[m/2][n/2] == 1 {
			flips += 1
		}
	}

	return flips
}
