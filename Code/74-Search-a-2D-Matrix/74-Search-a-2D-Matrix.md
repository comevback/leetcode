# 74. Search a 2D Matrix
Solved
Medium
Topics
Companies
You are given an m x n integer matrix matrix with the following two properties:

Each row is sorted in non-decreasing order.
The first integer of each row is greater than the last integer of the previous row.
Given an integer target, return true if target is in matrix or false otherwise.

You must write a solution in O(log(m * n)) time complexity.

 

Example 1:


Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true
Example 2:


Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false
 

Constraints:

m == matrix.length
n == matrix[i].length
1 <= m, n <= 100
-104 <= matrix[i][j], target <= 104


# Code
```go
package main

import "fmt"

func main() {
	// 测试用例：单元素矩阵
	matrix := [][]int{{1}}
	res := searchMatrix(matrix, 2) // 搜索目标值2
	fmt.Println(res)               // 输出搜索结果，期望输出false
}

// searchMatrix 在二维矩阵中搜索给定的目标值
func searchMatrix(matrix [][]int, target int) bool {
	// 将二维矩阵转换为一维数组以便进行二分查找
	arr := make([]int, len(matrix)*len(matrix[0])) // 创建一维数组
	index := 0

	// 遍历二维矩阵，填充一维数组
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			arr[index] = matrix[i][j]
			index += 1
		}
	}

	// 设置二分查找的初始边界
	left := 0
	right := len(matrix) * len(matrix[0]) // 计算一维数组的长度

	// 执行二分查找
	for left < right {
		mid := left + (right-left)/2 // 计算中间索引
		if arr[mid] == target {
			return true // 如果中间元素等于目标值，返回true
		} else if arr[mid] < target {
			left = mid + 1 // 如果中间元素小于目标值，调整左边界
		} else if arr[mid] > target {
			right = mid // 如果中间元素大于目标值，调整右边界
		}
	}

	return false // 如果查找结束未找到目标值，返回false
}

// ========================================================================================================
// 不用重新分配，只需转换坐标
func searchMatrix1(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	// 把二维数组映射到一维
	left, right := 0, m*n-1
	// 前文讲的标准的二分搜索框架
	for left <= right {
		mid := left + (right-left)/2
		if get(matrix, mid) == target {
			return true
		} else if get(matrix, mid) < target {
			left = mid + 1
		} else if get(matrix, mid) > target {
			right = mid - 1
		}
	}
	return false
}

// 通过一维坐标访问二维数组中的元素
func get(matrix [][]int, index int) int {
	n := len(matrix[0])
	// 计算二维中的横纵坐标
	i, j := index/n, index%n
	return matrix[i][j]
}
```