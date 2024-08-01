# 14. Longest Common Prefix

Solved
Easy
Topics
Companies
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters.

---

# Code

```go
package main

import "fmt"

func main() {
	strs := []string{"flower", "flow", "flight"}
	res := longestCommonPrefix(strs)
	fmt.Println(res)
}

func longestCommonPrefix1(strs []string) string {
	common := strs[0]
	for _, v := range strs {
		common = getPrefix(common, v)
		if common == "" {
			return ""
		}
	}
	return common
}

func getPrefix(str1 string, str2 string) string {
	minLen := min(len(str1), len(str2))
	res := ""
	for i := 0; i < minLen; i++ {
		if str1[i] == str2[i] {
			res += string(str1[i])
		} else {
			break
		}
	}

	return res
}

func min(num1 int, num2 int) int {
	if num1 > num2 {
		return num2
	} else {
		return num1
	}
}

// 字典树
func longestCommonPrefix(strs []string) string {

}
```
