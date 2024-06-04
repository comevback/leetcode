# 54. Spiral Matrix
Solved
Medium
Topics
Companies
Hint
Given an m x n matrix, return all elements of the matrix in spiral order.


Example 1:
> Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]

Example 2:
> Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]


Constraints:

m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100

---

# Code
```go
package main

import "fmt"

func main() {
	// matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	res := spiralOrder(matrix)
	fmt.Println(res)
}

// spiralOrder 函数用于以螺旋顺序遍历二维矩阵
func spiralOrder(matrix [][]int) []int {
	// 定义四个边界：右边界、下边界、左边界和上边界
	right, bottom, left, top := len(matrix[0])-1, len(matrix)-1, 0, 1
	totalNum := len(matrix) * len(matrix[0]) // 矩阵中的元素总数
	res := make([]int, 0, totalNum)          // 用于存储结果的切片，容量为矩阵中的元素总数
	i, j := 0, 0                             // 初始化当前的位置为矩阵的左上角

	// 当结果切片中的元素数量小于矩阵中的元素总数时，继续循环
	for len(res) < totalNum {
		// 向右移动
		for len(res) < totalNum {
			res = append(res, matrix[i][j]) // 将当前元素加入结果切片
			if j < right {
				j += 1 // 如果没有到达右边界，则继续向右移动
			} else {
				i += 1     // 否则向下移动到下一行
				right -= 1 // 右边界左移
				break
			}
		}

		// 向下移动
		for len(res) < totalNum {
			res = append(res, matrix[i][j]) // 将当前元素加入结果切片
			if i < bottom {
				i += 1 // 如果没有到达下边界，则继续向下移动
			} else {
				j -= 1      // 否则向左移动到上一列
				bottom -= 1 // 下边界上移
				break
			}
		}

		// 向左移动
		for len(res) < totalNum {
			res = append(res, matrix[i][j]) // 将当前元素加入结果切片
			if j > left {
				j -= 1 // 如果没有到达左边界，则继续向左移动
			} else {
				i -= 1    // 否则向上移动到上一行
				left += 1 // 左边界右移
				break
			}
		}

		// 向上移动
		for len(res) < totalNum {
			res = append(res, matrix[i][j]) // 将当前元素加入结果切片
			if i > top {
				i -= 1 // 如果没有到达上边界，则继续向上移动
			} else {
				j += 1   // 否则向右移动到下一列
				top += 1 // 上边界下移
				break
			}
		}
	}

	return res // 返回结果切片
}
```