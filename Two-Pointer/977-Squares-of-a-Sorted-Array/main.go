package main

import "fmt"

func main() {
	arr := []int{-1}
	res := sortedSquares(arr)
	fmt.Println(res)
}

func sortedSquares(nums []int) []int {
	// 如果数组第一个元素大于等于0，那么直接平方后返回
	if nums[0] >= 0 {
		for i := range nums {
			nums[i] = nums[i] * nums[i]
		}
		return nums
	}
	// 如果数组最后一个元素小于等于0，那么先反转数组，然后平方后返回
	if nums[len(nums)-1] <= 0 {
		reverse(nums)
		for i := range nums {
			nums[i] = nums[i] * nums[i]
		}
		return nums
	}

	// 如果数组中有正数和负数，那么找到正数和负数的分界点，然后分别向两边遍历，比较大小，小的先平方后加入结果数组，最后返回结果数组
	res := []int{}
	index := 0
	for nums[index] < 0 {
		index += 1
	}

	left := index - 1
	right := index

	for left >= 0 && right < len(nums) {
		if -nums[left] < nums[right] {
			res = append(res, nums[left]*nums[left])
			left -= 1
		} else {
			res = append(res, nums[right]*nums[right])
			right += 1
		}
	}

	for left >= 0 {
		res = append(res, nums[left]*nums[left])
		left -= 1
	}

	for right < len(nums) {
		res = append(res, nums[right]*nums[right])
		right += 1
	}

	return res
}

func reverse(nums []int) {
	head, tail := 0, len(nums)-1
	for head < tail {
		nums[head], nums[tail] = nums[tail], nums[head]
		head += 1
		tail -= 1
	}
}
