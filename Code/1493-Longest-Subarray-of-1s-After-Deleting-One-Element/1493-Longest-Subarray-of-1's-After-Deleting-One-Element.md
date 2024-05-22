# 1493. Longest Subarray of 1's After Deleting One Element

Medium

Given a binary array nums, you should delete one element from it.

Return the size of the longest non-empty subarray containing only 1's in the resulting array. Return 0 if there is no such subarray.


Example 1:
> Input: nums = [1,1,0,1]
Output: 3
Explanation: After deleting the number in position 2, [1,1,1] contains 3 numbers with value of 1's.

Example 2:
> Input: nums = [0,1,1,1,0,1,1,0,1]
Output: 5
Explanation: After deleting the number in position 4, [0,1,1,1,1,1,0,1] longest subarray with value of 1's is [1,1,1,1,1].

Example 3:
> Input: nums = [1,1,1]
Output: 2
Explanation: You must delete one element.
 

Constraints:
> 1 <= nums.length <= 105
nums[i] is either 0 or 1.

---

# Code
```go
package main

import "fmt"

func main() {
	arr := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	res := longestSubarray(arr)
	fmt.Println(res)
}

// 1.我的解法：维护1元素index队列
// 定义一个容量为2的队列，每当遇到0时，就把0的下一个塞进队列，如果已经有两个元素，把头元素退出
// length的长度：如果队列有两个，说明中间是间隔了一个0的，也就是删除了一个0，长度需要减掉这个0，所以是 length = i - oneIndex[0]
// 否则 length = i - oneIndex[0] + 1
func longestSubarray(nums []int) int {
	var length int
	var maxLength int
	var oneIndex [2]int // 记录了1元素下标的队列

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if len(oneIndex) == 0 {
				oneIndex[0] = i + 1
			} else if len(oneIndex) == 1 {
				oneIndex[1] = i + 1
			} else {
				oneIndex[0] = oneIndex[1]
				oneIndex[1] = i + 1
			}
		}

		if len(oneIndex) == 2 {
			length = i - oneIndex[0]
		} else {
			length = i - oneIndex[0] + 1
		}

		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
```