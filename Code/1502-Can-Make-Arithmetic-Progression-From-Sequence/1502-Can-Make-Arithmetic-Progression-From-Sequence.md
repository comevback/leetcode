# 1502. Can Make Arithmetic Progression From Sequence

Easy

Hint
A sequence of numbers is called an arithmetic progression if the difference between any two consecutive elements is the same.

Given an array of numbers arr, return true if the array can be rearranged to form an arithmetic progression. Otherwise, return false.

Example 1:
> Input: arr = [3,5,1]
Output: true
Explanation: We can reorder the elements as [1,3,5] or [5,3,1] with differences 2 and -2 respectively, between each consecutive elements.

Example 2:
> Input: arr = [1,2,4]
Output: false
Explanation: There is no way to reorder the elements to obtain an arithmetic progression.
 

Constraints:
> 2 <= arr.length <= 1000
-106 <= arr[i] <= 106

---

# Code
```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 4}
	fmt.Println(canMakeArithmeticProgression(arr))
}

// 排序后遍历差值
func canMakeArithmeticProgression(arr []int) bool {
	// 从小到大排序
	sort.Ints(arr)
	// 初始差值为最后一个元素减倒数第二个元素
	diff := arr[len(arr)-1] - arr[len(arr)-2]
	// 遍历每两个相邻的差值，如果有不相等的，返回false
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	// 如果都想等，返回true
	return true
}
```