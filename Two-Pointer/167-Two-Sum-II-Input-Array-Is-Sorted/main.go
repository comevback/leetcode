package main

func main() {

}

// 1.我的解法，快慢指针同时从左开始走
func twoSum(numbers []int, target int) []int {
	res := []int{}

	slow, fast := 0, 1

	for fast < len(numbers) {
		if numbers[slow]+numbers[fast] >= target {
			if numbers[slow]+numbers[fast] == target {
				res = append(res, slow+1)
				res = append(res, fast+1)
				break
			} else {
				slow -= 1
			}
		} else {
			fast += 1
			slow += 1
		}
	}

	return res
}

// 2. 前后指针分别从左右端向中间靠拢
func twoSum1(nums []int, target int) []int {
	// 一左一右两个指针相向而行
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			// 题目要求的索引是从 1 开始的
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++ // 让 sum 大一点
		} else if sum > target {
			right-- // 让 sum 小一点
		}
	}
	return []int{-1, -1}
}
