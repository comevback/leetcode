# 2352. Equal Row and Column Pairs

Medium

Given a 0-indexed n x n integer matrix grid, return the number of pairs (ri, cj) such that row ri and column cj are equal.

A row and column pair is considered equal if they contain the same elements in the same order (i.e., an equal array).


Example 1:
![ex1](https://assets.leetcode.com/uploads/2022/06/01/ex1.jpg)
> Input: grid = [[3,2,1],[1,7,6],[2,7,7]]
Output: 1
Explanation: There is 1 equal row and column pair:
- (Row 2, Column 1): [2,7,7]

Example 2:
![ex2](https://assets.leetcode.com/uploads/2022/06/01/ex2.jpg)
> Input: grid = [[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]]
Output: 3
Explanation: There are 3 equal row and column pairs:
- (Row 0, Column 0): [3,1,2,2]
- (Row 2, Column 2): [2,4,2,2]
- (Row 3, Column 2): [2,4,2,2]
 

Constraints:
> n == grid.length == grid[i].length
1 <= n <= 200
1 <= grid[i][j] <= 105

---

# Code
```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	grid := [][]int{{11, 1}, {1, 11}}
	fmt.Println(equalPairs(grid))
}

// 1. 我的解法
// 把行和列都变成字符串，在map里保存每个行字符串的出现次数
// 然后遍历每个列字符串，如果在map里出现，就加上这个字符串的出现次数，总和就是总的配对数。
func equalPairs(grid [][]int) int {
	hsMap := make(map[string]int)
	columns := []string{}
	pair := 0
	// 把每一行变成字符串，存进map，同时把这一列字符串加到一个数组里
	for i := 0; i < len(grid); i++ {
		str := arrToStr(grid[i])
		hsMap[str] += 1
		columns = append(columns, (getColumn(grid, i)))
	}

	// 遍历列字符串数组，在map中找相同的行字符串，如果找到相同的，累加行字符串出现的次数
	for _, v := range columns {
		if hsMap[v] > 0 {
			pair += hsMap[v]
		}
	}

	return pair
}

// 获取列字符串的方式
func getColumn(grid [][]int, column int) string {
	// 先把这一列变成一个数组
	var arr []int
	for i := 0; i < len(grid); i++ {
		arr = append(arr, grid[i][column])
	}
	// 然后用数组变字符串方式变成字符串，返回
	res := arrToStr(arr)
	return res
}

// 数组转化字符串方法
func arrToStr(arr []int) string {
	var str string

	for _, v := range arr {
		str += strconv.Itoa(v) + "," // 加逗号是为了防止[[11,1],[1,11]]这样前后交替变成字符串，都是“111”，分辨不出来的情况。
	}

	return str
}

// ============================================================================================================
// 2. 更简略的写法
func equalPairs1(grid [][]int) int {
	// 获取矩阵的大小
	n := len(grid)

	// 创建一个map，用来存储每一行的数组及其出现次数
	m := make(map[[200]int]int)

	// 创建一个长度为200的整型数组，用来临时存储行或列的数据
	arr := [200]int{}

	// 遍历每一行，将每一行的数据存储到arr数组中，并在map中记录其出现次数
	for i := 0; i < n; i++ {
		// 将当前行的数据复制到arr数组中
		copy(arr[:], grid[i])
		// 在map中记录arr数组出现的次数
		m[arr]++
	}

	// 初始化结果变量，用于统计符合条件的对数
	res := 0

	// 遍历每一列，检查该列与map中的行是否相同
	for i := 0; i < n; i++ {
		// 清空arr数组
		arr = [200]int{}
		// 将当前列的数据复制到arr数组中
		for j := 0; j < n; j++ {
			arr[j] = grid[j][i]
		}
		// 检查该列的数组是否在map中出现过
		if v, ok := m[arr]; ok {
			// 如果存在，则将其出现次数累加到结果中
			res += v
		}
	}

	// 返回结果，即符合条件的对数
	return res
}
```