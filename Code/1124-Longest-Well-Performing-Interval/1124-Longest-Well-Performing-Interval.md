# 1124. Longest Well-Performing Interval
Solved
Medium
Topics
Companies
Hint
We are given hours, a list of the number of hours worked per day for a given employee.

A day is considered to be a tiring day if and only if the number of hours worked is (strictly) greater than 8.

A well-performing interval is an interval of days for which the number of tiring days is strictly larger than the number of non-tiring days.

Return the length of the longest well-performing interval.


Example 1:
> Input: hours = [9,9,6,0,6,6,9]
Output: 3
Explanation: The longest well-performing interval is [9,9,6].

Example 2:
> Input: hours = [6,6,6]
Output: 0


Constraints:

1 <= hours.length <= 104
0 <= hours[i] <= 16

---

# Code
```go
package main

import "fmt"

func main() {
	// 测试 longestWPI 函数
	hours := []int{9, 9, 6, 0, 6, 6, 9}
	res := longestWPI(hours)
	fmt.Println(res) // 输出结果
}

// longestWPI 函数用于计算工作小时数组中表现良好的最长时间段
func longestWPI(hours []int) int {
	longest := 0               // 初始化记录最长表现良好的时间段长度的变量
	hsMap := make(map[int]int) // 哈希表，用于存储平衡值及其首次出现的索引
	hsMap[0] = -1              // 初始化哈希表，平衡值为 0 的初始索引为 -1
	balance := 0               // 初始化平衡值

	// 遍历 hours 数组，计算每个小时的平衡值
	for i := 0; i < len(hours); i++ {
		if hours[i] <= 8 {
			balance -= 1 // 小时数小于等于 8，平衡值减 1
		} else {
			balance += 1 // 小时数大于 8，平衡值加 1
		}

		if balance > 0 {
			longest = i + 1 // 如果平衡值大于 0，最长时间段为当前索引加 1
		} else {
			if _, exist := hsMap[balance-1]; exist {
				length := i - hsMap[balance-1] // 计算当前时间段的长度
				if length > longest {
					longest = length // 更新最长时间段长度
				}
			}
		}

		if _, exist := hsMap[balance]; !exist {
			hsMap[balance] = i // 将当前平衡值及其索引存入哈希表
		}
	}

	return longest // 返回最长表现良好的时间段长度
}
```