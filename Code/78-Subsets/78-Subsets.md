# 78. Subsets

Medium

Given an integer array nums of unique elements, return all possible 
subsets(the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.


Example 1:
> Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Example 2:
> Input: nums = [0]
Output: [[],[0]]
 

Constraints:
> 1 <= nums.length <= 10
-10 <= nums[i] <= 10
All the numbers of nums are unique.

---

# Idea

这个程序的思想在于，每一次递归，当前的 nums[index] 到底加不加入temp，都考虑两种情况
index从0开始，到len(temp)-1

  1. nums[0]，是否加入temp，分化为两支，
  2. nums[1], 是否加入，分化为四支，
  3. nums[3], 是否加入，分化为八支......

当每一支的index都到len(temp)之后，不能再加了，所以把这一支的切片，加入到结果数组中。

## 过程
- 包含 1：子集变为 [1]。
  - 包含 2：子集变为 [1, 2]。
    -  包含 3：子集变为 [1, 2, 3]。添加到结果，回溯。
    - 不包含 3：子集回到 [1, 2]。添加到结果，回溯。
  - 不包含 2：子集回到 [1]。
    - 包含 3：子集变为 [1, 3]。添加到结果，回溯。
    - 不包含 3：子集回到 [1]。添加到结果，回溯。
- 不包含 1：子集回到 []。
  - 包含 2：子集变为 [2]。
    - 包含 3：子集变为 [2, 3]。添加到结果，回溯。
    - 不包含 3：子集回到 [2]。添加到结果，回溯。
  - 不包含 2：子集回到 []。
    - 包含 3：子集变为 [3]。添加到结果，回溯。
    - 不包含 3：子集保持 []。添加到结果，结束。


# Code
```go
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	res := subsets(arr)
	fmt.Println(res)
}

func subsets(nums []int) [][]int {
	res := [][]int{}
	getSubs(nums, &res, []int{}, 0)
	return res
}

// 这个程序的思想在于，每一次递归，当前的 nums[index] 到底加不加入temp，都考虑两种情况
// index从0开始，到len(temp)-1
// 1. nums[0]，是否加入temp，分化为两支，
// 2. nums[1], 是否加入，分化为四支，
// 3. nums[3], 是否加入，分化为八支......
// 当每一支的index都到len(temp)之后，不能再加了，所以把这一支的切片，加入到结果数组中。
func getSubs(nums []int, res *[][]int, temp []int, index int) {
	if index == len(nums) {
		t := make([]int, len(temp))
		copy(t, temp)
		*res = append(*res, t)
	} else {
		getSubs(nums, res, temp, index+1)
		temp = append(temp, nums[index])
		getSubs(nums, res, temp, index+1)
	}
}
```