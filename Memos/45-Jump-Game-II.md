# 45. Jump Game II

Medium

You are given a 0-indexed array of integers nums of length n. You are initially positioned at nums[0].

Each element nums[i] represents the maximum length of a forward jump from index i. In other words, if you are at nums[i], you can jump to any nums[i + j] where:

0 <= j <= nums[i] and
i + j < n
Return the minimum number of jumps to reach nums[n - 1]. The test cases are generated such that you can reach nums[n - 1].

 

Example 1:
> Input: nums = [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:
> Input: nums = [2,3,0,1,4]
Output: 2
 

Constraints:
> 1 <= nums.length <= 104
0 <= nums[i] <= 1000
It's guaranteed that you can reach nums[n - 1].

---

```go
package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums))
	nums = []int{3, 4, 3, 1, 0, 7, 0, 3, 0, 2, 0, 3}
	fmt.Println(jump(nums))
}

// ===========================================================================================================================
// 1.从后往前 O(n^2)-O(1)
// 从后往前,找第一个达到target的元素，设为target，再从头找第一个能达到的，设为target，直到nums[0] 可以达到target，返回step+1
func jump1(nums []int) int {
	// 如果只有一个数在数组里，已经是最后一个，步数为0
	length := len(nums)
	if length == 1 {
		return 0
	}
	// 如果只有两个数在数组里，因为题目设定一定可以到达，所以返回1
	if length == 2 {
		return 1
	}

	// 设初始目标为最后一个元素，step也就是跳跃数为0
	target := length - 1
	step := 0
	// 从前往后找，找到第一个能到达target的元素，设为target，然后再从前往后找，找到第一个能到达target的元素，设为target，直到nums[0]可以到达target，返回step+1
	for {
		for i := 0; i < target; i++ {
			// 如果当前元素加上当前元素的值大于等于目标值，说明这个值就是下一个目标值
			if i+nums[i] >= target {
				// 如果是第一个元素，说明已经找到，返回step+1
				if i == 0 {
					return step + 1
				}
				// 如果不是第一个元素，说明还需要继续找，将当前元素设为目标值，step+1
				target = i
				step += 1
			}
		}
	}

}

// ===========================================================================================================================
// 2.从后往前找 O(n^2)-O(n)
func jump2(nums []int) int {
	// 如果只有一个数在数组里，已经是最后一个，肯定能到达，返回true
	length := len(nums)
	if length == 1 {
		return 0
	}
	// 如果只有两个数在数组里，因为题目设定一定可以到达，所以返回1
	if length == 2 {
		return 1
	}

	// 以末尾为目标，找到能达到末尾元素的最前一个元素，设为目标，然后循环找，知道第一个元素nums[0]大于等于目标下标target，则是找到，如果出现断层，返回false
	target := length - 1
	step := 0
	for {
		if nums[0] >= target {
			return step + 1
		} else {
			// 如果找不到能到达target的值，返回false
			arr := find(nums, target)

			// 如果找到多个，取最小的那个
			target = arr[0]
			step += 1

		}
	}

}

// find函数，找到能到达目标的元素，除了最后一个目标自身
func find(nums []int, target int) []int {
	var result []int
	for i := 0; i < target; i++ {
		if i+nums[i] >= target {
			result = append(result, i)
		}
	}
	return result
}

// ===========================================================================================================================
// 3.贪心算法 O(n)-O(1)
func jump(nums []int) int {
	// 如果只有一个数在数组里，已经是最后一个，肯定能到达，返回true
	length := len(nums)
	if length == 1 {
		return 0
	}
	// 如果只有两个数在数组里，因为题目设定一定可以到达，所以返回1
	if length == 2 {
		return 1
	}

	// 一开始是想每次碰到可以跳更远的元素，就跳到这个元素上，但是这会造成步数增多
	// 把当前元素能跳到最远距离nowMaxReach，和遍历中得到的最远距possibleMaxReach分别列出来，不断更新possibleMaxReach，但是只有当i走到了nowMaxReach时，才跳一次。
	nowMaxReach := 0
	possibleMaxReach := 0
	step := 0

	for i := 0; i < length-1; i++ {
		distance := i + nums[i]
		// 如果当前元素加上当前元素的值大于等于目标值，说明这个值就是下一个目标值
		if distance > nowMaxReach && distance > possibleMaxReach {
			possibleMaxReach = distance
		}

		// 如果i走到了nowMaxReach时，才跳一次
		if i == nowMaxReach {
			step += 1
			nowMaxReach = possibleMaxReach
			// 如果nowMaxReach已经大于等于length-1，说明已经到达最后一个元素，直接返回
			if nowMaxReach >= length-1 {
				break
			}
		}
	}
	return step
}
```