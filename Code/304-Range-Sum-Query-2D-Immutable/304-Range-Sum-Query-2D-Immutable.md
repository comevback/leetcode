# 304. Range Sum Query 2D - Immutable
Solved
Medium
Topics
Companies
Given a 2D matrix matrix, handle multiple queries of the following type:

Calculate the sum of the elements of matrix inside the rectangle defined by its upper left corner (row1, col1) and lower right corner (row2, col2).
Implement the NumMatrix class:

NumMatrix(int[][] matrix) Initializes the object with the integer matrix matrix.
int sumRegion(int row1, int col1, int row2, int col2) Returns the sum of the elements of matrix inside the rectangle defined by its upper left corner (row1, col1) and lower right corner (row2, col2).
You must design an algorithm where sumRegion works on O(1) time complexity.


Example 1:
![sum-grid](https://assets.leetcode.com/uploads/2021/03/14/sum-grid.jpg)
> Input
["NumMatrix", "sumRegion", "sumRegion", "sumRegion"]
[[[[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]]], [2, 1, 4, 3], [1, 1, 2, 2], [1, 2, 2, 4]]
Output
[null, 8, 11, 12]

Explanation
NumMatrix numMatrix = new NumMatrix([[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]]);
numMatrix.sumRegion(2, 1, 4, 3); // return 8 (i.e sum of the red rectangle)
numMatrix.sumRegion(1, 1, 2, 2); // return 11 (i.e sum of the green rectangle)
numMatrix.sumRegion(1, 2, 2, 4); // return 12 (i.e sum of the blue rectangle)
 

Constraints:

m == matrix.length
n == matrix[i].length
1 <= m, n <= 200
-104 <= matrix[i][j] <= 104
0 <= row1 <= row2 < m
0 <= col1 <= col2 < n
At most 104 calls will be made to sumRegion.

---

# Code
```go
package main

func main() {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	na := Constructor(matrix)
	res := na.SumRegion(2, 1, 4, 3)
	println(res)
}

// 设计NumMatrix结构，包含Matrix和Prefix两个属性
type NumMatrix struct {
	Matrix [][]int
	Prefix [][]int
}

// 构造函数
func Constructor(matrix [][]int) NumMatrix {
	// 初始化Prefix数组，长宽与	matrix 相同
	prefix := make([][]int, len(matrix))
	for i := range prefix {
		prefix[i] = make([]int, len(matrix[0]))
	}

	// 计算matrix中每个元素的前缀和（包含自身）
	for i := 0; i < len(matrix); i++ {
		// sum 用于记录当前行里每个元素的前缀和
		sum := 0
		// 遍历当前行的每个元素
		for j := 0; j < len(matrix[0]); j++ {
			// 计算当前元素的前缀和，每个元素的前缀和等于其左边元素的前缀和加上其自身
			sum += matrix[i][j]
			// 如果不是第一行，将当前元素的前缀和加上上一行同列元素的前缀和
			// 因为上一行的前缀和已经包含了左边元素的前缀和，所以还要减去左上角元素的前缀和
			if i > 0 {
				if j > 0 {
					sum += prefix[i-1][j] - prefix[i-1][j-1]
				} else {
					sum += prefix[i-1][j]
				}
			}
			// 将当前元素的前缀和赋值给prefix数组
			prefix[i][j] = sum
		}
	}

	return NumMatrix{
		Prefix: prefix,
		Matrix: matrix,
	}
}

// 计算指定区域的和
func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	// 获取右下角元素的前缀和
	res := this.Prefix[row2][col2]

	// 如果不是第一行，将结果减去上一行右下角元素的前缀和
	if row1 > 0 {
		res -= this.Prefix[row1-1][col2]
	}

	// 如果不是第一列，将结果减去左下角元素的前缀和
	if col1 > 0 {
		res -= this.Prefix[row2][col1-1]
	}

	// 如果不是第一行第一列，将结果加上左上角元素的前缀和，因为之前减去了两次
	if row1 > 0 && col1 > 0 {
		res += this.Prefix[row1-1][col1-1]
	}

	return res
}
```