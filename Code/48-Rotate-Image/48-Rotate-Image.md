# 48. Rotate Image（旋转正方形）
Solved
Medium
Topics
Companies
You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.


Example 1:
> Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]

Example 2:
> Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]


Constraints:

n == matrix.length == matrix[i].length
1 <= n <= 20
-1000 <= matrix[i][j] <= 1000

---

# Code
```go
package main

func main() {
	rotate([][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}})
}

func rotate(matrix [][]int) {
	// 计算旋转操作所需的层数
	var length int
	if len(matrix)%2 == 0 {
		length = len(matrix) / 2
	} else {
		length = len(matrix)/2 + 1
	}

	// 遍历每一层
	for i := 0; i < length; i++ {
		// 计算当前层最后一个元素的索引
		leng := len(matrix[0]) - i - 1
		// 遍历当前层的元素进行旋转，每一行遍历到倒数第二个元素即可，每次交换四个元素
		for j := i; j < leng; j++ {
			l := len(matrix) - 1
			// 进行四个元素的交换
			matrix[j][l-i], matrix[l-i][l-j], matrix[l-j][i], matrix[i][j] =
				matrix[i][j], matrix[j][l-i], matrix[l-i][l-j], matrix[l-j][i]
		}
	}
}

// 2.官方解法，先转置矩阵，再翻转每一行
func rotate1(matrix [][]int) {
	// 先从左上到右下沿对角线翻转
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 然后左右翻转
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0])/2; j++ {
			matrix[i][j], matrix[i][len(matrix)-1-j] = matrix[i][len(matrix)-1-j], matrix[i][j]
		}
	}
}
```