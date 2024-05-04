package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums []int = []int{3, 5, 1, 5}
	fmt.Println(maxOperations1(nums, 2))
}

// 1. map
// 遍历每个元素，元素值为v，如果hsMap[k-v]存在，这个元素不加入map，同时hsMap[k-v]-1，最大操作数加1
// 如果hsMap[k-v]不存在，则hsMap[v]加1
func maxOperations(nums []int, k int) int {
	var hsMap map[int]int = make(map[int]int)
	maxOp := 0

	for _, v := range nums {
		if hsMap[k-v] > 0 {
			hsMap[k-v] -= 1
			maxOp += 1
		} else {
			hsMap[v] += 1
		}
	}

	return maxOp
}

// 2.双指针 排序
// 先排序，然后使用双指针，head指向头，tail指向尾，从两侧向中间遍历，如果nums[head]+nums[tail] == k，head+1，tail-1，maxOp+1
func maxOperations1(nums []int, k int) int {
	// 定义最大操作数，头尾指针
	maxOp := 0
	var head, tail int = 0, len(nums) - 1
	// 把数组从小到大排序
	sort.Ints(nums)

	// 用双指针从两侧向中间遍历
	for head < tail {
		// 如果头尾指针的和大于k，尾指针向左移动
		if nums[head]+nums[tail] > k {
			tail -= 1
		}

		// 如果头尾指针的和小于k，头指针向右移动
		if nums[head]+nums[tail] < k {
			head += 1
		}

		// 如果前面的遍历操作导致头尾指针相遇，结束遍历
		if head > tail {
			break
		}

		// 如果头尾指针的和等于k，头尾指针同时向中间移动，最大操作数加1
		if nums[head]+nums[tail] == k && head != tail {
			maxOp += 1
			head += 1
			tail -= 1
		}
	}

	// 返回最大操作数
	return maxOp
}

// 在这个问题中，无法用数组代替map和排序，
// 要用数组代替map或排序，必须条件是，数组people有一个明确的最大值，
// 在（881. Boats to Save People）这个问题里，最大值就是limit，所以这个数组是可预期的有限长度。
// 但是在这个问题里，数组nums没有明确的最大值，nums可能极长，所以无法用数组代替map和排序。
