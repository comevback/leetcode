# 6. Zigzag Conversion

Solved
Medium
Topics
Companies
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P A H N
A P L S I I G
Y I R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);

Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P I N
A L S I G
Y A H R
P I
Example 3:

Input: s = "A", numRows = 1
Output: "A"

Constraints:

1 <= s.length <= 1000
s consists of English letters (lower-case and upper-case), ',' and '.'.
1 <= numRows <= 1000

---

# Code

```go
package main

import "fmt"

func main() {
	s := "PAYPALISHIRING"
	res := convert1(s, 4)
	fmt.Println(res)
}

// convert 将字符串 s 重新排列成 numRows 行的 Zigzag 模式。
func convert(s string, numRows int) string {
	res := "" // 结果字符串
	// 计算一个周期的长度，一个周期包括一个完整的竖直下落和斜向上升
	round := numRows + numRows - 2
	if round == 0 { // 如果 numRows 为 1，即没有Zigzag，直接返回原字符串
		return s
	}
	for i := 0; i < numRows; i++ { // 遍历每一行
		if i == 0 || i == numRows-1 { // 对于第一行和最后一行
			for j := i; j < len(s); j += round { // 每次跳过一个完整周期
				res += string(s[j])
			}
		} else { // 对于中间的行
			gap := 2 * (numRows - i - 1) // 计算中间行的字符间隔
			for j := i; j < len(s); j += round {
				res += string(s[j]) // 添加当前周期的字符
				if j+gap < len(s) {
					res += string(s[j+gap]) // 添加当前行的第二个字符（如果存在）
				}
			}
		}
	}

	return res
}

// convert1 使用模拟行遍历的方式将字符串 s 重新排列成 numRows 行的 Zigzag 模式。
func convert1(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) { // 如果只有一行，或行数不少于字符串长度，直接返回原字符串
		return s
	}

	rows := make([]string, numRows) // 创建一个数组来存储每一行的字符串
	current := 0                    // 当前行的索引
	up := false                     // 当前移动方向，false 表示向下，true 表示向上

	for _, char := range s { // 遍历字符串中的每个字符
		rows[current] += string(char) // 将字符添加到当前行
		// 确定下一个字符的行号
		if current == numRows-1 { // 如果到达最底行
			up = true // 改变方向向上
		} else if current == 0 { // 如果到达最顶行
			up = false // 改变方向向下
		}

		// 根据当前方向更新行索引
		if up {
			current -= 1
		} else {
			current += 1
		}
	}

	res := ""                  // 初始化结果字符串
	for _, row := range rows { // 将所有行合并成一个字符串
		res += row
	}

	return res
}
```
