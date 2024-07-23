# 496. Next Greater Element I

Solved
Easy
Topics
Companies
The next greater element of some element x in an array is the first greater element that is to the right of x in the same array.

You are given two distinct 0-indexed integer arrays nums1 and nums2, where nums1 is a subset of nums2.

For each 0 <= i < nums1.length, find the index j such that nums1[i] == nums2[j] and determine the next greater element of nums2[j] in nums2. If there is no next greater element, then the answer for this query is -1.

Return an array ans of length nums1.length such that ans[i] is the next greater element as described above.

Example 1:

Input: nums1 = [4,1,2], nums2 = [1,3,4,2]
Output: [-1,3,-1]
Explanation: The next greater element for each value of nums1 is as follows:

- 4 is underlined in nums2 = [1,3,4,2]. There is no next greater element, so the answer is -1.
- 1 is underlined in nums2 = [1,3,4,2]. The next greater element is 3.
- 2 is underlined in nums2 = [1,3,4,2]. There is no next greater element, so the answer is -1.
  Example 2:

Input: nums1 = [2,4], nums2 = [1,2,3,4]
Output: [3,-1]
Explanation: The next greater element for each value of nums1 is as follows:

- 2 is underlined in nums2 = [1,2,3,4]. The next greater element is 3.
- 4 is underlined in nums2 = [1,2,3,4]. There is no next greater element, so the answer is -1.

Constraints:

1 <= nums1.length <= nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 104
All integers in nums1 and nums2 are unique.
All the integers of nums1 also appear in nums2.

Follow up: Could you find an O(nums1.length + nums2.length) solution?

---

# Code

```go
package main

import "fmt"

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	res := nextGreaterElement(nums1, nums2)
	fmt.Println(res)
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	hsMap := make(map[int]int) // 创建哈希表来存储 nums2 中每个元素的索引
	for i, v := range nums2 {
		hsMap[v] = i // 存储每个元素及其索引
	}
	nextHigher := make([]int, len(nums2))  // 创建一个数组来保存 nums2 中每个元素的下一个更大元素
	temp := []int{}                        // 使用一个栈来帮助找出下一个更大的元素
	for i := len(nums2) - 1; i >= 0; i-- { // 从后向前遍历 nums2
		for len(temp) > 0 && nums2[i] >= temp[len(temp)-1] {
			temp = temp[:len(temp)-1] // 弹出栈中小于等于当前元素的所有元素
		}
		if len(temp) == 0 {
			nextHigher[i] = -1 // 如果栈为空，则当前元素没有更大的下一个元素
		} else {
			nextHigher[i] = temp[len(temp)-1] // 栈顶元素是下一个更大的元素
		}
		temp = append(temp, nums2[i]) // 当前元素入栈
	}

	res := []int{}
	for _, v := range nums1 {
		res = append(res, nextHigher[hsMap[v]]) // 使用哈希表和 nextHigher 数组来找出 nums1 中每个元素的下一个更大元素
	}
	return res // 返回结果
}
```
