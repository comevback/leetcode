# 2053. Kth Distinct String in an Array

Solved
Easy
Topics
Companies
Hint
A distinct string is a string that is present only once in an array.

Given an array of strings arr, and an integer k, return the kth distinct string present in arr. If there are fewer than k distinct strings, return an empty string "".

Note that the strings are considered in the order in which they appear in the array.

Example 1:

Input: arr = ["d","b","c","b","c","a"], k = 2
Output: "a"
Explanation:
The only distinct strings in arr are "d" and "a".
"d" appears 1st, so it is the 1st distinct string.
"a" appears 2nd, so it is the 2nd distinct string.
Since k == 2, "a" is returned.
Example 2:

Input: arr = ["aaa","aa","a"], k = 1
Output: "aaa"
Explanation:
All strings in arr are distinct, so the 1st string "aaa" is returned.
Example 3:

Input: arr = ["a","b","a"], k = 3
Output: ""
Explanation:
The only distinct string is "b". Since there are fewer than 3 distinct strings, we return an empty string "".

Constraints:

1 <= k <= arr.length <= 1000
1 <= arr[i].length <= 5
arr[i] consists of lowercase English letters.

---

# Code

```go
package main

import "fmt"

func main() {
	arr := []string{"d", "b", "c", "b", "c", "a"}
	res := kthDistinct(arr, 2)
	fmt.Println(res)
}

// kthDistinct 函数返回数组中第 k 个唯一字符串
func kthDistinct(arr []string, k int) string {
	// 用于存储唯一字符串的结果列表
	res := []string{}
	// 用于记录字符串出现次数的哈希映射
	hsMap := make(map[string]int)

	// 遍历数组，统计每个字符串的出现次数
	for _, v := range arr {
		hsMap[v] += 1
	}

	// 再次遍历数组，找到唯一的字符串并按顺序存储到结果列表中
	for _, v := range arr {
		// 如果字符串的出现次数为1，表示唯一
		if hsMap[v] == 1 {
			res = append(res, v)
			// 当结果列表中的唯一字符串个数等于k时，返回该字符串
			if len(res) == k {
				return v
			}
		}
	}

	// 如果没有找到第k个唯一字符串，返回空字符串
	return ""
}
```
