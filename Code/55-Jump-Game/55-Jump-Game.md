# 55. Jump Game （赋0）（从后往前）（贪心算法）

Medium

You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.

 

Example 1:
> Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:
> Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
 

Constraints:
> 1 <= nums.length <= 104
0 <= nums[i] <= 105

---

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	var nums []int = []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums))
	nums = []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
	nums = []int{1, 0, 1, 0}
	fmt.Println(canJump(nums))
	nums = []int{0}
	fmt.Println(canJump(nums))
	nums = []int{2, 5, 0, 0}
	fmt.Println(canJump(nums))
	nums = []int{3, 0, 2, 2, 0, 0, 1}
	fmt.Println(canJump(nums))
	nums = []int{1, 2, 3}
	fmt.Println(canJump(nums))
	nums = []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}
	fmt.Println(canJump(nums))
}

// ===========================================================================================================================
// 1.回溯，无法到达的元素，将其值变为0 O(n^2)-O(1)
// 在最坏情况下，时间复杂度接近 O(n^2)，但是不一定会达到最坏情况
func canJump(nums []int) bool {
	// 如果只有一个数在数组里，已经是最后一个，肯定能到达，返回true
	length := len(nums)
	if length < 2 {
		return true
	}
	// 如果第一个元素是0，且数组长度大于1，肯定走不通，返回false
	if nums[0] == 0 && length > 1 {
		return false
	}

	// 从第一个元素开始，distance是当前下标，preDis是上一个位置的下标
	distance := 0
	preDis := 0
	// count是当前元素在可能的情况下，走不通的次数
	count := 0

	// 如果当前元素值不为0，就走能走通的最快路径，也就是走到当前下标distance+nums[distance]的位置
	// 如果当前元素值为0，就走不通，回到上一个位置preDis的位置，然后走第二快的路径，也就是走到distance+nums[distance]-1的位置
	// 第一次走不通，count=1，下一次就走到distance+nums[distance]-1的位置，第二次走不通，count=2，下一次就走到distance+nums[distance]-2的位置，以此类推
	// 如果这个元素的全部可能都走不通，也就是count=nums[distance]，把这个元素的值也变为0，然后回到上一个位置，继续走
	for {
		if nums[distance] == 0 {
			count += 1
			distance = preDis
			// 如果一个元素的全部可能都走不通，把它的值也变为0，
			if nums[distance]-count == 0 {
				nums[distance] = 0
				count = 1
			}
			distance = distance + nums[distance] - count
			if distance < 0 {
				return false
			}
		} else {
			count = 0
			if distance+nums[distance] >= length-1 {
				return true
			} else {
				preDis = distance
				distance = distance + nums[distance]
			}
		}
	}
}

// 这个方法的时间复杂度可能非常高，主要是因为它在遇到0时采用了回溯的策略。
// 它可能需要多次迭代到同一个位置，并且每次迭代都可能递减地尝试从当前位置向前跳跃，直到跳跃次数达到该位置数字的大小。这可能导致在接近最坏的情况下，时间复杂度接近 O(n^2)。
// 最坏情况：考虑一个数组，其中几乎每个元素都是1，但在末尾前有一个0，这可能导致需要从每一个1开始重复尝试，从而导致几乎是平方级别的时间复杂度。

// ===========================================================================================================================
// 2.从后往前找 O(n^2)-O(n)
func canJump1(nums []int) bool {
	// 如果只有一个数在数组里，已经是最后一个，肯定能到达，返回true
	length := len(nums)
	if length < 2 {
		return true
	}
	// 如果第一个元素是0，且数组长度大于1，肯定走不通，返回false
	if nums[0] == 0 && length > 1 {
		return false
	}

	// 以末尾为目标，找到能达到末尾元素的最前一个元素，设为目标，然后循环找，知道第一个元素nums[0]大于等于目标下标target，则是找到，如果出现断层，返回false
	target := length - 1
	for {
		if nums[0] >= target {
			return true
		} else {
			// 如果找不到能到达target的值，返回false
			arr := find(nums, target)
			if len(arr) == 0 {
				return false
			} else {
				// 如果找到多个，取最小的那个
				target = arr[0]
			}
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
// 从前往后找,每次更新最远能到达的位置，最终如果最远能到达的位置大于等于最后一个元素的下标，就是能到达
// 如果在途中，最远能到达的位置小于当前下标，就是走不通，返回false
func canJump2(nums []int) bool {
	// 如果只有一个数在数组里，已经是最后一个，肯定能到达，返回true
	length := len(nums)
	if length < 2 {
		return true
	}
	// 如果第一个元素是0，且数组长度大于1，肯定走不通，返回false
	if nums[0] == 0 && length > 1 {
		return false
	}

	// 设maxReach为最远能到达的位置
	maxReach := 0
	for i := 0; i < length; i++ {
		// 如果本次下标大于最远能到达的位置，就是走不通，返回false
		if i > maxReach {
			return false
		}
		// 如果本次下标加上本次元素的值大于最远能到达的位置，就把这个值赋更新给maxReach
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}
		// 如果最远能到达的位置大于等于最后一个元素的下标，就是能到达，返回true
		if maxReach >= length-1 {
			return true
		}
	}
	return false
}

// 贪心算法 变种 O(n)-O(1)
// 遍历数组，每次都找出当前元素可以走的最远距离，如果发现当前可以走的距离大于dist，就更新dist，在过程中，如果能走的最远距离始终不为负数，就是可以继续走，返回true
func canJump3(nums []int) bool {
	runtime.GC()
	dist := 0 // 这里理解为剩余需要可以跳跃的距离
	for _, n := range nums {
		if dist < 0 {
			return false
		}
		// 如果发现当前可以走的距离大于dist，就更新dist
		if n > dist {
			dist = n
		}
		// 本来的贪心算法是用i+nums[i]和maxReach比较，这里只用了n，因为这里的dist就是maxReach，本身
		dist--
	}
	// 第一遍循环，只要nums[0]不为0，就能走，dist变为nums[0]，最后dist-1 = nums[0]-1
	// 第二遍循环，如果nums[1]大于nums[0]-1，就把dist变为nums[1]，否则dist-1，等于 nums[0]-2
	// 这个过程其实是前一个元素和后一个元素能到的最远距离的比较，
	// 后一个元素因为本身就比前一个元素远一步，所以要在比较时，前一个元素先dist-1，再去和后一个元素比，相当于 (i + nums[i])-1 与（i+1）+ nums[i+1] 比较
	return true
}
```