package main

func main() {

}

// containsNearbyDuplicate 检查在给定的整数切片 nums 中是否存在间隔不超过 k 的重复整数。
func containsNearbyDuplicate(nums []int, k int) bool {
	left, right := 0, 0          // 初始化两个指针 left 和 right，指向切片的起始位置。
	charMap := make(map[int]int) // 创建一个映射，用于存储当前窗口内每个整数的计数。

	for right < len(nums) { // 使用 right 指针遍历整个切片。
		charMap[nums[right]] += 1 // 增加当前整数在映射中的计数。

		if right-left > k { // 如果窗口大小超过了 k，
			charMap[nums[left]] -= 1 // 减少 left 指针所指整数在映射中的计数，
			left += 1                // 然后移动 left 指针向右一步，以保持窗口大小不超过 k。
		}

		if charMap[nums[right]] > 1 { // 如果当前整数的计数大于 1，说明存在重复，
			return true // 返回 true。
		}

		right += 1 // 移动 right 指针向右一步，继续检查下一个整数。
	}

	return false // 如果遍历完整个切片后没有找到满足条件的重复整数，返回 false。
}
