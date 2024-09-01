# 118. Pascal's Triangle

Solved
Easy
Topics
Companies
Given an integer numRows, return the first numRows of Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:

Example 1:

Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
Example 2:

Input: numRows = 1
Output: [[1]]

Constraints:

1 <= numRows <= 30

---

# Code

```go
package main

func main() {

}

// generate 根据指定的行数生成帕斯卡三角形。
func generate(numRows int) [][]int {
	// 初始化结果数组，这个数组将包含每一行的整数数组。
	res := make([][]int, numRows)
	// 第一行始终是 [1]。
	res[0] = []int{1}

	// 从第二行开始构建帕斯卡三角形。
	for i := 1; i < len(res); i++ {
		// 每行的开始元素是 1。
		res[i] = append(res[i], 1)
		// 计算当前行的中间元素，它们是上一行相邻两数之和。
		for j := 1; j < len(res[i-1]); j++ {
			// res[i-1][j] 和 res[i-1][j-1] 是上一行的相邻两个元素。
			res[i] = append(res[i], res[i-1][j]+res[i-1][j-1])
		}
		// 每行的结束元素也是 1。
		res[i] = append(res[i], 1)
	}

	// 返回构建好的帕斯卡三角形。
	return res
}
```
