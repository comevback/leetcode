# 556. Next Greater Element III

Solved
Medium
Topics
Companies
Given a positive integer n, find the smallest integer which has exactly the same digits existing in the integer n and is greater in value than n. If no such positive integer exists, return -1.

Note that the returned integer should fit in 32-bit integer, if there is a valid answer but it does not fit in 32-bit integer, return -1.

Example 1:

Input: n = 12
Output: 21
Example 2:

Input: n = 21
Output: -1

Constraints:

1 <= n <= 231 - 1

---

# Code

```go
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	num := 12443322
	res := nextGreaterElement(num)
	fmt.Println(res)
}

// nextGreaterElement finds the next greater element with the same digits.
// nextGreaterElement 寻找具有相同数字的下一个更大元素。
func nextGreaterElement(n int) int {
	// Convert the integer to a string to easily access each digit.
	// 将整数转换为字符串以便于访问每个数字。
	numStr := strconv.Itoa(n)
	// Initialize a slice to hold the digits of the number.
	// 初始化一个切片来存储数字的各个位。
	numArr := []int{}
	// Populate the numArr with individual digits of the number.
	// 用数字的每个位填充 numArr。
	for i := 0; i < len(numStr); i++ {
		num, _ := strconv.Atoi(string(numStr[i]))
		numArr = append(numArr, num)
	}

	// If there's only one digit, return -1 as there can be no greater number.
	// 如果只有一个数字，则返回 -1，因为没有更大的数。
	if len(numArr) == 1 {
		return -1
	}

	// tempArr is used to track the digits we've seen so far in reverse.
	// tempArr 用来反向跟踪我们迄今为止看到的数字。
	tempArr := []int{}
	changed := false // flag to check if we found a valid next greater number.
	// 标志位，用来检查我们是否找到一个有效的下一个更大的数字。
	// Iterate over the number from right to left.
	// 从右到左迭代数字。
	for i := len(numArr) - 1; i >= 0; i-- {
		// If current digit is larger than or equal to the last digit in tempArr, just append it.
		// 如果当前数字大于或等于 tempArr 中的最后一个数字，只需添加它。
		if len(tempArr) == 0 || numArr[i] >= tempArr[len(tempArr)-1] {
			tempArr = append(tempArr, numArr[i])
		} else {
			// Find the smallest digit in tempArr which is bigger than numArr[i]
			// 找到 tempArr 中大于 numArr[i] 的最小数字
			j := 0
			for j < len(tempArr) && numArr[i] >= tempArr[j] {
				j += 1
			}
			// Swap the current digit with this smallest bigger digit found.
			// 将当前数字与找到的最小较大数字交换。
			numArr[i], tempArr[j] = tempArr[j], numArr[i]
			// Copy the sorted tempArr back to numArr to form the next permutation.
			// 将排序好的 tempArr 复制回 numArr 以形成下一个排列。
			copy(numArr[i+1:], tempArr)
			changed = true
			break
		}
	}

	// If no valid next greater number was found, return -1.
	// 如果没有找到有效的下一个更大的数字，返回 -1。
	if !changed {
		return -1
	} else {
		numRes := 0
		// Construct the integer from the digits array.
		// 从数字数组构造整数。
		for i := 0; i < len(numArr); i++ {
			pro := math.Pow10(len(numArr) - 1 - i) // power of 10 needed for each digit.
			// 每个数字需要的 10 的幂。
			tempNum := numArr[i] * int(pro)
			numRes += tempNum
		}
		// Check if the resulting number is larger than the maximum int32 value.
		// 检查结果数字是否大于 int32 的最大值。
		if numRes > math.MaxInt32 {
			return -1
		}
		return numRes
	}
}
```
