# 59. Spiral Matrix II
Solved
Medium
Topics
Companies
Given a positive integer n, generate an n x n matrix filled with elements from 1 to n2 in spiral order.


Example 1:
> Input: n = 3
Output: [[1,2,3],[8,9,4],[7,6,5]]
Example 2:

Input: n = 1
Output: [[1]]

Constraints:

1 <= n <= 20

---

# Code
```go
package main

func main() {

}

func generateMatrix(n int) [][]int {
	totalNums := n * n         // 矩阵中的元素总数
	length := n                // 矩阵的边长
	matrix := make([][]int, n) // 创建一个 n 行的二维切片
	for i := range matrix {
		matrix[i] = make([]int, n) // 初始化每一行的切片，长度为 n
	}

	value := 1                                           // 初始化要填入矩阵的值
	right, bottom, left, top := length-1, length-1, 0, 1 // 定义四个边界
	X, Y := 0, 0                                         // 初始化当前位置为矩阵的左上角

	// 当填入的值小于总元素数时，继续循环
	for value <= totalNums {
		// 向右移动
		for right >= left {
			matrix[X][Y] = value // 将当前值填入矩阵
			value += 1           // 值加 1
			if Y < right {
				Y += 1 // 如果没有到达右边界，则继续向右移动
			} else {
				X += 1     // 否则向下移动到下一行
				right -= 1 // 右边界左移
				break
			}
		}

		// 向下移动
		for bottom >= top {
			matrix[X][Y] = value // 将当前值填入矩阵
			value += 1           // 值加 1
			if X < bottom {
				X += 1 // 如果没有到达下边界，则继续向下移动
			} else {
				Y -= 1      // 否则向左移动到上一列
				bottom -= 1 // 下边界上移
				break
			}
		}

		// 向左移动
		for right >= left {
			matrix[X][Y] = value // 将当前值填入矩阵
			value += 1           // 值加 1
			if Y > left {
				Y -= 1 // 如果没有到达左边界，则继续向左移动
			} else {
				X -= 1    // 否则向上移动到上一行
				left += 1 // 左边界右移
				break
			}
		}

		// 向上移动
		for bottom >= top {
			matrix[X][Y] = value // 将当前值填入矩阵
			value += 1           // 值加 1
			if X > top {
				X -= 1 // 如果没有到达上边界，则继续向上移动
			} else {
				Y += 1   // 否则向右移动到下一列
				top += 1 // 上边界下移
				break
			}
		}
	}

	return matrix // 返回生成的螺旋矩阵
}
```